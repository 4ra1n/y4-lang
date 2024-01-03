package base

import (
	"sync"
)

type HashSet[T comparable] struct {
	lock sync.Mutex
	set  map[T]struct{}
}

func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{
		set: make(map[T]struct{}),
	}
}

func (h *HashSet[T]) Add(item T) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.set[item] = struct{}{}
}

func (h *HashSet[T]) Contains(item T) bool {
	h.lock.Lock()
	defer h.lock.Unlock()
	_, exists := h.set[item]
	return exists
}

func (h *HashSet[T]) Remove(item T) {
	h.lock.Lock()
	defer h.lock.Unlock()
	delete(h.set, item)
}

func (h *HashSet[T]) Length() int {
	h.lock.Lock()
	defer h.lock.Unlock()
	return len(h.set)
}

func (h *HashSet[T]) Items() []T {
	h.lock.Lock()
	defer h.lock.Unlock()
	items := make([]T, 0, len(h.set))
	for key := range h.set {
		items = append(items, key)
	}
	return items
}
