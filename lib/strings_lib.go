package lib

import (
	"strings"

	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/native"
)

const (
	StringsLibName = "strings"
)

var (
	StringsLib = base.NewList[*native.NativeFunction]()
)

func init() {
	isEmptyFun, err := native.NewNativeFunction(StringsLibName+SEP+"isEmpty", isEmpty)
	if err != nil {
		log.Error("load string lib error")
		return
	}
	StringsLib.Add(isEmptyFun)
}

func isEmpty(data string) int {
	if strings.TrimSpace(data) == "" {
		return envir.TRUE
	} else {
		return envir.FALSE
	}
}
