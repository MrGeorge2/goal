package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	numbers := goalist.Goalist[int]{1, 2, 3, 4, 5}
	reversedNumbers := goalist.Goalist[int]{5, 4, 3, 2, 1}

	numbers.Reverse()

	assert.Equal(t, reversedNumbers, numbers)
}
