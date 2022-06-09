package goalist

func (l *Goalist[T]) Insert(index int, object T) {
	*l = append((*l)[:index+1], (*l)[index:]...)
	(*l)[index] = object
}
