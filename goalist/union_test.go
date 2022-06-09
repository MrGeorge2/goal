package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestUnionSame(t *testing.T) {
	numbers1 := goalist.Goalist[int]{0, 0, 1, 1, 1, 2, 3, 4, 5, 6, 1024}
	numbers2 := goalist.Goalist[int]{0, 0, 1, 1, 1, 2, 3, 4, 5, 6, 1024}

	united := numbers1.Union(numbers2)
	testUnited(t, united, numbers1, numbers2)
}

func TestUnionOneSame(t *testing.T) {
	numbers1 := goalist.Goalist[int]{0, 1}
	numbers2 := goalist.Goalist[int]{0}

	united := numbers1.Intersect(numbers2)
	testUnited(t, united, numbers1, numbers2)

	assert.Len(t, united, 1)
}

func TestUnionEmptyFirst(t *testing.T) {
	numbers1 := goalist.Goalist[int]{}
	numbers2 := goalist.Goalist[int]{0, 0, 1, 1, 1, 2, 3, 4, 5, 6, 1024}

	united := numbers1.Union(numbers2)
	testUnited(t, united, numbers1, numbers2)
}

func testUnited[T comparable](t *testing.T, united goalist.Goalist[T], goalist1 goalist.Goalist[T], goalist2 goalist.Goalist[T]) {
	for _, item := range united {
		assert.Equal(t, 1, united.Count(func(x T) bool { return x == item }))

		// item in any of two lists
		assert.True(t, goalist1.Contains(item) || goalist2.Contains(item))
	}
}
