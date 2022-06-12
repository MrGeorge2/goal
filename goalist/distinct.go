package goalist

// Remove duplicite items from goalist
func (l Goalist[T]) Distinct() Goalist[T] {
	result := Goalist[T]{}

	for _, item := range l {
		if !result.Contains(item) {
			result.Add(item)
		}
	}

	return result
}
