package native

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
)

const (
	nativeNewListFunction     = "newList"
	nativeListAddFunction     = "listAdd"
	nativeListGetFunction     = "listGet"
	nativeListRemoveFunction  = "listRemove"
	nativeListClearFunction   = "listClear"
	nativeListToArrayFunction = "listToArray"
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
