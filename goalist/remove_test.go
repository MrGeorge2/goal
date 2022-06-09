package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/MrGeorge2/goal/tests/seeder"
	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	cars := seeder.CreateShuffledCarList()

	anySkoda := cars.Any(func(x seeder.Car) bool { return x.Brand == "Skoda" })
	assert.True(t, anySkoda)

	cars.Remove(func(x seeder.Car) bool { return x.Brand == "Skoda" })
	anySkoda = cars.Any(func(x seeder.Car) bool { return x.Brand == "Skoda" })
	assert.False(t, anySkoda)
}

func TestRemoveAll(t *testing.T) {
	cars := seeder.CreateShuffledCarList()
	assert.NotZero(t, cars)

	cars.Remove(func(x seeder.Car) bool { return true })
	assert.Len(t, cars, 0)
}

func TestRemoveNotExistingItem(t *testing.T) {
	numbers := goalist.Goalist[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	originalLen := len(numbers)

	numbers.Remove(func(x int) bool { return x == 1024 })

	assert.Len(t, numbers, originalLen)
}
