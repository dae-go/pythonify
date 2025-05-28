package pkg

import (
	"fmt"
	"math"
	"slices"
)

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
		for _, iter := range iters {
			next = append(next, iter[i])
		}
		zipped = append(zipped, next)
	}

	return zipped
}

func Any(iter []any, predicate func(t any) bool) bool {
	return slices.ContainsFunc(iter, predicate)
}

func All(iter []any, predicate func(t any) bool) bool {
	rVal := true
	for _, item := range iter {
		rVal = predicate(item)
	}
	return rVal
}

func Sum[T Number | string](iter []T) T {
	var rVal T
	for _, val := range iter {
		rVal += val
	}
	return rVal
}

// This is not an OG python function but... why not?
func Product[T Number](iter []T) T {
	var rVal T
	for _, val := range iter {
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

func Map()    {}
func Filter() {}

func Print(val ...any) {
	fmt.Println(val...) // LOL
}
