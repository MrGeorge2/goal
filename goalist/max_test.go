package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestMaxEmpty(t *testing.T) {
	empty := goalist.Goalist[int]{}

	MaxResult := goalist.Max(empty, func(x int) int { return x })

	assert.Nil(t, MaxResult)
}

func TestMaxtestNumbers(t *testing.T) {
	testMax(t, goalist.Goalist[int]{1, 2, 3, 4}, 4)
	testMax(t, goalist.Goalist[int]{50, 2, 3, -1}, 50)
	testMax(t, goalist.Goalist[int]{1, -5, 1024, 4}, 1024)

}

func testMax(t *testing.T, lst goalist.Goalist[int], MaxValue int) {
	resultMax := goalist.Max(lst, func(x int) int { return x })
	assert.Equal(t, MaxValue, *resultMax)
}
