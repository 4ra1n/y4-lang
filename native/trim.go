package native

import (
	"strings"
)

const nativeTrimFunction = "trim"

func y4Trim(v interface{}) string {
	if vs, ok := v.(string); !ok {
		return ""
	} else {
		return strings.TrimSpace(vs)
	}
}
