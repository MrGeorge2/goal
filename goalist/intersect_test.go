package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestIntersectSame(t *testing.T) {
	numbers1 := goalist.Goalist[int]{0, 0, 1, 1, 1, 2, 3, 4, 5, 6, 1024}
	numbers2 := goalist.Goalist[int]{0, 0, 1, 1, 1, 2, 3, 4, 5, 6, 1024}

	intersected := numbers1.Intersect(numbers2)

	assert.Equal(t, numbers1, intersected)
	assert.Len(t, intersected, len(numbers1))

	assert.Equal(t, numbers2, intersected)
	assert.Len(t, intersected, len(numbers2))
}

func TestIntersectOneSame(t *testing.T) {
	numbers1 := goalist.Goalist[int]{0, 1}
	numbers2 := goalist.Goalist[int]{0}

	intersected := numbers1.Intersect(numbers2)
	assert.Len(t, intersected, 1)
	assert.Equal(t, numbers2, intersected)
}
