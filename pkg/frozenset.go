package py

import (
	"maps"
	"slices"
)

type FrozenSet[T comparable] struct {
	value map[T]struct{}
}

func NewFrozenSet[T comparable](items ...T) FrozenSet[T] {
	result := make(map[T]struct{})
	for _, v := range items {
		result[v] = struct{}{}
	}
	return FrozenSet[T]{
		value: result,
	}
}

// Set Comparisons

func (s *Set[T]) IsDisjoint(other Set[T]) bool {
	l := s.Union(other)
	return len(l) == len(s.value)+len(other.value)
}

func (s *Set[T]) IsSubset(other Set[T]) bool {
	for val := range s.value {
		if !other.Contains(val) {
			return false
		}
	}
	return true
}

func (s *Set[T]) IsSuperset(other Set[T]) bool {
	for val := range other.value {
		if !s.Contains(val) {
			return false
		}
	}
	return true
}

// Other Methods

func (s *Set[T]) Copy() Set[T] { return NewSet(s.ToSlice()...) }

func (s *Set[T]) Contains(item T) bool {
	_, ok := s.value[item]
	return ok
}

func (s Set[T]) ToSlice() []T { return slices.Collect(maps.Keys(s.value)) }

func (s *Set[T]) Union(other ...Sliceable[T]) Set[T] {
	result := s.Copy()
	result.Update(other...)
	return result
}

func (s *Set[T]) Intersection(other Set[T]) Set[T] {
	result := s.Copy()
	result.IntersectionUpdate(other)
	return result
}

func (s *Set[T]) Difference(other Set[T]) Set[T] {
	result := s.Copy()
	result.DifferenceUpdate(other)
	return result
}
