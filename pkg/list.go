package pkg

type list []any

func (l *list) Append(item any) {
	*l = append(*l, item)
}

func (l *list) Extend(item ...any) {
	*l = append(*l, item...)
}

func (l *list) Pop(index int) any {
	r := (*l)[index]
	*l = append((*l)[:index], (*l)[index+1:]...)
	return r
}
