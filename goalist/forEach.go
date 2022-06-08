package goalist

func (l Goalist[T]) ForEach(predicate func(x T)) {
	for _, item := range l {
		predicate(item)
	}
}
