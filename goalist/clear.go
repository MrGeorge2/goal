package goalist

// Clear the goalist
// Mutates to goalist
func (l *Goalist[T]) Clear() {
	*l = Goalist[T]{}
}
