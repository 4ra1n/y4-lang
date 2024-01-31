package base

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMap(t *testing.T) {
	y4Assert := assert.New(t)
	omap := NewMap[int, string]()
	omap.Set(1, "one")
	omap.Set(2, "two")
	omap.Set(4, "four")
	omap.Set(6, "six")
	omap.Set(3, "three")
	value, exists := omap.Get(2)
	y4Assert.True(exists)
	y4Assert.Equal("two", value)
	omap.Delete(2)
	_, exists = omap.Get(2)
	y4Assert.False(exists)
	keys := omap.Keys()
	expectedKeys := []int{1, 3, 4, 6}
	sort.Ints(keys)
	y4Assert.True(reflect.DeepEqual(keys, expectedKeys))
	for _, key := range keys {
		value, _ := omap.Get(key)
		t.Logf("%d: %s", key, value)
	}
	t.Log("------------------")
	omap.Sort(func(a, b int) bool {
		return a > b
	})
	sortedKeys := omap.Keys()
	expectedSortedKeys := []int{6, 4, 3, 1}
	y4Assert.True(reflect.DeepEqual(sortedKeys, expectedSortedKeys))
	for _, key := range sortedKeys {
		value, _ := omap.Get(key)
		t.Logf("%d: %s", key, value)
	}
}
