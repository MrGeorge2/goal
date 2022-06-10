package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestForEach(t *testing.T) {
	numbers := goalist.Goalist[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numbersToCopy := goalist.Goalist[int]{}

	numbers.ForEach(func(x int) {
		numbersToCopy.Add(x)
	})

	assert.Equal(t, len(numbers), len(numbersToCopy))
	assert.Equal(t, numbers, numbersToCopy)
}

func TestForEachSum(t *testing.T) {
	numbers := goalist.Goalist[int]{}

	const SUM = 15
	const ONE = 1
	for i := 0; i < SUM; i++ {
		numbers.Add(ONE)
	}

	sum := 0
	numbers.ForEach(func(x int) {
		sum += x
	})
}

func TestGoForEachSum(t *testing.T) {
	numbers := goalist.Goalist[int]{}

	const SUM = 15
	const ONE = 1
	for i := 0; i < SUM; i++ {
		numbers.Add(ONE)
	}

	sum := 0
	numbers.GoForEach(func(x int) {
		sum += x
	})
}
