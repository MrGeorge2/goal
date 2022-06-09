package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestWhere(t *testing.T) {
	testWhereLen(t, 0)
	testWhereLen(t, 1)
	testWhereLen(t, 1024)
}

func testWhereLen(t *testing.T, count int) {
	const ZERO = 0
	numbers := goalist.Goalist[int]{1, 2, 3, 4, 5}

	for i := 0; i < count; i++ {
		numbers.Add(ZERO)
	}

	zeros := numbers.Where(func(x int) bool { return x == ZERO })
	assert.Len(t, zeros, count)

	for _, item := range zeros {
		assert.Equal(t, ZERO, item)
	}
}
