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
}

func Append(en envir.Environment, name string, fn interface{}) {
	nativeFn, err := NewNativeFunction(name, fn)
	if err != nil {
		log.Error(err)
		return
	}
	en.Put(name, nativeFn)
}
