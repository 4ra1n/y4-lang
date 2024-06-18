package lib

import (
	"encoding/base64"

	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/native"
)

const (
	Base64LibName = "贝斯六四"
)

var (
	Base64Lib = base.NewList[*native.NativeFunction]()
)

func init() {
	encodeFun, err := native.NewNativeFunction(Base64LibName+SEP+"编码", encode)
	if err != nil {
		return
	}
	decodeFun, err := native.NewNativeFunction(Base64LibName+SEP+"解码", decode)
	if err != nil {
		return
	}

	Base64Lib.Add(encodeFun)
	Base64Lib.Add(decodeFun)
}

func encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func decode(s string) string {
	res, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Errorf("base64 lib error: %s", err)
		return "<null>"
	} else {
		return string(res)
	}
}
