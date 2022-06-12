package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	numbers := goalist.Goalist[int]{1, 2, 3, 4, 5}
	const AVERAGE = 3

	averageResult := goalist.Average(numbers, func(x int) int { return x })

	assert.Equal(t, AVERAGE, *averageResult)
}

func TestAverageNil(t *testing.T) {
	empty := goalist.Goalist[int]{}

	averageResult := goalist.Average(empty, func(x int) int { return x })

	assert.Nil(t, averageResult)
}
