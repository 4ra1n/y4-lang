package lib

import (
	"strings"

	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/native"
)

const (
	StringsLibName = "字符串"
)

var (
	StringsLib = base.NewList[*native.NativeFunction]()
)

func init() {
	isEmptyFun, err := native.NewNativeFunction(StringsLibName+SEP+"是空", isEmpty)
	if err != nil {
		return
	}
	notEmptyFun, err := native.NewNativeFunction(StringsLibName+SEP+"notEmpty", notEmpty)
	if err != nil {
		return
	}
	containsFun, err := native.NewNativeFunction(StringsLibName+SEP+"contains", contains)
	if err != nil {
		return
	}
	hasPrefixFun, err := native.NewNativeFunction(StringsLibName+SEP+"hasPrefix", hasPrefix)
	if err != nil {
		return
	}
	hasSuffixFun, err := native.NewNativeFunction(StringsLibName+SEP+"hasSuffix", hasSuffix)
	if err != nil {
		return
	}
	toLowerFun, err := native.NewNativeFunction(StringsLibName+SEP+"toLower", toLower)
	if err != nil {
		return
	}
	toUpperFun, err := native.NewNativeFunction(StringsLibName+SEP+"toUpper", toUpper)
	if err != nil {
		return
	}
	splitFun, err := native.NewNativeFunction(StringsLibName+SEP+"split", split)
	if err != nil {
		return
	}
	replaceFun, err := native.NewNativeFunction(StringsLibName+SEP+"replace", replace)
	if err != nil {
		return
	}

	StringsLib.Add(isEmptyFun)
	StringsLib.Add(notEmptyFun)
	StringsLib.Add(containsFun)
	StringsLib.Add(hasPrefixFun)
	StringsLib.Add(hasSuffixFun)
	StringsLib.Add(toLowerFun)
	StringsLib.Add(toUpperFun)
	StringsLib.Add(splitFun)
	StringsLib.Add(replaceFun)
}

func isEmpty(data string) int {
	if strings.TrimSpace(data) == "" {
		return envir.TRUE
	} else {
		return envir.FALSE
	}
}

func notEmpty(data string) int {
	if isEmpty(data) == envir.TRUE {
		return envir.FALSE
	} else {
		return envir.TRUE
	}
}

func contains(data string, c string) int {
	if strings.Contains(data, c) {
		return envir.TRUE
	} else {
		return envir.FALSE
	}
}

func hasPrefix(data string, s string) int {
	if strings.HasPrefix(data, s) {
		return envir.TRUE
	} else {
		return envir.FALSE
	}
}

func hasSuffix(data string, s string) int {
	if strings.HasSuffix(data, s) {
		return envir.TRUE
	} else {
		return envir.FALSE
	}
}

func toLower(data string) string {
	return strings.ToLower(data)
}

func toUpper(data string) string {
	return strings.ToUpper(data)
}

func split(data string, sp string) []string {
	return strings.Split(data, sp)
}

func replace(data string, old string, new string) string {
	return strings.ReplaceAll(data, old, new)
}
