package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/tests/seeder"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	carList := seeder.CreateShuffledCarList()

	const CAR_BRAND = "Skoda"
	allSkodas := carList.All(func(x seeder.Car) bool { return x.Brand == CAR_BRAND })
	assert.False(t, allSkodas)

	allHaveBrandName := carList.All(func(x seeder.Car) bool { return len(x.Brand) > 0 })
	assert.True(t, allHaveBrandName)
}
