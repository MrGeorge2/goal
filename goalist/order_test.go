package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestOrderIntegersAsc(t *testing.T) {
	numbers := goalist.Goalist[int]{4, 2, 5, 1, 3}
	numbersRightOrder := goalist.Goalist[int]{1, 2, 3, 4, 5}

	numbers.Order(func(a, b int) int { return a - b })

	assert.Equal(t, numbersRightOrder, numbers)
}

func TestOrderIntegersDesc(t *testing.T) {
	numbers := goalist.Goalist[int]{4, 2, 5, 1, 3}
	numbersRightOrder := goalist.Goalist[int]{5, 4, 3, 2, 1}

	numbers.Order(func(a, b int) int { return b - a })

	assert.Equal(t, numbersRightOrder, numbers)
}

func TestOrderIntegerNoChange(t *testing.T) {
	numbers := goalist.Goalist[int]{1, 1, 1, 1, 1}
	numbersRightOrder := goalist.Goalist[int]{1, 1, 1, 1, 1}

	numbers.Order(func(a, b int) int { return b - a })

	assert.Equal(t, numbersRightOrder, numbers)
}
