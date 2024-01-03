package base

import (
	"testing"
)

func TestNewList(t *testing.T) {
	intList := NewList[int]()
	intList.Add(1)
	intList.Add(2)
	if val, _ := intList.Get(0); val != 1 {
		t.Errorf("Expected first element to be 1, got %v", val)
	}
	if length := intList.Length(); length != 2 {
		t.Errorf("Expected length to be 2, got %d", length)
	}
	stringList := NewList[string]()
	stringList.Add("hello")
	stringList.Add("world")
	if val, _ := stringList.Get(1); val != "world" {
		t.Errorf("Expected second element to be 'world', got '%s'", val)
	}
	if length := stringList.Length(); length != 2 {
		t.Errorf("Expected length to be 2, got %d", length)
	}
}
