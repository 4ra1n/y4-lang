package base

import (
	"testing"

	"github.com/4ra1n/y4-lang/assert"
)

func TestNewList(t *testing.T) {
	y4Assert := assert.New(t)

	intList := NewList[int]()
	intList.Add(1)
	intList.Add(2)
	val, _ := intList.Get(0)
	y4Assert.Equal(1, val)

	length := intList.Length()
	y4Assert.Equal(2, length)

	stringList := NewList[string]()
	stringList.Add("hello")
	stringList.Add("world")
	valStr, _ := stringList.Get(1)
	y4Assert.Equal("world", valStr)

	lengthStr := stringList.Length()
	y4Assert.Equal(2, lengthStr)
}
