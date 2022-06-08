package goalist

func (l *Goalist[T]) AddRange(lst Goalist[T]) {
	*l = append(*l, lst...)
}
