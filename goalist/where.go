package goalist

// Returns items that met the predicate
func (l Goalist[T]) Where(predicate func(x T) bool) Goalist[T] {
	result := Goalist[T]{}

	for _, item := range l {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}
