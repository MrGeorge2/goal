package testlist

import (
	"testing"

	"github.com/MrGeorge2/goal/tests/seeder"
	"github.com/stretchr/testify/assert"
)

func TestAddRange(t *testing.T) {
	carList := seeder.CreateShuffledCarList()
	carListLen := len(carList)

	listToAppend := seeder.CreateShuffledCarList()
	listAppendLen := len(listToAppend)

	carList.AddRange(listToAppend)
	assert.Len(t, carList, carListLen+listAppendLen)
}
