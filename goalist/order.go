package goalist

// Order golist
// Return values of predicate 1, 0, -1
// > 0 sort a after b
// < osort a before b
// 0 keep original order of a and b
func (l *Goalist[T]) Order(predicate func(a, b T) int) {
	sorted := false

	for !sorted {
		moved := false

		for i, item := range *l {
			j := i + 1

			if j >= len(*l) {
				break
			}

			actual := item
			next := (*l)[j]

			orderResult := predicate(actual, next)

			if orderResult == 0 {
				continue

			} else if orderResult > 0 {
				if (*l)[i] != next && (*l)[j] != actual {
					(*l)[i] = next
					(*l)[j] = actual

					moved = true
				}

			}
		}

		sorted = !moved
	}
}
