package goalist

func (l Goalist[T]) GetEnumerator() func() *T {
	i := 0

	return func() *T {
		if i >= len(l) {
			return nil
		}

		value := l[i]
		i++
		return &value
	}
}
