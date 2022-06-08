package list

func (l List[T]) All(predicate func(x T) bool) bool {
	for _, item := range l {
		if !predicate(item) {
			return false
		}
	}

	return true
}
