package goalist

func (l Goalist[T]) Union(unionList Goalist[T]) Goalist[T] {
	result := l.Distinct()

	for _, item := range l {
		if !result.Contains(item) {
			result.Add(item)
		}
	}

	for _, item := range unionList {
		if !result.Contains(item) {
			result.Add(item)
		}
	}

	return result
}
