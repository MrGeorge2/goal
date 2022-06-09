package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/tests/seeder"
	"github.com/stretchr/testify/assert"
)

func TestPrepend(t *testing.T) {
	insertToIndex(t, 0)
	insertToIndex(t, 1)
	insertToIndex(t, 3)
}

func insertToIndex(t *testing.T, index uint) {
	carList := seeder.CreateShuffledCarList()

	originalCarList := len(carList)

	newCar := seeder.Car{
		Brand: "BMW",
		ID:    1000,
		IsNew: false,
	}

	carList.Insert(index, newCar)

	assert.Len(t, carList, originalCarList+1)

	assert.Equal(t, newCar, carList[index])
}
