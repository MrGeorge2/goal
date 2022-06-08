package goalist

func (l Goalist[T]) Intersect(intersectList Goalist[T]) Goalist[T] {
	result := Goalist[T]{}

	for _, item := range l {
		if intersectList.Contains(item) {
			result.Add(item)
		}
	}

	return result
}
