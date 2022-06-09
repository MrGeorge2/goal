package goalist_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/MrGeorge2/goal/tests/seeder"
	"github.com/stretchr/testify/assert"
)

func TestIndexOfExisting(t *testing.T) {
	carList := seeder.CreateShuffledCarList()
	originalLen := len(carList)

	// Existing
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

func TestIndexOfNotExistig(t *testing.T) {
	numbers := goalist.Goalist[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index := numbers.IndexOf(1024)
	assert.Nil(t, index)
}
