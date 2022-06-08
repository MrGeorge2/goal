package list

func (l List[T]) Contains(object T) bool {
	return l.Any(func(x T) bool {
		return x == object
	})
}
