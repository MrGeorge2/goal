package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	numbers := goalist.Goalist[int]{1, 2, 3, 4, 5}
	originalLen := len(numbers)

	const ZERO = 0

	numbers.Add(ZERO)

	assert.Len(t, numbers, originalLen+1)

	// Try select added number
	selectedNumber := numbers.First(func(x int) bool { return x == ZERO })

	assert.NotNil(t, selectedNumber)
	assert.Equal(t, ZERO, *selectedNumber)
}
