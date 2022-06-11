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

	// Add items matching predicate to result
	result := Goalist[T]{}
	orderedIndexes := indexes.Order(func(a, b int) int { return a - b })
	for _, index := range orderedIndexes {
		result.Add(l[index])
	}

	return result
}
