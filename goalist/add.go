package goalist

// Add item to goalist
// Mutate the goalist
func (l *Goalist[T]) Add(item T) {
	*l = append(*l, item)
}
