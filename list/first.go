package list

func (l List[T]) First(predicate func(x T) bool) *T {
	var result *T = nil

	for _, item := range l {
		if predicate(item) {
			return &item
		}
	}

	return result
}
