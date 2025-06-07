package py

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

type File struct {
	file     *os.File
	name     string
	mode     string
	closed   bool
	reader   *bufio.Reader
	writer   *bufio.Writer
	position int64
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
	f.closed = true
	f.writer.Flush()
	return f.file.Close()
}

func (f *File) Fileno() int {
	return int(f.file.Fd())
}

func (f *File) Flush() {
	f.writer.Flush()
}

func (f *File) Isatty() bool {
	fileInfo, err := f.file.Stat()
	if err != nil {
		return false
	}
	return fileInfo.Mode()&os.ModeCharDevice != 0
}

func (f *File) Read(size ...int) []byte {
	if f.closed {
		return nil
	}
	
	readSize := -1
	if len(size) > 0 {
		readSize = size[0]
	}
	
	if readSize < 0 {
		data, _ := io.ReadAll(f.reader)
		f.position += int64(len(data))
		return data
	}
	
	buf := make([]byte, readSize)
	n, _ := f.reader.Read(buf)
	f.position += int64(n)
	return buf[:n]
}

func (f *File) Readable() bool {
	return !f.closed && (f.mode == "r" || f.mode == "r+" || f.mode == "w+" || f.mode == "a+")
}

func (f *File) Readline(size ...int) string {
	if f.closed {
		return ""
	}
	
	maxSize := -1
	if len(size) > 0 {
		maxSize = size[0]
	}
	
	if maxSize < 0 {
		line, _ := f.reader.ReadString('\n')
		f.position += int64(len(line))
		return line
	}
	
	var buf bytes.Buffer
	for i := 0; i < maxSize; i++ {
		b, err := f.reader.ReadByte()
		if err != nil {
			break
		}
		buf.WriteByte(b)
		f.position++
		if b == '\n' {
			break
		}
	}
	return buf.String()
}

func (f *File) Readlines(hint ...int) []string {
	if f.closed {
		return nil
	}
	
	var lines []string
	totalSize := 0
	maxSize := -1
	
	if len(hint) > 0 {
		maxSize = hint[0]
	}
	
	for {
		line := f.Readline()
		if line == "" {
			break
		}
		lines = append(lines, line)
		totalSize += len(line)
		if maxSize > 0 && totalSize >= maxSize {
			break
		}
	}
	
	return lines
}

func (f *File) Seek(offset int, whence ...int) int {
	if f.closed {
		return -1
	}
	
	w := 0
	if len(whence) > 0 {
		w = whence[0]
	}
	
	newPos, err := f.file.Seek(int64(offset), w)
	if err != nil {
		return -1
	}
	
	f.position = newPos
	f.reader = bufio.NewReader(f.file)
	f.writer = bufio.NewWriter(f.file)
	
	return int(newPos)
}

func (f *File) Seekable() bool {
	_, err := f.file.Seek(0, io.SeekCurrent)
	return err == nil
}

func (f *File) Tell() int {
	return int(f.position)
}

func (f *File) Truncate(size ...int) {
	if f.closed {
		return
	}
	
	var truncateSize int64
	if len(size) > 0 {
		truncateSize = int64(size[0])
	} else {
		truncateSize = f.position
	}
	
	f.file.Truncate(truncateSize)
}

func (f *File) Writable() bool {
	return !f.closed && (f.mode == "w" || f.mode == "w+" || f.mode == "a" || f.mode == "a+" || f.mode == "r+")
}

func (f *File) Write(b []byte) int {
	if f.closed {
		return 0
	}
	
	n, _ := f.writer.Write(b)
	f.position += int64(n)
	return n
}

func (f *File) Writelines(lines []string) {
	if f.closed {
		return
	}
	
	for _, line := range lines {
		f.writer.WriteString(line)
		f.position += int64(len(line))
	}
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