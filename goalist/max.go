package goalist

import "golang.org/x/exp/constraints"

func Max[T comparable, N constraints.Ordered](lst Goalist[T], maxFunc func(x T) N) *N {
	if len(lst) == 0 {
		return nil
	}

	maxResult := maxFunc(lst[0])
	for _, item := range lst {
		maxFuncResult := maxFunc(item)

		if maxResult < maxFuncResult {
			maxResult = maxFuncResult
		}
	}

	return &maxResult
}
