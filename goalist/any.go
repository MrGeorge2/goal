package goalist

// Return true if all values meet the predicate
func (l Goalist[T]) Any(predicate func(x T) bool) bool {
	if len(l) == 0 {
		return false
	}

	doneChan := make(chan struct{}, len(l))
	anyFoundChan := make(chan struct{})

	for _, item := range l {
		go func(predicateItem T) {
			defer func() {
				doneChan <- struct{}{}
			}()

			if predicate(predicateItem) {
				anyFoundChan <- struct{}{}
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

		case <-anyFoundChan:
			{
				return true
			}

		default:
			{
				if doneCounter >= len(l) {
					return false
				}
			}
		}
	}
}
