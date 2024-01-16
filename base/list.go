package base

import (
	"fmt"
	"sync"
)

type List[T any] struct {
	lock  sync.Mutex
	items []T
}

func NewList[T any]() *List[T] {
	return &List[T]{
		lock:  sync.Mutex{},
		items: []T{},
	}
}

func (l *List[T]) Clear() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.items = []T{}
}

func (l *List[T]) Add(item T) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.items = append(l.items, item)
}

func (l *List[T]) Get(index int) (T, error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	var zeroValue T
	if index < 0 || index >= len(l.items) {
		return zeroValue, fmt.Errorf("index out of range")
	}
	return l.items[index], nil
}

func (l *List[T]) Remove(index int) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	if index < 0 || index >= len(l.items) {
		return fmt.Errorf("index out of range")
	}
	l.items = append(l.items[:index], l.items[index+1:]...)
	return nil
}

func (l *List[T]) Length() int {
	l.lock.Lock()
	defer l.lock.Unlock()
	return len(l.items)
}

func (l *List[T]) Items() []T {
	l.lock.Lock()
	defer l.lock.Unlock()
	copiedItems := make([]T, len(l.items))
	copy(copiedItems, l.items)
	return copiedItems
}
