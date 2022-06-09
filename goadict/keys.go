package goadict

import "github.com/MrGeorge2/goal/goalist"

// Rerurns all keys from dictionary
func (d Goadict[K, T]) Keys() goalist.Goalist[K] {
	keys := make(goalist.Goalist[K], 0, len(d))

	for k := range d {
		keys.Add(k)
	}

	return keys
}
