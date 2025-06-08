package py

import (
	"fmt"
	"maps"
	"slices"
)

type Set[T comparable] struct {
	value map[T]struct{}
}

func NewSet[T comparable](items ...T) Set[T] {
	s := Set[T]{}
	i := List[T](items)
	s.Update(i)
	return s
}

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

// Adding Elements

func (s *Set[T]) Add(item T) {
	s.value[item] = struct{}{}
}

func (s *Set[T]) Update(iterators ...Sliceable[T]) {

	for _, iterator := range iterators {
		i := iterator.ToSlice()
		for _, item := range i {
			s.Add(item)
		}
	}
}

// Removing Elements

func (s *Set[T]) Remove(item T) {
	_, ok := s.value[item]
	if !ok {
		panic(fmt.Sprintf("%v not found", item))
	}
	delete(s.value, item)
}

func (s *Set[T]) Discard(item T) { delete(s.value, item) }

func (s *Set[T]) Pop() T {

	for item := range s.value {
		delete(s.value, item)
		return item
	}
	panic("Can't pop from an empty set.")
}

func (s *Set[T]) Clear() { clear(s.value) }

// Set Operations

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

func (s *Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	result := s.Copy()
	result.SymmetricDifferenceUpdate(other)
	return result
}

func (s *Set[T]) IntersectionUpdate(other ...Set[T]) {
	if len(other) == 0 {
		return
	}
	for k := range s.value {
		for _, o := range other {
			if !o.Contains(k) {
				s.Remove(k)
				break
			}
		}
	}
}

func (s *Set[T]) DifferenceUpdate(other Set[T]) {
	for k := range s.value {
		_, ok := other.value[k]
		if ok {
			s.Remove(k)
		}
	}
}

func (s *Set[T]) SymmetricDifferenceUpdate(other Set[T]) {
	for k := range s.value {
		_, ok := other.value[k]
		if !ok {
			s.Remove(k)
		}
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
