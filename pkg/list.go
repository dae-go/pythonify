package py

import (
	"cmp"
	"slices"
)

type Alist[T any] []T

func NewAlist[T any](items ...T) Alist[T] {
	return Alist[T](items)
}

type List[T comparable] Alist[T]

func NewList[T comparable](items ...T) List[T] {
	return List[T](items)
}

type Olist[T cmp.Ordered] List[T]

func NewOlist[T cmp.Ordered](items ...T) Olist[T] {
	return Olist[T](items)
}

func (l *Olist[T]) Sort() { slices.Sort(*l) }

// Adding Elements

func (l *Alist[T]) Append(item T) {
	*l = append(*l, item)
}

func (l *Alist[T]) Extend(item ...T) {
	*l = append(*l, item...)
}

func (l *Alist[T]) Insert(index int, item T) {
	if index < 0 || index > len(*l) {
		index = len(*l)
	}
	buf := append(Alist[T]{item}, (*l)[index:]...)
	*l = append((*l)[:index], buf...)
}

// Removing Elements

func (l *Alist[T]) Pop(index int) T {

	r := (*l)[index]
	*l = append((*l)[:index], (*l)[index+1:]...)
	return r
}

func (l *Alist[T]) Clear() { clear(*l) }

// Modify Elements/List

func (l *Alist[T]) Reverse() { slices.Reverse(*l) }

func (l *Alist[T]) SortFunc(cmp func(a, b T) int) { slices.SortFunc(*l, cmp) }

func (l *Alist[T]) Copy() List[T] {
	newList := make(List[T], len(*l))
	copy(newList, *l)
	return newList
}

func (l Alist[T]) ToSlice() []T {
	return append([]T{}, (l)...)
}

// List only

func (l *List[T]) Index(item T) int {
	return slices.Index(*l, item)
}

func (l *List[T]) Remove(item T) {
	for i, v := range *l {
		if v == item {
			*l = append((*l)[:i], (*l)[i+1:]...)
			return
		}
	}
}
func (l *List[T]) Count(item T) int {
	count := 0
	for _, v := range *l {
		if v == item {
			count++
		}
	}
	return count
}
