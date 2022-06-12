package goalist

// Return true if all values meet the predicate
func (l Goalist[T]) Alll(predicate func(x T) bool) bool {
	if len(l) == 0 {
		return false
	}

	for _, item := range l {
		if !predicate(item) {
			return false
		}
	}

	return true
}

// Return true if all values meet the predicate
func (l Goalist[T]) All(predicate func(x T) bool) bool {
	if len(l) == 0 {
		return false
	}

	doneChan := make(chan struct{}, len(l))
	allNotFound := make(chan struct{})

	for _, item := range l {
		go func(predicateItem T) {
			defer func() {
				doneChan <- struct{}{}
			}()

			if !predicate(predicateItem) {
				allNotFound <- struct{}{}
			}
		}(item)
	}

	doneCounter := 0
	for {
		select {
		case <-doneChan:
			{
				doneCounter++
			}
		case <-allNotFound:
			{
				return false
			}
		default:
			{
				if doneCounter >= len(l) {
					return true
				}
			}
		}
	}
}
