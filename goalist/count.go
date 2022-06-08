package goalist

func (l Goalist[T]) Count(predicate func(x T) bool) int {
	counter := 0

	for _, item := range l {
		if predicate(item) {
			counter++
		}
	}

	return counter
}
