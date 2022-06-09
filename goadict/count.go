package goadict

// Returns number of keys inside dict
func (d Goadict[K, T]) Count() int {
	return len(d)
}
