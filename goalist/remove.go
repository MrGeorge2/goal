package goalist

func (l *Goalist[T]) Remove(predicate func(x T) bool) {
	for _, item := range *l {
		if !predicate(item) {
			continue
		}

		index := l.IndexOf(item)

		l.RemoveAt(*index)
	}
}
