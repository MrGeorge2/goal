package goadict_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goadict"
	"github.com/stretchr/testify/assert"
)

func TestValues(t *testing.T) {
	dict := goadict.Goadict[string, string]{}

	const WORLD = "WORLD"
	const HO = "HO"

	dict["HELLO"] = WORLD
	dict["HEY"] = HO

	values := dict.Values()

	assert.True(t, values.Contains(WORLD))
	assert.True(t, values.Contains(HO))
}
