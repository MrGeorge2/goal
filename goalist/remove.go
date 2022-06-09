package goalist

// Remove element according to predicate
// numbers := goalist.Goalist[int]{1, 2 }
// numbers.Remove(func(x int) bool { return x == 1 })
// will mutate list and remove 1
func (l *Goalist[T]) Remove(predicate func(x T) bool) {
	for _, item := range *l {
		if !predicate(item) {
			continue
		}

		index := l.IndexOf(item)

		l.RemoveAt(*index)
	}
}
