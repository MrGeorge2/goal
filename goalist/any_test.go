package goalist_test

import (
	"sync"
	"testing"
	"time"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestAnyExists(t *testing.T) {
	const NUMBER_OF_NUMBERS = 100

	numbers := goalist.Goalist[int]{}

	for i := 0; i < NUMBER_OF_NUMBERS; i++ {
		numbers.Add(i)
	}

	assert.Len(t, numbers, NUMBER_OF_NUMBERS)

	var wg sync.WaitGroup
	for _, value := range numbers {
		wg.Add(1)

		go func(parValue int) {
			defer wg.Done()
			assert.True(t, numbers.Any(func(x int) bool {
				time.Sleep(300 * time.Millisecond)
				return x == parValue
			}))
		}(value)
	}

	wg.Wait()
}

func TestNotAnyExists(t *testing.T) {
	numbers := goalist.Goalist[int]{0, 1, 2, 3, 4, 5}

	assert.False(t, numbers.Any(func(x int) bool {
		time.Sleep(300 * time.Millisecond)
		return x == 1024
	}))
}

func TestAnyEmpty(t *testing.T) {
	numbers := goalist.Goalist[int]{}

	assert.False(t, numbers.Any(func(x int) bool {
		time.Sleep(300 * time.Millisecond)
		return true
	}))
}
