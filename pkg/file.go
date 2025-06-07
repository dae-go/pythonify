package py

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"reflect"
	"slices"
	"strings"
)

var (
	ErrFileClosed  = errors.New("I/O operation on closed file")
	ErrInvalidSeek = errors.New("invalid seek operation")
	ErrInvalidMode = errors.New("invalid file mode")
	ErrReadOnly    = errors.New("file not open for writing")
	ErrWriteOnly   = errors.New("file not open for reading")
)

type File struct {
	file   *os.File
	name   string
	mode   string
	closed bool
	reader *bufio.Reader
	writer *bufio.Writer
}

func NewFile(file *os.File, name string, mode string) *File {
	return &File{
		file:   file,
		name:   name,
		mode:   mode,
		closed: false,
		reader: bufio.NewReader(file),
		writer: bufio.NewWriter(file),
	}
}

func (f *File) Close() error {
	if f.closed {
		return nil
	}

	if err := f.writer.Flush(); err != nil {
		return err
	}

	f.closed = true
	return f.file.Close()
}

func (f *File) Fileno() (int, error) {
	if f.closed {
		return -1, ErrFileClosed
	}
	return int(f.file.Fd()), nil
}

func (f *File) Flush() error {
	if f.closed {
		return ErrFileClosed
	}
	return f.writer.Flush()
}

func (f *File) Isatty() bool {
	if f.closed {
		return false
	}

	fileInfo, err := f.file.Stat()
	if err != nil {
		return false
	}

	mode := fileInfo.Mode()
	return mode&os.ModeCharDevice != 0 && mode&os.ModeNamedPipe == 0
}

func (f *File) Read(size ...int) ([]byte, error) {
	if f.closed {
		return nil, ErrFileClosed
	}

	if !f.Readable() {
		return nil, ErrWriteOnly
	}

	readSize := -1
	if len(size) > 0 {
		readSize = size[0]
	}

	if readSize < 0 {
		data, err := io.ReadAll(f.reader)
		return data, err
	}

	if readSize == 0 {
		return []byte{}, nil
	}

	buf := make([]byte, readSize)
	n, err := f.reader.Read(buf)

	if err == io.EOF && n > 0 {
		return buf[:n], nil
	}

	return buf[:n], err
}

func (f *File) Readable() bool {
	options := []string{"r", "r+", "w+", "a+"}
	return !f.closed && slices.Contains(options, f.mode)
}

func (f *File) Readline(size ...int) (string, error) {
	if f.closed {
		return "", ErrFileClosed
	}

	if !f.Readable() {
		return "", ErrWriteOnly
	}

	maxSize := -1
	if len(size) > 0 {
		maxSize = size[0]
	}

	if maxSize < 0 {
		line, err := f.reader.ReadString('\n')
		return line, err
	}

	if maxSize == 0 {
		return "", nil
	}

	var buf bytes.Buffer
	for i := 0; i < maxSize; i++ {
		b, err := f.reader.ReadByte()
		if err != nil {
			if err == io.EOF && buf.Len() > 0 {
				return buf.String(), nil
			}
			return buf.String(), err
		}
		buf.WriteByte(b)
		if b == '\n' {
			break
		}
	}
	return buf.String(), nil
}

func (f *File) Readlines(hint ...int) ([]string, error) {
	if f.closed {
		return nil, ErrFileClosed
	}

	if !f.Readable() {
		return nil, ErrWriteOnly
	}

	var lines []string
	totalSize := 0
	maxSize := -1

	if len(hint) > 0 {
		maxSize = hint[0]
	}

	for {
		line, err := f.Readline()
		if err != nil {
			if err == io.EOF {
				if line != "" {
					lines = append(lines, line)
				}
				break
			}
			return lines, err
		}

		if line == "" {
			break
		}

		lines = append(lines, line)
		totalSize += len(line)

		if maxSize > 0 && totalSize >= maxSize {
			break
		}
	}

	return lines, nil
}

func (f *File) Seek(offset int, whence ...int) (int, error) {
	if f.closed {
		return -1, ErrFileClosed
	}

	w := io.SeekStart
	if len(whence) > 0 {
		w = whence[0]
	}

	if w < 0 || w > 2 {
		return -1, ErrInvalidSeek
	}

	newPos, err := f.file.Seek(int64(offset), w)
	if err != nil {
		return -1, err
	}

	// Reset buffers after seek
	f.reader = bufio.NewReader(f.file)
	f.writer = bufio.NewWriter(f.file)

	return int(newPos), nil
}

func (f *File) Seekable() bool {
	if f.closed {
		return false
	}

	_, err := f.file.Seek(0, io.SeekCurrent)
	return err == nil
}

func (f *File) Tell() (int, error) {
	if f.closed {
		return -1, ErrFileClosed
	}

	currentPos, err := f.file.Seek(0, io.SeekCurrent)
	if err != nil {
		return -1, err
	}

	// Account for buffered data
	if f.reader != nil {
		currentPos -= int64(f.reader.Buffered())
	}

	if f.writer != nil {
		currentPos += int64(f.writer.Buffered())
	}

	return int(currentPos), nil
}

func (f *File) Truncate(size ...int) error {
	if f.closed {
		return ErrFileClosed
	}

	if !f.Writable() {
		return ErrReadOnly
	}

	var truncateSize int64
	if len(size) > 0 {
		truncateSize = int64(size[0])
	} else {
		currentPos, err := f.Tell()
		if err != nil {
			return err
		}
		truncateSize = int64(currentPos)
	}

	return f.file.Truncate(truncateSize)
}

func (f *File) Writable() bool {
	options := []string{"r+", "w", "w+", "a", "a+"}
	return !f.closed && slices.Contains(options, f.mode)
}

func (f *File) Write(b []byte) (int, error) {
	if f.closed {
		return 0, ErrFileClosed
	}

	if !f.Writable() {
		return 0, ErrReadOnly
	}

	n, err := f.writer.Write(b)
	return n, err
}

func (f *File) Writelines(lines []string) error {
	if f.closed {
		return ErrFileClosed
	}

	if !f.Writable() {
		return ErrReadOnly
	}

	for _, line := range lines {
		_, err := f.writer.WriteString(line)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Mode() string {
	return f.mode
}

func (f *File) Closed() bool {
	return f.closed
}

func (f *File) Get(t string) any {
	v := reflect.ValueOf(f).Elem()
	field := v.FieldByName(strings.ToTitle(t))

	if !field.IsValid() {
		panic("not a gettable value")
	}

	return field.Interface()
}
