package py

import (
	"fmt"
	"strings"
)

type Tuple[T comparable] struct {
	elements []T
}

func NewTuple[T comparable](elements ...T) *Tuple[T] {
	return &Tuple[T]{
		elements: append([]T(nil), elements...),
	}
}

func (t *Tuple[T]) Count(value T) int {
	count := 0
	for _, elem := range t.elements {
		if elem == value {
			count++
		}
	}
	return count
}

func (t *Tuple[T]) Index(value T) int {
	for i, elem := range t.elements {
		if elem == value {
			return i
		}
	}
	return -1
}

func (t *Tuple[T]) ToSlice() []T {
	result := make([]T, len(t.elements))
	copy(result, t.elements)
	return result
}

func (t *Tuple[T]) String() string {
	if len(t.elements) == 0 {
		return "()"
	}

	var parts []string
	for _, elem := range t.elements {
		parts = append(parts, fmt.Sprintf("%v", elem))
	}

	if len(t.elements) == 1 {
		return fmt.Sprintf("(%s,)", parts[0])
	}

	return fmt.Sprintf("(%s)", strings.Join(parts, ", "))
}
