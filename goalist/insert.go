package goalist

func (l *Goalist[T]) Insert(index uint, object T) {
	*l = append((*l)[:index+1], (*l)[index:]...)
	(*l)[index] = object
}
