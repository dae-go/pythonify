package abf

import (
	"cmp"
	"slices"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Filter[T any](fn func(T) bool, iterator []T) []T {
	result := *new([]T)
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

func Reversed(iterable []any) []any {
	result := slices.Clone(iterable)
	slices.Reverse(result)
	return result
}
func Sorted[T cmp.Ordered](iterable []T, reverse bool) []T {
	result := slices.Clone(iterable)
	slices.Sort(result)
	return result
}

func SortedFunc(iterable []any, key func(any, any) int, reverse bool) []any {
	result := slices.Clone(iterable)
	slices.SortFunc(result, key)
	if reverse {
		return Reversed(result)
	}
	return result
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

func Product[T Number](iterator []T) T {
	var rVal T
	for _, val := range iterator {
		rVal *= val
	}
	return rVal
}
