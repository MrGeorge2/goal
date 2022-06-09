package goalist

// Return true if any of item from goalist meet the predicate
func (l Goalist[T]) Any(predicate func(x T) bool) bool {
	for _, item := range l {
		if predicate(item) {
			return true
		}
	}

	return false
}
