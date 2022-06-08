package goalist

func (l Goalist[T]) TakeWhile(predicate func(x T) bool) Goalist[T] {
	result := Goalist[T]{}

	for _, item := range l {
		if !predicate(item) {
			break
		}

		result.Add(item)
	}

	return result
}
