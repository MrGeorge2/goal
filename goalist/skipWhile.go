package goalist

// Returns items that are after the predicate is met
func (l Goalist[T]) SkipWhile(predicate func(x T) bool) Goalist[T] {
	result := Goalist[T]{}

	skipping := true

	for _, item := range l {
		if skipping && !predicate(item) {
			skipping = false
		}

		if !skipping {
			result.Add(item)
		}
	}

	return result
}
