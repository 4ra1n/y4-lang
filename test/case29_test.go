package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/assert"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase29(t *testing.T) {
	Finish()
	Redirect()
	code := `
#引入 "海克斯"

a = "ACED0005";
c = 海克斯.解码(a);
b = 海克斯.编码(c);

打印(b);
打印(c);
`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "aced0005\n")
	assert.Contains(t, result, "[172 237 0 5]\n")
}
