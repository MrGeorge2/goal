package testgoalist

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestExceptSame(t *testing.T) {
	numbers1 := goalist.Goalist[int]{0, 0, 1, 1, 1, 2, 3, 4, 5, 6, 1024}
	numbers2 := goalist.Goalist[int]{0, 0, 1, 1, 1, 2, 3, 4, 5, 6, 1024}

	excepted := numbers1.Except(numbers2)
	testExcept(t, excepted, numbers1, numbers2)
}

func TestExceptOneSame(t *testing.T) {
	numbers1 := goalist.Goalist[int]{0, 1}
	numbers2 := goalist.Goalist[int]{0}

	excepted := numbers1.Except(numbers2)
	testExcept(t, excepted, numbers1, numbers2)

	assert.Len(t, excepted, 1)
	assert.Equal(t, excepted[0], 1)
}

func testExcept[T comparable](t *testing.T, excepted goalist.Goalist[T], src goalist.Goalist[T], exceptList goalist.Goalist[T]) {
	for _, item := range excepted {
		assert.Equal(t, 1, excepted.Count(func(x T) bool { return x == item }))
		assert.True(t, src.Contains(item))
		assert.False(t, exceptList.Contains(item))
	}
}
