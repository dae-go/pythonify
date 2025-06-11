package py

import "cmp"

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type Sliceable[T any] interface {
	ToSlice() []T
}

type BasicList[T any] interface {
	Len() int

	Append(item T)
	Extend(items ...T)
	Insert(index int, item T)

	Pop(index int) T
	Clear()

	Reverse()
	SortFunc(cmp func(a, b T) int)
	Copy() BasicList[T]
	ToSlice() []T
}

type ComparableList[T comparable] interface {
	BasicList[T]

	Index(item T) int
	Remove(item T)
	Count(item T) int
}

type OrderedList[T cmp.Ordered] interface {
	ComparableList[T]

	Sort()
}

type BasicSet[T comparable] interface {
	Sliceable[T]

	Len() int
	Contains(item T) bool
	Copy() BasicSet[T]

	IsDisjoint(other Sliceable[T]) bool
	IsSubset(other Sliceable[T]) bool
	IsSuperset(other Sliceable[T]) bool

	Union(others ...Sliceable[T]) BasicSet[T]
	Intersection(others ...Sliceable[T]) BasicSet[T]
	Difference(other Sliceable[T]) BasicSet[T]
	SymmetricDifference(other Sliceable[T]) BasicSet[T]
}

type MutableSet[T comparable] interface {
	BasicSet[T]

	Add(item T)
	Remove(item T)
	Discard(item T)
	Pop() T
	Clear()
	Update(others ...Sliceable[T])

	IntersectionUpdate(others ...Sliceable[T])
	DifferenceUpdate(other Sliceable[T])
	SymmetricDifferenceUpdate(other Sliceable[T])
}
