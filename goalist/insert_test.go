package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestPrepend(t *testing.T) {
	const ZERO = 0
	numbers := goalist.Goalist[int]{1, 2, 3, 4, 5}
	insertToIndex(t, 0, numbers, ZERO)
	insertToIndex(t, 1, numbers, ZERO)
	insertToIndex(t, len(numbers)-1, numbers, ZERO)
}

func insertToIndex[T comparable](t *testing.T, index int, lst goalist.Goalist[T], value T) {
	originalLen := len(lst)

	lst.Insert(index, value)

	assert.Len(t, lst, originalLen+1)

	assert.Equal(t, value, lst[index])
}
