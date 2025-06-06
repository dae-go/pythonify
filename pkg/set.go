package py

import (
	"fmt"
	"maps"
	"slices"
)

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](items ...T) Set[T] {
	s := make(Set[T], len(items))
	s.Update(items)
	return s
}

// Adding Elements

func (s *Set[T]) Add(item T) {
	(*s)[item] = struct{}{}
}

func (s *Set[T]) Update(iterators ...any) {

	for _, iterator := range iterators {
		i := Sliced[T](iterator)
		for _, item := range i {
			s.Add(item)
		}
	}
}

// Removing Elements

func (s *Set[T]) Remove(item T) {
	_, ok := (*s)[item]
	if !ok {
		panic(fmt.Sprintf("%v not found", item))
	}
	delete(*s, item)
}

func (s *Set[T]) Delete(item T) { delete(*s, item) }

func (s *Set[T]) Pop() T {

	for item := range *s {
		delete(*s, item)
		return item
	}
	panic("Can't pop from an empty set.")
}

func (s *Set[T]) Clear() { clear(*s) }

func (s *Set[T]) ToSlice() []T { return slices.Collect(maps.Keys(*s)) }

// Set Operations

func (s *Set[T]) Union(other Set[T]) Set[T] {
	result := s.Copy()
	result.Update(other)
}

func (s *Set[T]) Intersection(other Set[T]) Set[T] {}

func (s *Set[T]) Difference(other Set[T]) Set[T] {}

func (s *Set[T]) SymmetricDifference(other Set[T]) Set[T] {}

func (s *Set[T]) IntersectionUpdate(other Set[T]) {}

func (s *Set[T]) DifferenceUpdate(other Set[T]) {}

func (s *Set[T]) SymmetricDifferenceUpdate(other Set[T]) {}

// Set Comparisons

func (s *Set[T]) IsDisjoint(other Set[T]) bool {
	l := s.Union(other)
	return len(l) == len(*s)+len(other)
}

func (s *Set[T]) IsSubset(other Set[T]) bool {
	for val := range *s {
		if !other.Contains(val) {
			return false
		}
	}
	return true
}

func (s *Set[T]) IsSuperset(other Set[T]) bool {
	for val := range other {
		if !s.Contains(val) {
			return false
		}
	}
	return true
}

// Other Methods

func (s *Set[T]) Copy() Set[T] { return maps.Clone(*s) }

func (s *Set[T]) Contains(item T) bool {
	_, ok := (*s)[item]
	return ok
}
