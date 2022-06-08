package testlist

import (
	"testing"

	"github.com/MrGeorge2/goal/tests/seeder"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	carList := seeder.CreateShuffledCarList()

	newCar := seeder.Car{
		Brand: "Volvo",
		IsNew: true,
		ID:    1024,
	}

	carList.Add(newCar)

	assert.True(t, carList.Contains(newCar))
}
