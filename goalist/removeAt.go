package goalist

// Remove item from goalist on certain index
func (l *Goalist[T]) RemoveAt(index int) {
	ret := make([]T, 0)
	ret = append(ret, (*l)[:index]...)
	*l = append(ret, (*l)[index+1:]...)
}
