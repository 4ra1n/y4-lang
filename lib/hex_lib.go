package lib

import (
	"encoding/hex"

	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/native"
)

const (
	HexLibName = "海克斯"
)

var (
	HexLib = base.NewList[*native.NativeFunction]()
)

func init() {
	encodeFun, err := native.NewNativeFunction(HexLibName+SEP+"编码", encodeHex)
	if err != nil {
		return
	}
	decodeFun, err := native.NewNativeFunction(HexLibName+SEP+"解码", decodeHex)
	if err != nil {
		return
	}

	HexLib.Add(encodeFun)
	HexLib.Add(decodeFun)
}

func encodeHex(s interface{}) string {
	b, ok := s.([]byte)
	if !ok {
		log.Errorf("hex lib encode error")
		return "<空的>"
	}
	return hex.EncodeToString(b)
}

func decodeHex(s string) interface{} {
	res, err := hex.DecodeString(s)
	if err != nil {
		log.Errorf("hex lib decode error: %s", err)
		return "<空的>"
	} else {
		return res
	}
}
