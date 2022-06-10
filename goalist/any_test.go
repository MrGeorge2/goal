package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestAnyExists(t *testing.T) {
	numbers := goalist.Goalist[int]{0, 1, 2, 3, 4, 5}
	assert.True(t, numbers.Any(func(x int) bool { return x == 0 }))
}

func TestNotAnyExists(t *testing.T) {
	numbers := goalist.Goalist[int]{0, 1, 2, 3, 4, 5}
	assert.False(t, numbers.Any(func(x int) bool { return x == 1024 }))
}

func TestAnyEmpty(t *testing.T) {
	numbers := goalist.Goalist[int]{}
	assert.False(t, numbers.Any(func(x int) bool { return true }))
}
