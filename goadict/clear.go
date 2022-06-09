package goadict

// Clears the list
func (d *Goadict[K, T]) Clear() {
	*d = Goadict[K, T]{}
}
