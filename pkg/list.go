package pkg

type list []any

// Adding Elements

func (l *list) Append(item any) {
	*l = append(*l, item)
}

func (l *list) Extend(item ...any) {
	*l = append(*l, item...)
}

func (l *list) Insert(index int, item any) {
	if index < 0 || index > len(*l) {
		index = len(*l)
	}
	buf := append([]any{item}, (*l)[index:]...)
	*l = append((*l)[:index], buf...)
}

// Removing Elements

func (l *list) Pop(index int) any {

	r := (*l)[index]
	*l = append((*l)[:index], (*l)[index+1:]...)
	return r
}

func (l *list) Remove(item any) {
	for i, v := range *l {
		if v == item {
			*l = append((*l)[:i], (*l)[i+1:]...)
			return
		}
	}
}

func (l *list) Clear() {
	*l = (*l)[:0]
}

// Searching Elements

func (l *list) Index(item any) int {
	for i, v := range *l {
		if v == item {
			return i
		}
	}

	panic("Item not found in list")
}

func (l *list) Count(item any) int {
	count := 0
	for _, v := range *l {
		if v == item {
			count++
		}
	}
	return count
}

// Modify Elements/List

func (l *list) Reverse() {

}

func (l *list) Sort() {

}

func (l *list) Copy() list {
	newList := make(list, len(*l))
	copy(newList, *l)
	return newList
}
