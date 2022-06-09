package goalist

// Add goalist to goalist
// Mutate the goalist
func (l *Goalist[T]) AddRange(lst Goalist[T]) {
	*l = append(*l, lst...)
}
