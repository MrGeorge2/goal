package list

func (l List[T]) Where(predicate func(x T) bool) List[T] {
	result := List[T]{}

	for _, item := range l {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}
