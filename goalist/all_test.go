package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestAllExists(t *testing.T) {
	numbers := goalist.Goalist[int]{1, 1, 1, 1}
	assert.True(t, numbers.All(func(x int) bool { return x == 1 }))
}

func TestAllNotExisis(t *testing.T) {
	numbers := goalist.Goalist[int]{1, 1, 1, 1}
	assert.False(t, numbers.All(func(x int) bool { return x == 2 }))
}
