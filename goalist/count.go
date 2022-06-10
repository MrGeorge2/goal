package goalist

import "sync"

// Get count of items that meet the predicate
func (l Goalist[T]) Count(predicate func(x T) bool) int {
	counter := 0

	var wg sync.WaitGroup
	var lock sync.Mutex

	for _, item := range l {

		wg.Add(1)
		go func(predicateItem T) {
			defer wg.Done()

			if predicate(predicateItem) {
				lock.Lock()
				defer lock.Unlock()

				counter++
			}

		}(item)
	}

	wg.Wait()

	return counter
}
