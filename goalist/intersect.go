package goalist

// Return items that are in both actual Goalist and Goalist from parameter
func (l Goalist[T]) Intersect(intersectList Goalist[T]) Goalist[T] {
	result := Goalist[T]{}

	for _, item := range l {
		if intersectList.Contains(item) {
			result.Add(item)
		}
	}

	return result
}
