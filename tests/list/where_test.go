package testlist

import (
	"testing"

	"github.com/MrGeorge2/goal/tests/seeder"
	"github.com/stretchr/testify/assert"
)

func TestWhereLen(t *testing.T) {
	const CAR_BRAND = "Skoda"

	carList := seeder.CreateShuffledCarList()

	skodas := carList.Where(func(x seeder.Car) bool { return x.Brand == CAR_BRAND })

	assert.Len(t, skodas, 2)

	for _, car := range skodas {
		assert.Equal(t, CAR_BRAND, car.Brand)
	}
}
