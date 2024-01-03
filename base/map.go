package base

import (
	"sort"
	"sync"
)

type Map[K comparable, V any] struct {
	lock   sync.Mutex
	keys   []K
	values map[K]V
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		values: make(map[K]V),
	}
}

func (m *Map[K, V]) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.keys = make([]K, 0)
	m.values = make(map[K]V)
}

func (m *Map[K, V]) Set(key K, value V) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, ok := m.values[key]; !ok {
		m.keys = append(m.keys, key)
	}
	m.values[key] = value
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	val, ok := m.values[key]
	return val, ok
}

func (m *Map[K, V]) Delete(key K) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, ok := m.values[key]; ok {
		delete(m.values, key)
		for i, k := range m.keys {
			if k == key {
				m.keys = append(m.keys[:i], m.keys[i+1:]...)
				break
			}
		}
	}
}

func (m *Map[K, V]) Keys() []K {
	m.lock.Lock()
	defer m.lock.Unlock()
	return append([]K(nil), m.keys...)
}

func (m *Map[K, V]) Length() int {
	m.lock.Lock()
	defer m.lock.Unlock()
	return len(m.keys)
}

func (m *Map[K, V]) Sort(compare func(a, b K) bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	sort.Slice(m.keys, func(i, j int) bool {
		return compare(m.keys[i], m.keys[j])
	})
}
