package goadict

func (d Goadict[K, T]) Count() int {
	return len(d)
}
