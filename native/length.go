package native

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
)

const nativeLengthFunction = "长度"

func y4Length(v interface{}) int {
	s, isS := v.(string)
	if isS {
		return len(s)
	}
	a, isA := v.([]interface{})
	if isA {
		return len(a)
	}
	as, isAs := v.([]string)
	if isAs {
		return len(as)
	}
	m, isM := v.(*base.Map[string, interface{}])
	if isM {
		return m.Length()
	}
	return envir.FALSE
}
