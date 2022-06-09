package goadict_test

import (
	"testing"

	"github.com/MrGeorge2/goal/goadict"
	"github.com/stretchr/testify/assert"
)

func TestValues(t *testing.T) {
	dict := goadict.Goadict[string, string]{}

	dict["HELLO"] = "WORLD"
	dict["HEY"] = "HO"

	assert.True(t, dict.ContainsKey("HELLO"))
	assert.True(t, dict.ContainsKey("HEY"))

	assert.False(t, dict.ContainsKey("WORLD"))
	assert.False(t, dict.ContainsKey("HO"))

	assert.True(t, dict.ContainsValue("WORLD"))
	assert.True(t, dict.ContainsValue("HO"))

	assert.False(t, dict.ContainsValue("HELLO"))
	assert.False(t, dict.ContainsValue("HEY"))

}
