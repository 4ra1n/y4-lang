package native

import "github.com/4ra1n/y4-lang/base"

const (
	nativeNewMapFunction   = "新建哈希表"
	nativeMapPutFunction   = "放入哈希表"
	nativeMapGetFunction   = "从哈希表取出"
	nativeMapClearFunction = "清空哈希表"
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
	return "<空的>"
}

func y4MapClear(m *base.Map[string, interface{}]) {
	m.Clear()
}
