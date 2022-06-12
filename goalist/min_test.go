package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestMinEmpty(t *testing.T) {
	empty := goalist.Goalist[int]{}

	minResult := goalist.Min(empty, func(x int) int { return x })

	assert.Nil(t, minResult)
}

func TestMintestNumbers(t *testing.T) {
	testMin(t, goalist.Goalist[int]{1, 2, 3, 4}, 1)
	testMin(t, goalist.Goalist[int]{1, 2, 3, -1}, -1)
	testMin(t, goalist.Goalist[int]{1, -5, 3, 4}, -5)

}

func testMin(t *testing.T, lst goalist.Goalist[int], minValue int) {
	resultMin := goalist.Min(lst, func(x int) int { return x })
	assert.Equal(t, minValue, *resultMin)
}
