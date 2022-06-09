package goalist

// Insert value to certain index
func (l *Goalist[T]) Insert(index int, value T) {
	*l = append((*l)[:index+1], (*l)[index:]...)
	(*l)[index] = value
}
