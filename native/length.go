package native

import "github.com/4ra1n/y4-lang/envir"

const nativeLengthFunction = "length"

func y4Length(v interface{}) int {
	s, isS := v.(string)
	if isS {
		return len(s)
	}
	a, isA := v.([]interface{})
	if isA {
		return len(a)
	}
	return envir.FALSE
}
