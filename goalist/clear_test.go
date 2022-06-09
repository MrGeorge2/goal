package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestClear(t *testing.T) {
	numbers := goalist.Goalist[int]{1, 2, 3, 4, 5}
	assert.NotZero(t, len(numbers))

	numbers.Clear()
	assert.Len(t, numbers, 0)
}
