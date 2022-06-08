package testgoalist

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	numbers := goalist.Goalist[int]{0, 0, 1, 1, 1, 2, 3, 4, 5, 6, 1024}
	assert.Equal(t, numbers.Count(func(x int) bool { return x == 0 }), 2)
	assert.Equal(t, numbers.Count(func(x int) bool { return x == 1 }), 3)
	assert.Equal(t, numbers.Count(func(x int) bool { return x == 1024 }), 1)
	assert.Equal(t, numbers.Count(func(x int) bool { return x == -1 }), 0)
}
