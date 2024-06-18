package test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase04(t *testing.T) {
	Finish()
	Redirect()
	code := `
#引入 "字符串"
a = "1"
如果 a == "" {
	打印("a is null native")
} 另外 {
	打印("a is not null native")
}
如果 字符串.是空(a) {
	打印("a is null strings lib")
} 另外 {
	打印("a is not null strings lib")
}

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "a is not null native")
	assert.Contains(t, result, "a is not null strings lib")
}
