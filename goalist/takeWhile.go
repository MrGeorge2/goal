package goalist

// Returns items that are before the predicate is met
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
