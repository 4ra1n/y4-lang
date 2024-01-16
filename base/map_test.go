package base

import (
	"reflect"
	"sort"
	"testing"
)

func TestNewMap(t *testing.T) {
	omap := NewMap[int, string]()
	omap.Set(1, "one")
	omap.Set(2, "two")
	omap.Set(4, "four")
	omap.Set(6, "six")
	omap.Set(3, "three")
	value, exists := omap.Get(2)
	if !exists || value != "two" {
		t.Errorf("Get(2) = %v, %v; want 'two', true", value, exists)
	}
	omap.Delete(2)
	_, exists = omap.Get(2)
	if exists {
		t.Errorf("Expected 2 to be deleted")
	}
	keys := omap.Keys()
	expectedKeys := []int{1, 3, 4, 6}
	sort.Ints(keys)
	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("Keys = %v; want %v", keys, expectedKeys)
	}
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
	if !reflect.DeepEqual(sortedKeys, expectedSortedKeys) {
		t.Errorf("Sorted Keys = %v; want %v", sortedKeys, expectedSortedKeys)
	}
	for _, key := range sortedKeys {
		value, _ := omap.Get(key)
		t.Logf("%d: %s", key, value)
	}
}
