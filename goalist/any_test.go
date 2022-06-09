package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/tests/seeder"
	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	carList := seeder.CreateShuffledCarList()

	const CAR_IN_LIST = "Skoda"
	hasSkoda := carList.Any(func(x seeder.Car) bool { return x.Brand == CAR_IN_LIST })
	assert.True(t, hasSkoda)

	const CAR_NOT_IN_LIST = "Lada"
	hasLada := carList.Any(func(x seeder.Car) bool { return x.Brand == CAR_NOT_IN_LIST })
	assert.False(t, hasLada)
}
