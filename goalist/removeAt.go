package goalist

func (l *Goalist[T]) RemoveAt(index int) {
	(*l)[index] = (*l)[len(*l)-1]
	(*l) = (*l)[:len(*l)-1]
}
