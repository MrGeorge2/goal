package testgoalist

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestOrderIntegers(t *testing.T) {
	numbers := goalist.Goalist[int]{4, 2, 5, 1, 3}
	numbersRightOrder := goalist.Goalist[int]{1, 2, 3, 4, 5}

	ordered := numbers.Order(func(a, b int) int { return a - b })

	assert.Equal(t, numbersRightOrder, ordered)
}
