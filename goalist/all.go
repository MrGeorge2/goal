package goalist

func (l Goalist[T]) All(predicate func(x T) bool) bool {
	for _, item := range l {
		if !predicate(item) {
			return false
		}
	}

	return true
}
