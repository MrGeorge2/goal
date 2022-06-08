package goalist

func (l Goalist[T]) Except(exceptList Goalist[T]) Goalist[T] {
	result := Goalist[T]{}

	for _, item := range l {
		if !exceptList.Contains(item) {
			result.Add(item)
		}
	}

	return result
}
