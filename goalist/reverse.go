package goalist

// Reverse list
// reversing goalist.Goalist[int]{1, 2, 3, 4, 5} will result to goalist.Goalist[int]{5, 4, 3, 2, 1}
func (l *Goalist[T]) Reverse() {
	for i, j := 0, len(*l)-1; i < j; i, j = i+1, j-1 {
		(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
	}
}
