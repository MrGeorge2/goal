package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestFirstExist(t *testing.T) {
	const ZERO = 0

	numbers := goalist.Goalist[int]{ZERO, 1, 2, 3, 4, 5}
	firstZero := numbers.First(func(x int) bool { return x == ZERO })
	assert.NotNil(t, firstZero)
	assert.Equal(t, ZERO, *firstZero)
}

func TestMultipleExists(t *testing.T) {
	const ZERO = 0

	numbers := goalist.Goalist[int]{ZERO, ZERO, 1, 2, 3, 4, 5}
	firstZero := numbers.First(func(x int) bool { return x == ZERO })
	assert.NotNil(t, firstZero)
	assert.Equal(t, ZERO, *firstZero)
}

func TestFirstNotExists(t *testing.T) {
	const ZERO = 0
	numbers := goalist.Goalist[int]{1, 2, 3, 4, 5}
	firstZero := numbers.First(func(x int) bool { return x == ZERO })
	assert.Nil(t, firstZero)
}
