package goalist

import (
	"sync"

	"golang.org/x/exp/constraints"
)

func Sum[T comparable, N constraints.Integer | constraints.Float | constraints.Complex](lst Goalist[T], sumFunc func(x T) N) *N {
	if len(lst) == 0 {
		return nil
	}

	sumChan := make(chan N, len(lst))

	var wg sync.WaitGroup
	for _, item := range lst {
		wg.Add(1)

		go func(sumItem T) {
			defer wg.Done()

			sumChan <- sumFunc(sumItem)
		}(item)
	}

	wg.Wait()
	close(sumChan)

	sumResult := *new(N)

	for sumItem := range sumChan {
		sumResult += sumItem
	}

	return &sumResult
}
