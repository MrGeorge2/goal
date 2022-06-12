package goalist

import "golang.org/x/exp/constraints"

func Min[T comparable, N constraints.Ordered](lst Goalist[T], minFunc func(x T) N) *N {
	if len(lst) == 0 {
		return nil
	}

	minResult := minFunc(lst[0])
	for _, item := range lst {
		minFuncResult := minFunc(item)

		if minResult > minFuncResult {
			minResult = minFuncResult
		}
	}

	return &minResult
}
