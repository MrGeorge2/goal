package testgoalist

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestSkipWhileChars(t *testing.T) {
	charList := goalist.Goalist[string]{"H", "E", "L", "L", "O", " ", "W", "O", "R", "L", "D"}
	world := charList.SkipWhile(func(x string) bool {
		return x != " "
	})

	worldCharList := goalist.Goalist[string]{" ", "W", "O", "R", "L", "D"}
	assert.Equal(t, worldCharList, world)

}
