package native

import (
	"strings"
)

const nativeTrimFunction = "去空"

func y4Trim(v interface{}) string {
	if vs, ok := v.(string); !ok {
		return ""
	} else {
		return strings.TrimSpace(vs)
	}
}
