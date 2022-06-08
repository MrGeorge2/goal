package testlist

import (
	"testing"

	"github.com/MrGeorge2/goal/tests/seeder"
	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	carList := seeder.CreateShuffledCarList()
	originalLen := len(carList)

	newCar := seeder.Car{
		Brand: "Volvo",
		IsNew: true,
		ID:    1024,
	}

	carList.Add(newCar)

	index := carList.IndexOf(newCar)
	assert.NotNil(t, index)
	assert.Equal(t, originalLen, *index)
}
