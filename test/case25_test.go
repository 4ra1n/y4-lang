package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
	"github.com/stretchr/testify/assert"
)

func TestCase25(t *testing.T) {
	Finish()
	Redirect()
	code := `
#引入 "贝斯六四"

a = 贝斯六四.编码("4ra1n");
打印(a);

b = 贝斯六四.解码(a);
打印(b);

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "NHJhMW4=\n")
	assert.Contains(t, result, "4ra1n\n")
}
