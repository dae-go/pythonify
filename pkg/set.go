package py

import "fmt"

type Set[T comparable] struct {
	FrozenSet[T]
}

func NewSet[T comparable](items ...T) Set[T] {
	s := Set[T]{
		FrozenSet: FrozenSet[T]{
			value: make(map[T]struct{}),
		},
	}
	for _, item := range items {
		s.value[item] = struct{}{}
	}
	return s
}

func (s *Set[T]) Add(item T) {
	s.value[item] = struct{}{}
}

func (s *Set[T]) Update(others ...Sliceable[T]) {
	for _, other := range others {
		for _, v := range other.ToSlice() {
			s.value[v] = struct{}{}
		}
	}
}

func (s *Set[T]) Remove(item T) {
	if _, ok := s.value[item]; !ok {
		panic(fmt.Sprintf("%v not found", item))
	}
	delete(s.value, item)
}

func (s *Set[T]) Discard(item T) {
	delete(s.value, item)
}

func (s *Set[T]) Pop() T {
	for item := range s.value {
		delete(s.value, item)
		return item
	}
	panic("Can't pop from an empty set.")
}

func (s *Set[T]) Clear() {
	clear(s.value)
}

// In-place set operations
func (s *Set[T]) IntersectionUpdate(others ...Sliceable[T]) {
	if len(others) == 0 {
		return
	}

	otherMaps := make([]map[T]struct{}, len(others))
	for i, other := range others {
		otherMaps[i] = make(map[T]struct{})
		for _, v := range other.ToSlice() {
			otherMaps[i][v] = struct{}{}
		}
	}

	for k := range s.value {
		for _, otherMap := range otherMaps {
			if _, ok := otherMap[k]; !ok {
				delete(s.value, k)
				break
			}
		}
	}
}

func (s *Set[T]) DifferenceUpdate(other Sliceable[T]) {
	for _, v := range other.ToSlice() {
		delete(s.value, v)
	}
}

func (s *Set[T]) SymmetricDifferenceUpdate(other Sliceable[T]) {
	otherMap := make(map[T]struct{})
	for _, v := range other.ToSlice() {
		otherMap[v] = struct{}{}
	}
	for k := range otherMap {
		if _, ok := s.value[k]; ok {
			delete(s.value, k)
		} else {
			s.value[k] = struct{}{}
		}
	}
}

func (s Set[T]) Copy() BasicSet[T] {
	return Set[T]{
		FrozenSet: s.FrozenSet.Copy().(FrozenSet[T]),
	}
}

func (s Set[T]) Union(others ...Sliceable[T]) BasicSet[T] {
	return Set[T]{
		FrozenSet: s.FrozenSet.Union(others...).(FrozenSet[T]),
	}
}

func (s Set[T]) Intersection(others ...Sliceable[T]) BasicSet[T] {
	return Set[T]{
		FrozenSet: s.FrozenSet.Intersection(others...).(FrozenSet[T]),
	}
}

func (s Set[T]) Difference(other Sliceable[T]) BasicSet[T] {
	return Set[T]{
		FrozenSet: s.FrozenSet.Difference(other).(FrozenSet[T]),
	}
}

func (s Set[T]) SymmetricDifference(other Sliceable[T]) BasicSet[T] {
	return Set[T]{
		FrozenSet: s.FrozenSet.SymmetricDifference(other).(FrozenSet[T]),
	}
}
