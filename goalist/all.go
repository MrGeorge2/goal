package goalist

import "context"

// Return true if all values meet the predicate
func (l Goalist[T]) All(predicate func(x T) bool) bool {
	allCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	doneChan := make(chan struct{}, len(l))
	allResultChan := make(chan bool, len(l))

	for _, item := range l {
		go func(ctx context.Context, cancelation context.CancelFunc, predicateItem T) {
			defer func() { doneChan <- struct{}{} }()

			for {
				select {
				case <-ctx.Done():
					allResultChan <- false
					return
				default:
					if !predicate(predicateItem) {
						allResultChan <- false
						cancelation()
					} else {
						allResultChan <- true
					}

					return
				}
			}
		}(allCtx, cancel, item)
	}

	doneCounter := 0
	for {
		select {
		case <-doneChan:
			doneCounter++
			if doneCounter == len(l) {
				return true
			}
		case result := <-allResultChan:
			if !result {
				return false
			}
		}
	}
}
