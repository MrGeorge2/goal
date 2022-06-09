package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestAddRange(t *testing.T) {
	numbers1 := goalist.Goalist[int]{1, 2, 3, 4, 5}
	numbers1Len := len(numbers1)

	numbers2 := goalist.Goalist[int]{6, 7, 8, 9, 10}
	numbers2len := len(numbers2)

	numbers1.AddRange(numbers2)
	assert.Len(t, numbers1, numbers1Len+numbers2len)

	for _, item := range numbers2 {
		assert.True(t, numbers1.Contains(item))
	}
}
