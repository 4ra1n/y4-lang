package native

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
)

const (
	nativeNewListFunction     = "新列表"
	nativeListAddFunction     = "列表添加"
	nativeListGetFunction     = "列表获取"
	nativeListRemoveFunction  = "列表删除"
	nativeListClearFunction   = "列表清空"
	nativeListToArrayFunction = "列表转数组"
)

func y4NewList() *base.List[interface{}] {
	return base.NewList[interface{}]()
}

func y4ListAdd(l *base.List[interface{}], val interface{}) {
	l.Add(val)
}

func y4ListGet(l *base.List[interface{}], index int) interface{} {
	val, err := l.Get(index)
	if err != nil {
		return "<null>"
	} else {
		return val
	}
}

func y4ListRemove(l *base.List[interface{}], index int) int {
	err := l.Remove(index)
	if err != nil {
		return envir.FALSE
	} else {
		return envir.TRUE
	}
}

func y4ListClear(l *base.List[interface{}]) {
	l.Clear()
}

func y4ListToArray(l *base.List[interface{}]) []interface{} {
	return l.Items()
}
