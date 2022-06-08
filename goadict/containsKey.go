package goadict

func (d Goadict[K, T]) ContainsKey(keyToCheck K) bool {
	for key := range d {
		if key == keyToCheck {
			return true
		}
	}

	return false
}
