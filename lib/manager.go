package lib

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/native"
)

const SEP = "."

var (
	LibMap = base.NewMap[string,
		*base.List[*native.NativeFunction]]()
)

func init() {
	LibMap.Set(StringsLibName, StringsLib)
	LibMap.Set(Base64LibName, Base64Lib)
}

func AddLib(libName string, en envir.Environment) {
	ll, exist := LibMap.Get(libName)
	if exist {
		for _, function := range ll.Items() {
			if en.Get(function.GetName()) != nil {
				continue
			}
			en.PutNew(function.GetName(), function)
		}
	}
}
