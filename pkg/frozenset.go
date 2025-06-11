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
	return FrozenSet[T]{value: result}
}

func (s FrozenSet[T]) Contains(item T) bool {
	_, ok := s.value[item]
	return ok
}

func (s FrozenSet[T]) Len() int {
	return len(s.value)
}

func (s FrozenSet[T]) ToSlice() []T {
	return slices.Collect(maps.Keys(s.value))
}

func (s FrozenSet[T]) IsDisjoint(other Sliceable[T]) bool {
	otherSlice := other.ToSlice()
	otherMap := make(map[T]struct{})
	for _, v := range otherSlice {
		otherMap[v] = struct{}{}
	}
	for k := range s.value {
		if _, ok := otherMap[k]; ok {
			return false
		}
	}
	return true
}

func (s FrozenSet[T]) IsSubset(other Sliceable[T]) bool {
	otherSlice := other.ToSlice()
	otherMap := make(map[T]struct{})
	for _, v := range otherSlice {
		otherMap[v] = struct{}{}
	}
	for k := range s.value {
		if _, ok := otherMap[k]; !ok {
			return false
		}
	}
	return true
}

func (s FrozenSet[T]) IsSuperset(other Sliceable[T]) bool {
	otherSlice := other.ToSlice()
	for _, v := range otherSlice {
		if _, ok := s.value[v]; !ok {
			return false
		}
	}
	return true
}

func (s FrozenSet[T]) Union(others ...Sliceable[T]) BasicSet[T] {
	result := make(map[T]struct{})
	for k := range s.value {
		result[k] = struct{}{}
	}
	for _, other := range others {
		for _, v := range other.ToSlice() {
			result[v] = struct{}{}
		}
	}
	return FrozenSet[T]{value: result}
}

func (s FrozenSet[T]) Intersection(others ...Sliceable[T]) BasicSet[T] {
	if len(others) == 0 {
		return s.Copy()
	}
	result := make(map[T]struct{})

	otherMaps := make([]map[T]struct{}, len(others))
	for i, other := range others {
		otherMaps[i] = make(map[T]struct{})
		for _, v := range other.ToSlice() {
			otherMaps[i][v] = struct{}{}
		}
	}

outer:
	for k := range s.value {
		for _, otherMap := range otherMaps {
			if _, ok := otherMap[k]; !ok {
				continue outer
			}
		}
		result[k] = struct{}{}
	}
	return FrozenSet[T]{value: result}
}

func (s FrozenSet[T]) Difference(other Sliceable[T]) BasicSet[T] {
	result := make(map[T]struct{})
	otherMap := make(map[T]struct{})
	for _, v := range other.ToSlice() {
		otherMap[v] = struct{}{}
	}
	for k := range s.value {
		if _, ok := otherMap[k]; !ok {
			result[k] = struct{}{}
		}
	}
	return FrozenSet[T]{value: result}
}

func (s FrozenSet[T]) SymmetricDifference(other Sliceable[T]) BasicSet[T] {
	result := make(map[T]struct{})
	otherMap := make(map[T]struct{})
	for _, v := range other.ToSlice() {
		otherMap[v] = struct{}{}
	}

	for k := range s.value {
		if _, ok := otherMap[k]; !ok {
			result[k] = struct{}{}
		}
	}
	for k := range otherMap {
		if _, ok := s.value[k]; !ok {
			result[k] = struct{}{}
		}
	}
	return FrozenSet[T]{value: result}
}

func (s FrozenSet[T]) Copy() BasicSet[T] {
	result := make(map[T]struct{})
	for k := range s.value {
		result[k] = struct{}{}
	}
	return FrozenSet[T]{value: result}
}
