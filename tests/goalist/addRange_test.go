package testgoalist

import (
	"testing"

	"github.com/MrGeorge2/goal/tests/seeder"
	"github.com/stretchr/testify/assert"
)

func TestAddRange(t *testing.T) {
	carList := seeder.CreateShuffledCarList()
	carListLen := len(carList)

	goalistToAppend := seeder.CreateShuffledCarList()
	goalistAppendLen := len(goalistToAppend)

	carList.AddRange(goalistToAppend)
	assert.Len(t, carList, carListLen+goalistAppendLen)
}
