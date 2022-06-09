package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestEnumerate(t *testing.T) {
	numbers := goalist.Goalist[int]{0, 1, 2, 3, 4, 5}
	enumerator := numbers.GetEnumerator()

	for num := range numbers {
		assert.Equal(t, num, *enumerator())
	}

	assert.Nil(t, enumerator())
}

func TestEnumerateEmptyList(t *testing.T) {
	emptyArray := goalist.Goalist[int]{}
	enumerator := emptyArray.GetEnumerator()

	assert.Nil(t, enumerator())
}
