package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	numbers := goalist.Goalist[int]{1, 2, 3, 4, 5}
	const SUM = 15

	resultSum := goalist.Sum(numbers, func(x int) int { return x })

	assert.Equal(t, SUM, *resultSum)
}

func TestWithStruct(t *testing.T) {
	type testStruct struct {
		value int
	}

	numbers := goalist.Goalist[testStruct]{
		testStruct{
			value: 1,
		},
		testStruct{
			value: 2,
		},
		testStruct{
			value: 3,
		},
		testStruct{
			value: 4,
		},

		testStruct{
			value: 5,
		},
	}

	const SUM = 15

	resultSum := goalist.Sum(numbers, func(x testStruct) int {
		return x.value
	})

	assert.Equal(t, SUM, *resultSum)
}

func TestEmptyList(t *testing.T) {
	numbers := goalist.Goalist[int]{}
	resultSum := goalist.Sum(numbers, func(x int) int { return x })
	assert.Nil(t, resultSum)
}
