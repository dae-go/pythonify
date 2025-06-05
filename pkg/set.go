package pkg

import (
	"fmt"
	"maps"
	"slices"
)

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](items ...T) Set[T] {
	s := make(Set[T], len(items))
	s.Update(items...)
	return s
}

// Adding Elements

func (s *Set[T]) Add(item T) {
	(*s)[item] = struct{}{}
}

func (s *Set[T]) Update(items ...T) {
	for _, item := range items {
		s.Add(item)
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
