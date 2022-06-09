package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/tests/seeder"
	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	carList := seeder.CreateShuffledCarList()

	const CAR_IN_LIST = "Skoda"
	firstSkoda := carList.First(func(x seeder.Car) bool { return x.Brand == CAR_IN_LIST })
	assert.NotNil(t, firstSkoda)

	const CAR_NOT_IN_LIST = "Lada"
	firstLada := carList.First(func(x seeder.Car) bool { return x.Brand == CAR_NOT_IN_LIST })
	assert.Nil(t, firstLada)
}
