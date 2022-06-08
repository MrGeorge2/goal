package testgoalist

import (
	"testing"

	"github.com/MrGeorge2/goal/tests/seeder"
	"github.com/stretchr/testify/assert"
)

func TestClear(t *testing.T) {
	carList := seeder.CreateShuffledCarList()
	assert.NotZero(t, len(carList))

	carList.Clear()
	assert.Len(t, carList, 0)
}
