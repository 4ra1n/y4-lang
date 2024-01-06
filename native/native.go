package native

import (
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/log"
)

type Native struct {
	en envir.Environment
}

func NewNative(e envir.Environment) *Native {
	return &Native{
		en: e,
	}
}

func (n *Native) Environment() envir.Environment {
	AppendNatives(n.en)
	return n.en
}

func AppendNatives(en envir.Environment) {
	Append(en, nativePrintFunction, y4Print)
	Append(en, nativeTypeFunction, y4Type)
	Append(en, nativeTrimFunction, y4Trim)
	Append(en, nativeLengthFunction, y4Length)
	Append(en, nativeNowFunction, y4Now)
	Append(en, nativeSleepFunction, y4Sleep)
	// List
	Append(en, nativeNewListFunction, y4NewList)
	Append(en, nativeListAddFunction, y4ListAdd)
	Append(en, nativeListGetFunction, y4ListGet)
	Append(en, nativeListClearFunction, y4ListClear)
	Append(en, nativeListRemoveFunction, y4ListRemove)
	Append(en, nativeListToArrayFunction, y4ListToArray)
	// HashMap
	Append(en, nativeNewMapFunction, y4NewMap)
	Append(en, nativeMapPutFunction, y4MapPut)
	Append(en, nativeMapGetFunction, y4MapGet)
	Append(en, nativeMapClearFunction, y4MapClear)
	// File
	Append(en, nativeExistFileFunction, y4ExistFile)
	Append(en, nativeReadFileFunction, y4ReadFile)
	Append(en, nativeWriteFileFunction, y4WriteFile)
	Append(en, nativeDeleteFileFunction, y4DeleteFile)
	// Convert
	Append(en, nativeIntFunction, y4Int)
}

func Append(en envir.Environment, name string, fn interface{}) {
	nativeFn, err := NewNativeFunction(name, fn)
	if err != nil {
		log.Error(err)
		return
	}
	en.Put(name, nativeFn)
}
