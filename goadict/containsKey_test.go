package goadict_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goadict"
	"github.com/stretchr/testify/assert"
)

func TestContainsKeyExisting(t *testing.T) {
	dict := goadict.Goadict[int, int]{}

	const DICT_LEN = 10

	for i := 0; i < DICT_LEN; i++ {
		dict[i] = i
	}

	for i := 0; i < dict.Count(); i++ {
		assert.True(t, dict.ContainsKey(i))
	}
}

func TestContainsKeyNotExisting(t *testing.T) {
	dict := goadict.Goadict[int, int]{}

	const DICT_LEN = 10

	for i := 0; i < DICT_LEN; i++ {
		dict[i] = i
	}

	assert.False(t, dict.ContainsKey(DICT_LEN+1024))
}
