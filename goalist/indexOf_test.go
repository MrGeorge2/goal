package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestIndexOfExisting(t *testing.T) {
	const ZERO = 0
	numbers := goalist.Goalist[int]{ZERO, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	index := numbers.IndexOf(ZERO)
	assert.NotNil(t, index)
	assert.Equal(t, 0, *index)
}

func TestIndexOfNotExistig(t *testing.T) {
	numbers := goalist.Goalist[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index := numbers.IndexOf(1024)
	assert.Nil(t, index)
}
