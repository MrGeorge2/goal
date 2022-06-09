package goadict_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goadict"
	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	dict := goadict.Goadict[int, int]{}

	const DICT_LEN = 10

	for i := 0; i < DICT_LEN; i++ {
		dict[i] = i
	}

	keys := dict.Keys()
	assert.Len(t, keys, DICT_LEN)

	for i, _ := range keys {
		assert.True(t, keys.Contains(i))
	}
}
