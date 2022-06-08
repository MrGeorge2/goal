package testgoalist

import (
	"testing"

	"github.com/MrGeorge2/goal/tests/seeder"
	"github.com/stretchr/testify/assert"
)

func TestRemoveAt(t *testing.T) {
	cars := seeder.CreateShuffledCarList()

	for _, item := range cars {
		index := cars.IndexOf(item)
		assert.NotNil(t, index)

		cars.RemoveAt(*index)
		assert.False(t, cars.Contains(item))
	}

	assert.Len(t, cars, 0)
}
