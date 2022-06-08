package testgoalist

import (
	"testing"

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
