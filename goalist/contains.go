package goalist

// Returns true if goalist contains the item else return false
func (l Goalist[T]) Contains(item T) bool {
	return l.Any(func(x T) bool {
		return x == item
	})
}
