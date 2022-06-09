package goadict

// Check if dictionary contains value
func (d Goadict[K, T]) ContainsValue(valueToCheck T) bool {
	for key := range d {
		if d[key] == valueToCheck {
			return true
		}
	}

	return false
}
