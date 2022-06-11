package goalist

import "sync"

// Returns items that met the predicate
func (l Goalist[T]) Where(predicate func(x T) bool) Goalist[T] {
	var wg sync.WaitGroup

	matchingIndexes := make(chan int, len(l))

	// Selecting items that meet the predicate
	for index, item := range l {
		wg.Add(1)
		go func(parIndex int, parItem T) {
			defer wg.Done()

			if predicate(parItem) {
				matchingIndexes <- parIndex
			}
		}(index, item)
	}

	wg.Wait()
	close(matchingIndexes)

	// Get all indexes to goalist and sort them
	indexes := Goalist[int]{}
	for index := range matchingIndexes {
		indexes.Add(index)
	}
	indexes.Order(func(a, b int) int { return a - b })

	// Add items matching predicate to result
	result := Goalist[T]{}
	for _, index := range indexes {
		result.Add(l[index])
	}

	return result
}
