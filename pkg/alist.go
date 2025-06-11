package py

import (
	"slices"
)

type Alist[T any] []T

func NewAlist[T any](items ...T) Alist[T] {
	return Alist[T](items)
}

func (l *Alist[T]) Len() int { return len(*l) }

func (l *Alist[T]) Append(item T) {
	*l = append(*l, item)
}

func (l *Alist[T]) Extend(items ...T) {
	*l = append(*l, items...)
}

func (l *Alist[T]) Insert(index int, item T) {
	if index < 0 || index > len(*l) {
		index = len(*l)
	}
	buf := append(Alist[T]{item}, (*l)[index:]...)
	*l = append((*l)[:index], buf...)
}

func (l *Alist[T]) Pop(index int) T {
	r := (*l)[index]
	*l = append((*l)[:index], (*l)[index+1:]...)
	return r
}

func (l *Alist[T]) Clear() { clear(*l) }

func (l *Alist[T]) Reverse() { slices.Reverse(*l) }

func (l *Alist[T]) SortFunc(cmp func(a, b T) int) { slices.SortFunc(*l, cmp) }

func (l *Alist[T]) Copy() BasicList[T] {
	newList := make(Alist[T], len(*l))
	copy(newList, *l)
	return &newList
}

func (l Alist[T]) ToSlice() []T {
	return append([]T{}, l...)
}
