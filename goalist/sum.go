package goalist

import "github.com/MrGeorge2/goal/interfaces"

func (l Goalist[T, N interfaces.Number]) Sum[](predicate func(x T) N) interfaces.Number {
	counter := 0

	for _, item := range l {
		if predicate(item) {
			counter++
		}
	}

	return counter
}
