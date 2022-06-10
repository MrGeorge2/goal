package goalist

import "sync"

// Iterate over the goalist and call the function from parameter on each one
func (l Goalist[T]) ForEach(function func(x T)) {
	for _, item := range l {
		function(item)
	}
}

// Iterate over the goalist and call the function from parameter on each one
func (l Goalist[T]) GoForEach(function func(x T)) {
	var wg sync.WaitGroup

	for _, item := range l {
		wg.Add(1)

		go func(functionParam T) {
			defer wg.Done()

			go function(functionParam)
		}(item)
	}

	wg.Wait()
}
