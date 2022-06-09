package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	const ZERO = 0
	numbers := goalist.Goalist[int]{ZERO, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numbers.Remove(func(x int) bool { return x == ZERO })

	assert.False(t, numbers.Contains(ZERO))
}

func TestRemoveAll(t *testing.T) {
	numbers := goalist.Goalist[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.NotZero(t, numbers)

	numbers.Remove(func(x int) bool { return true })
	assert.Len(t, numbers, 0)
}

func TestRemoveNotExistingItem(t *testing.T) {
	numbers := goalist.Goalist[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	originalLen := len(numbers)

	numbers.Remove(func(x int) bool { return x == 1024 })

	assert.Len(t, numbers, originalLen)
}
