package goadict_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goadict"
	"github.com/stretchr/testify/assert"
)

func TestClear(t *testing.T) {
	dict := goadict.Goadict[int, int]{}

	const DICT_LEN = 10

	for i := 0; i < DICT_LEN; i++ {
		dict[i] = i
	}
	assert.Equal(t, DICT_LEN, dict.Count())

	dict.Clear()

	assert.Zero(t, dict.Count())
}
