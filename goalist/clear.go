package goalist

func (l *Goalist[T]) Clear() {
	*l = Goalist[T]{}
}
