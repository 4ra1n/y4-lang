package base

import (
	"testing"
)

func TestHashSet(t *testing.T) {
	hashSet := NewHashSet[int]()
	hashSet.Add(1)
	hashSet.Add(2)
	hashSet.Add(3)
	if hashSet.Length() != 3 {
		t.Errorf("Expected length 3, got %d", hashSet.Length())
	}
	if !hashSet.Contains(1) {
		t.Errorf("Expected to contain 1")
	}
	if hashSet.Contains(4) {
		t.Errorf("Expected not to contain 4")
	}
	hashSet.Remove(2)
	if hashSet.Contains(2) {
		t.Errorf("Expected not to contain 2 after removal")
	}
	if hashSet.Length() != 2 {
		t.Errorf("Expected length 2 after removal, got %d", hashSet.Length())
	}
	items := hashSet.Items()
	if len(items) != 2 || !contains(items, 1) || !contains(items, 3) {
		t.Errorf("Items method returned incorrect items: %v", items)
	}
}

func contains[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
