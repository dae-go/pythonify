package py

import (
	"cmp"
	"slices"
)

type Olist[T cmp.Ordered] List[T]

func NewOlist[T cmp.Ordered](items ...T) Olist[T] {
	return Olist[T](items)
}

func (l *Olist[T]) Len() int                      { return (*List[T])(l).Len() }
func (l *Olist[T]) Append(item T)                 { (*List[T])(l).Append(item) }
func (l *Olist[T]) Extend(items ...T)             { (*List[T])(l).Extend(items...) }
func (l *Olist[T]) Insert(index int, item T)      { (*List[T])(l).Insert(index, item) }
func (l *Olist[T]) Pop(index int) T               { return (*List[T])(l).Pop(index) }
func (l *Olist[T]) Clear()                        { (*List[T])(l).Clear() }
func (l *Olist[T]) Reverse()                      { (*List[T])(l).Reverse() }
func (l *Olist[T]) SortFunc(cmp func(a, b T) int) { (*List[T])(l).SortFunc(cmp) }
func (l *Olist[T]) Index(item T) int              { return (*List[T])(l).Index(item) }
func (l *Olist[T]) Remove(item T)                 { (*List[T])(l).Remove(item) }
func (l *Olist[T]) Count(item T) int              { return (*List[T])(l).Count(item) }

func (l *Olist[T]) Copy() BasicList[T] {
	newList := make(Olist[T], len(*l))
	copy(newList, *l)
	return &newList
}

func (l Olist[T]) ToSlice() []T {
	return append([]T{}, l...)
}

func (l *Olist[T]) Sort() { slices.Sort(*l) }
