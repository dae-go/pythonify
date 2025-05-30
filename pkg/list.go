package pkg

import (
	"cmp"
	"slices"
)

type List[T comparable] []T

// Adding Elements

func (l *List[T]) Append(item T) {
	*l = append(*l, item)
}

func (l *List[T]) Extend(item ...T) {
	*l = append(*l, item...)
}

func (l *List[T]) Insert(index int, item T) {
	if index < 0 || index > len(*l) {
		index = len(*l)
	}
	buf := append([]T{item}, (*l)[index:]...)
	*l = append((*l)[:index], buf...)
}

// Removing Elements

func (l *List[T]) Pop(index int) T {

	r := (*l)[index]
	*l = append((*l)[:index], (*l)[index+1:]...)
	return r
}

func (l *List[T]) Remove(item T) {
	for i, v := range *l {
		if v == item {
			*l = append((*l)[:i], (*l)[i+1:]...)
			return
		}
	}
}

func (l *List[T]) Clear() { clear(*l) }

// Searching Elements

func (l *List[T]) Index(item T) int {
	return slices.Index(*l, item)
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

// Modify Elements/List

func (l *List[T]) Reverse() { slices.Reverse(*l) }

func (l *List[T]) Sort() {
	defer func() {
		r := recover()
		if r != nil {
			panic("List type not sortable")
		}
	}()
	slices.SortFunc(*l, func(a, b T) int {
		if cmp.Ordered(a) {

		}
		return cmp.Compare(cmp.Ordered(a), any(b))
	})
}

func (l *List[T]) Copy() List[T] {
	newList := make(List[T], len(*l))
	copy(newList, *l)
	return newList
}
