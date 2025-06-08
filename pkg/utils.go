package py

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"reflect"
	"slices"
	"strconv"
	"unicode/utf8"

	"github.com/dae-go/pythonify/pkg/abf"
)

type kvp [2]any
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}
type Sliceable[T any] interface {
	ToSlice() []T
}

func Sliced[T any](iterator any) []T {
	result := *new([]T)
	switch v := iterator.(type) {
	case []T:
		result = append(result, v...)
	case Sliceable[T]:
		result = v.ToSlice()
	default:
		panic("unsupported type")
	}
	return result
}

func DivMod[T Number](x, y T) (int, float64) {
	raw := x / y
	q := math.Trunc(float64(raw))
	r := float64(raw) - q
	return int(q), r
}

func Pow[T Number](x, y T) T {
	rVal := x
	q, r := DivMod(y, 1)
	for range q {
		rVal *= x
	}
	if r != 0 {
		rVal *= T(r * float64(x))
	}

	return rVal
}

func Print(val ...any) {
	fmt.Println(val...) // LOL
}

// Type conversion and creation functions
func Abs[T Number](x T) T {
	return T(math.Abs(float64(x)))
}

func Ascii(obj any) string {
	// Get string representation
	str := fmt.Sprintf("%v", obj)

	// Escape non-ASCII characters
	result := make([]byte, 0, len(str)*2)
	result = append(result, '\'')

	for _, r := range str {
		if r < 32 || r > 126 {
			// Non-printable or non-ASCII character
			switch r {
			case '\a':
				result = append(result, '\\', 'a')
			case '\b':
				result = append(result, '\\', 'b')
			case '\f':
				result = append(result, '\\', 'f')
			case '\n':
				result = append(result, '\\', 'n')
			case '\r':
				result = append(result, '\\', 'r')
			case '\t':
				result = append(result, '\\', 't')
			case '\v':
				result = append(result, '\\', 'v')
			case '\\':
				result = append(result, '\\', '\\')
			case '\'':
				result = append(result, '\\', '\'')
			default:
				if r < 0x10000 {
					result = append(result, fmt.Sprintf("\\u%04x", r)...)
				} else {
					result = append(result, fmt.Sprintf("\\U%08x", r)...)
				}
			}
		} else {
			result = append(result, byte(r))
		}
	}

	result = append(result, '\'')
	return string(result)
}

func Bin(x int) string {
	if x >= 0 {
		return "0b" + strconv.FormatInt(int64(x), 2)
	}
	// Handle negative numbers with two's complement
	return "-0b" + strconv.FormatInt(int64(-x), 2)
}

func ByteArray(source any, encoding string, errors string) []byte {
	switch v := source.(type) {
	case string:
		return []byte(v)
	case []byte:
		return slices.Clone(v)
	case int:
		// Create a byte array of given size
		return make([]byte, v)
	default:
		// Try to convert to string first
		str := fmt.Sprintf("%v", source)
		return []byte(str)
	}
}

func Bytes(source any, encoding string, errors string) []byte {
	// Bytes is immutable in Python, but in Go we'll return a byte slice
	return ByteArray(source, encoding, errors)
}

func Chr(i int) string {
	if i < 0 || i > 0x10FFFF {
		panic("chr() arg not in range(0x110000)")
	}
	return string(rune(i))
}

func Complex(real float64, imag float64) complex128 {
	return complex(real, imag)
}

func Hex(x int) string {
	if x >= 0 {
		return "0x" + strconv.FormatInt(int64(x), 16)
	}
	return "-0x" + strconv.FormatInt(int64(-x), 16)
}

func Oct(x int) string {
	if x >= 0 {
		return "0o" + strconv.FormatInt(int64(x), 8)
	}
	return "-0o" + strconv.FormatInt(int64(-x), 8)
}

func Ord(c string) int {
	if len(c) == 0 {
		panic("ord() expected a character, but string of length 0 found")
	}
	r, size := utf8.DecodeRuneInString(c)
	if size != len(c) {
		panic(fmt.Sprintf("ord() expected a character, but string of length %d found", len(c)))
	}
	return int(r)
}

// Object and attribute functions
func Callable(obj any) bool {
	if obj == nil {
		return false
	}
	t := reflect.TypeOf(obj)
	return t.Kind() == reflect.Func
}

func Dir(obj any) []string {
	if obj == nil {
		return []string{}
	}

	var result []string
	t := reflect.TypeOf(obj)

	// Get fields if it's a struct
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() == reflect.Struct {
		for i := range t.NumField() {
			result = append(result, t.Field(i).Name)
		}
	}

	// Get methods
	for i := range t.NumMethod() {
		result = append(result, t.Method(i).Name)
	}

	slices.Sort(result)
	return result
}

func Hash(obj any) int {
	// Simple hash implementation
	str := fmt.Sprintf("%v", obj)
	hash := 0
	for _, r := range str {
		hash = 31*hash + int(r)
	}
	return hash
}

func IsInstance(obj any, classinfo any) bool {
	if obj == nil || classinfo == nil {
		return false
	}

	objType := reflect.TypeOf(obj)

	if t, ok := classinfo.(reflect.Type); ok {
		return objType.AssignableTo(t)
	}

	if types, ok := classinfo.([]reflect.Type); ok {
		return slices.ContainsFunc(types, func(t reflect.Type) bool {
			return objType.AssignableTo(t)
		})
	}

	return false
}

func IsSubClass(class any, classinfo any) bool {
	classType, ok := class.(reflect.Type)
	if !ok {
		return false
	}

	if t, ok := classinfo.(reflect.Type); ok {
		return classType.AssignableTo(t)
	}

	if types, ok := classinfo.([]reflect.Type); ok {
		return slices.ContainsFunc(types, func(t reflect.Type) bool {
			return classType.AssignableTo(t)
		})
	}

	return false
}

func Repr(obj any) string {
	if obj == nil {
		return "None"
	}

	switch v := obj.(type) {
	case string:
		return strconv.Quote(v)
	default:
		return fmt.Sprintf("%#v", obj)
	}
}

func Type(obj any) any {
	if obj == nil {
		return "NoneType"
	}
	return reflect.TypeOf(obj)
}

func Vars(obj any) map[string]any {
	result := make(map[string]any)

	if obj == nil {
		return result
	}

	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return result
	}

	t := v.Type()
	for i := range v.NumField() {
		field := v.Field(i)
		if field.CanInterface() {
			result[t.Field(i).Name] = field.Interface()
		}
	}

	return result
}

// Math and numeric functions
func Round(number float64, ndigits int) float64 {
	if ndigits == 0 {
		return math.Round(number)
	}

	multiplier := math.Pow(10, float64(ndigits))
	return math.Round(number*multiplier) / multiplier
}

func Input(prompt string) string {
	if prompt != "" {
		fmt.Print(prompt)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func Open(file string, mode string, buffering int, encoding string, errors string, newline string, closefd bool, opener func(string, int) int) any {

	return nil
}

func Len[T comparable](val Sliceable[T]) int {
	return len(val.ToSlice())
}

// Generic versions of ABF

func Filter[T comparable](fn func(T) bool, iterator Sliceable[T]) List[T] {
	return abf.Filter(fn, iterator.ToSlice())
}

func Map[T any, U comparable](fn func(T) U, iterator Sliceable[T]) List[U] {
	return abf.Map(fn, iterator.ToSlice())
}

func Reversed[T comparable](iterable Sliceable[T]) List[T] {
	return NewList(abf.Reversed(iterable.ToSlice())...)
}

func Sorted[T cmp.Ordered](iterable Sliceable[T], reverse bool) List[T] {
	return NewList(abf.Sorted(iterable.ToSlice(), reverse)...)
}

func SortedFunc[T comparable](iterable Sliceable[T], key func(T, T) int, reverse bool) List[T] {
	return NewList(abf.SortedFunc(iterable.ToSlice(), key, reverse)...)
}

func Zip[T comparable](iters ...Sliceable[T]) []List[T] {
	slices := make([][]T, 0, len(iters))
	for _, i := range iters {
		slices = append(slices, i.ToSlice())
	}
	
	zipped := abf.Zip(slices...)
	result := make([]List[T], len(zipped))
	for i, z := range zipped {
		result[i] = List[T](z)
	}
	return result
}

func Any[T any](iterator Sliceable[T], predicate func(t T) bool) bool {
	return abf.Any(iterator.ToSlice(), predicate)
}

func All[T any](iterator Sliceable[T], predicate func(t T) bool) bool {
	return abf.All(iterator.ToSlice(), predicate)
}

func Sum[T cmp.Ordered | string](iterator Sliceable[T]) T {
	return abf.Sum(iterator.ToSlice())
}

// This is not an OG python function but... why not?
func Product[T Number](iterator Sliceable[T]) T {
	return abf.Product(iterator.ToSlice())
}
