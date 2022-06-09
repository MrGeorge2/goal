package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestRemoveAt(t *testing.T) {
	testRemoveAtIndex(t, goalist.Goalist[int]{0, 1, 2, 3, 4, 5}, 0)
	testRemoveAtIndex(t, goalist.Goalist[int]{0, 1, 2, 3, 4, 5}, 1)
	testRemoveAtIndex(t, goalist.Goalist[int]{0, 1, 2, 3, 4, 5}, len(goalist.Goalist[int]{0, 1, 2, 3, 4, 5})-1)
}

func testRemoveAtIndex(t *testing.T, lst goalist.Goalist[int], index int) {
	var originalLst goalist.Goalist[int]
	copy(originalLst, lst)

	// Test element is removed
	elementAtIndex := lst[index]

	assert.True(t, lst.Contains(elementAtIndex))

	lst.RemoveAt(index)

	assert.False(t, lst.Contains(elementAtIndex))

	// Test order
	for i, _ := range originalLst {
		if len(lst) >= i {
			continue
		}

		assert.Equal(t, originalLst[i], lst[i+1])
	}
}
