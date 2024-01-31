package base

import (
	"testing"

	"github.com/4ra1n/y4-lang/assert"
)

func TestHashSet(t *testing.T) {
	y4Assert := assert.New(t)
	hashSet := NewHashSet[int]()
	hashSet.Add(1)
	hashSet.Add(2)
	hashSet.Add(3)
	y4Assert.Equal(3, hashSet.Length())
	y4Assert.Equal(true, hashSet.Contains(1))
	y4Assert.Equal(false, hashSet.Contains(4))
	hashSet.Remove(2)
	y4Assert.Equal(false, hashSet.Contains(2))
	y4Assert.Equal(2, hashSet.Length())
	items := hashSet.Items()
	y4Assert.Equal(2, len(items))
	y4Assert.Equal(true, contains(items, 1))
	y4Assert.Equal(true, contains(items, 3))
}

func contains[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
