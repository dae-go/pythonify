package py

type Tuple[T comparable] struct {
	value []T
}

func NewTuple[T comparable](value ...T) *Tuple[T] {
	return &Tuple[T]{
		value: append([]T(nil), value...),
	}
}

func (t *Tuple[T]) Count(value T) int {
	count := 0
	for _, elem := range t.value {
		if elem == value {
			count++
		}
	}
	return count
}

func (t *Tuple[T]) Index(value T) int {
	for i, elem := range t.value {
		if elem == value {
			return i
		}
	}
	return -1
}

func (t Tuple[T]) ToSlice() []T {
	result := make([]T, len(t.value))
	copy(result, t.value)
	return result
}

// Tuple not an actual slice so myTuple[x] does not work.
func (t *Tuple[T]) Get(index int) (T, bool) {
	var zero T

	if index < 0 {
		index = len(t.value) + index
	}

	if index >= len(t.value) {
		return zero, false
	}

	return t.value[index], true
}
