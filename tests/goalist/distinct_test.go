package testgoalist

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestDistinct(t *testing.T) {
	numbers := goalist.Goalist[int]{0, 0, 1, 2, 3, 1, 2, 3, 4, 100, 2}
	distinctedNumbers := numbers.Distinct()

	for _, num := range distinctedNumbers {
		assert.Equal(t, distinctedNumbers.Count(func(x int) bool { return x == num }), 1)
	}
}
