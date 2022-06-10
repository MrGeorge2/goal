package goalist

import "context"

// Return true if all values meet the predicate
func (l Goalist[T]) Any(predicate func(x T) bool) bool {
	if len(l) == 0 {
		return false
	}

	allCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	doneChan := make(chan struct{}, len(l))
	AnyResultChan := make(chan bool, len(l))

	for _, item := range l {
		go func(ctx context.Context, cancelation context.CancelFunc, predicateItem T) {
			defer func() { doneChan <- struct{}{} }()

			for {
				select {
				case <-ctx.Done():
					AnyResultChan <- false
					return
				default:
					if predicate(predicateItem) {
						AnyResultChan <- true
						cancelation()
					} else {
						AnyResultChan <- false
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
				return false
			}
		case result := <-AnyResultChan:
			if result {
				return true
			}
		}
	}
}
