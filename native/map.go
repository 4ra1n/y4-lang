package native

import "github.com/4ra1n/y4-lang/base"

const (
	NativeNewMapFunction   = "newMap"
	NativeMapPutFunction   = "mapPut"
	NativeMapGetFunction   = "mapGet"
	NativeMapClearFunction = "mapClear"
)

func y4NewMap() *base.Map[string, interface{}] {
	return base.NewMap[string, interface{}]()
}

func y4MapPut(m *base.Map[string, interface{}], key string, value interface{}) {
	m.Set(key, value)
}

func y4MapGet(m *base.Map[string, interface{}], key string) interface{} {
	val, ok := m.Get(key)
	if ok {
		return val
	}
	return "<null>"
}

func y4MapClear(m *base.Map[string, interface{}]) {
	m.Clear()
}
