package py

type Sliceable[T any] interface {
	ToSlice() []T
}

func Sliced[T any](iterator any) []T {
	result := *new([]T)
	switch v := iterator.(type) {
	case []T:
		result = append(result, v...)
	case Sliceable[T]:
		result = v.ToSlice()
	default:
		panic("unsupported type")
	}
	return result
}
