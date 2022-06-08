package goalist

func (l Goalist[T]) Union(unionList Goalist[T]) Goalist[T] {
	result := l.Distinct()

	for _, item := range l {
		if !l.Contains(item) {
			result.Add(item)
		}
	}
	return result
}
