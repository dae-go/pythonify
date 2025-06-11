package py

import "slices"

type List[T comparable] Alist[T]

func NewList[T comparable](items ...T) List[T] {
	return List[T](items)
}

func (l *List[T]) Len() int                      { return (*Alist[T])(l).Len() }
func (l *List[T]) Append(item T)                 { (*Alist[T])(l).Append(item) }
func (l *List[T]) Extend(items ...T)             { (*Alist[T])(l).Extend(items...) }
func (l *List[T]) Insert(index int, item T)      { (*Alist[T])(l).Insert(index, item) }
func (l *List[T]) Pop(index int) T               { return (*Alist[T])(l).Pop(index) }
func (l *List[T]) Clear()                        { (*Alist[T])(l).Clear() }
func (l *List[T]) Reverse()                      { (*Alist[T])(l).Reverse() }
func (l *List[T]) SortFunc(cmp func(a, b T) int) { (*Alist[T])(l).SortFunc(cmp) }

func (l *List[T]) Copy() BasicList[T] {
	newList := make(List[T], len(*l))
	copy(newList, *l)
	return &newList
}

func (l List[T]) ToSlice() []T {
	return append([]T{}, l...)
}

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
