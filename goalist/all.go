package goalist

// Return true if all values meet the predicate
func (l Goalist[T]) All(predicate func(x T) bool) bool {
	for _, item := range l {
		if !predicate(item) {
			return false
		}
	}

	return true
}
