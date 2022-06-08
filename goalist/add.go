package goalist

func (l *Goalist[T]) Add(object T) {
	*l = append(*l, object)
}
