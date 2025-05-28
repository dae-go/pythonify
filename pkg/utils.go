package pkg

func Zip[T []any](iterables ...T) []T {
	if len(iterables) == 0 {
		return []T{}
	}

	var (
		minLen = len(iterables[0]) // updates in next for loop
		zipped = []T{}
	)

	for _, i := range iterables {
		if len(i) < minLen {
			minLen = len(i)
		}
	}

	for i := range minLen {
		var next T
		for _, iterable := range iterables {
			next = append(next, iterable[i])
		}
		zipped = append(zipped, next)
	}

	return zipped
}
