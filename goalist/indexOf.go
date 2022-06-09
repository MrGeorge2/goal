package goalist

// Returns index of first item
// If the item is not in the goalist, then returns nil
func (l Goalist[T]) IndexOf(value T) *int {
	for i, item := range l {
		if item == value {
			return &i
		}
	}
	return nil
}
