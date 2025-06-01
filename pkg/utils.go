package pkg

import (
	"cmp"
	"fmt"
	"math"
	"slices"
)

type kvp [2]any
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Zip[T []any](iters ...T) []T {
	if len(iters) == 0 {
		return []T{}
	}

	var (
		minLen = len(iters[0]) // updates in next for loop
		zipped = []T{}
	)

	for _, i := range iters {
		if len(i) < minLen {
			minLen = len(i)
		}
	}

	for i := range minLen {
		var next T
		for _, iterator := range iters {
			next = append(next, iterator[i])
		}
		zipped = append(zipped, next)
	}

	return zipped
}

func Any[T any](iterator []T, predicate func(t T) bool) bool {
	return slices.ContainsFunc(iterator, predicate)
}

func All[T any](iterator []T, predicate func(t T) bool) bool {
	rVal := true
	for _, item := range iterator {
		rVal = predicate(item)
	}
	return rVal
}

func Sum[T cmp.Ordered | string](iterator []T) T {
	var rVal T
	for _, val := range iterator {
		rVal += val
	}
	return rVal
}

// This is not an OG python function but... why not?
func Product[T Number](iterator []T) T {
	var rVal T
	for _, val := range iterator {
		rVal *= val
	}
	return rVal
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

func Filter[T any](fn func(T) bool, iterator []T) []T {
	result := make([]T, 0)
	for _, v := range iterator {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map[T any, U any](fn func(T) U, iterator []T) []U {
	result := make([]U, len(iterator))
	for i, v := range iterator {
		result[i] = fn(v)
	}
	return result
}

func Reversed(seq []any) []any                                     {}
func Sorted(iterable []any, key func(any) any, reverse bool) []any {}

// Type conversion and creation functions
func Abs[T Number](x T) T {
	return T(math.Abs(float64(x)))
}

func Ascii(obj any) string                                        {}
func Bin(x int) string                                            {}
func Bool(x any) bool                                             {}
func ByteArray(source any, encoding string, errors string) []byte {}
func Bytes(source any, encoding string, errors string) []byte     {}
func Chr(i int) string                                            {}
func Complex(real float64, imag float64) complex128               {}
func Dict(mapping any, kwargs ...any) map[any]any                 {}
func Float(x any) float64                                         {}
func FrozenSet(iterable []any) map[any]struct{}                   {}
func Hex(x int) string                                            {}
func Int(x any, base int) int                                     {}
func Oct(x int) string                                            {}
func Ord(c string) int                                            {}
func Str(obj any, encoding string, errors string) string          {}
func Tuple(iterable []any) []any                                  {}

// Object and attribute functions
func Callable(obj any) bool                              {}
func DelAttr(obj any, name string)                       {}
func Dir(obj any) []string                               {}
func GetAttr(obj any, name string, defaultValue any) any {}
func Globals() map[string]any                            {}
func HasAttr(obj any, name string) bool                  {}
func Hash(obj any) int                                   {}
func Id(obj any) uintptr                                 {}
func IsInstance(obj any, classinfo any) bool             {}
func IsSubClass(class any, classinfo any) bool           {}
func Locals() map[string]any                             {}
func Repr(obj any) string                                {}
func SetAttr(obj any, name string, value any)            {}
func Type(obj any) any                                   {}
func Vars(obj any) map[string]any                        {}

// Iterator and generator functions
func Iter(obj any, sentinel any) any          {}
func Next(iterator any, defaultValue any) any {}

// Math and numeric functions
func Round(number float64, ndigits int) float64 {}

// I/O functions
func Input(prompt string) string {}
func Open(file string, mode string, buffering int, encoding string, errors string, newline string, closefd bool, opener func(string, int) int) any {
}

// Code execution and compilation
func Compile(source string, filename string, mode string, flags int, dontInherit bool, optimize int) any {
}
func Eval(expression string, globals map[string]any, locals map[string]any) any {}
func Exec(source string, globals map[string]any, locals map[string]any)         {}
