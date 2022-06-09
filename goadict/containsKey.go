package goadict

// Returns true if dictionary contains key, else returns false
func (d Goadict[K, T]) ContainsKey(keyToCheck K) bool {
	for key := range d {
		if key == keyToCheck {
			return true
		}
	}

	return false
}
