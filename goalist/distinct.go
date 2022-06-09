package goalist

// Remove duplicite items from goalist
func (l Goalist[T]) Distinct() Goalist[T] {
	result := Goalist[T]{}

	for _, item := range l {
		duplicates := l.Where(func(x T) bool { return x == item })

		if len(duplicates) > 1 {
			// remove all elements except first one
			duplicates.Remove(func(x T) bool { return x != duplicates[0] })
		}
	}

	return result
}
