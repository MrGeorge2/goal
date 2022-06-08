package testgoalist

import (
	"testing"

	"github.com/MrGeorge2/goal/goalist"
	"github.com/stretchr/testify/assert"
)

func TestTakeWhileChars(t *testing.T) {
	charList := goalist.Goalist[string]{"H", "E", "L", "L", "O", " ", "W", "O", "R", "L", "D"}
	hello := charList.TakeWhile(func(x string) bool {
		return x != " "
	})

	helloCharList := goalist.Goalist[string]{"H", "E", "L", "L", "O"}
	assert.Equal(t, helloCharList, hello)

}
