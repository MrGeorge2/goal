package list

func (l *List[T]) AddRange(lst List[T]) {
	*l = append(*l, lst...)
}
