package testlist

import (
	"testing"

	"github.com/MrGeorge2/goal/tests/seeder"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	carList := seeder.CreateShuffledCarList()
	originalLen := len(carList)

	// Try add
	lada := seeder.Car{
		Brand: "Lada",
		ID:    100,
		IsNew: false,
	}

	carList.Add(lada)

	assert.Len(t, carList, originalLen+1)

	// Try select added car
	selectedCar := carList.First(func(x seeder.Car) bool {
		return x == lada
	})

	assert.NotNil(t, selectedCar)
	assert.Equal(t, lada, *selectedCar)
}
