package list

func (l List[T]) Any(predicate func(x T) bool) bool {
	for _, item := range l {
		if predicate(item) {
			return true
		}
	}

	return false
}
