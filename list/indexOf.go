package list

func (l List[T]) IndexOf(value T) *int {
	for i, item := range l {
		if item == value {
			return &i
		}
	}
	return nil
}
