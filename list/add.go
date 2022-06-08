package list

func (l *List[T]) Add(object T) {
	*l = append(*l, object)
}
