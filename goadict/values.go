package goadict

import "github.com/MrGeorge2/goal/goalist"

func (d Goadict[K, T]) Values() goalist.Goalist[T] {
	values := make(goalist.Goalist[T], 0, len(d))

	for k := range d {
		values.Add(d[k])
	}

	return values
}
