package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	numbers := goalist.Goalist[int]{1, 2, 3, 4, 5}
	assert.True(t, numbers.Contains(1))
}

func TestNotContains(t *testing.T) {
	numbers := goalist.Goalist[int]{1, 2, 3, 4, 5}
	assert.False(t, numbers.Contains(1024))
}
