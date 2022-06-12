package goalist

import "golang.org/x/exp/constraints"

func Average[T comparable, N constraints.Integer | constraints.Float](lst Goalist[T], averageFunc func(x T) N) *N {
	if len(lst) == 0 {
		return nil
	}

	sum := Sum(lst, averageFunc)

	if sum == nil {
		return nil
	}

	count := N(len(lst))

	average := *sum / count
	return &average
}
