package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/assert"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase10(t *testing.T) {
	Finish()
	Redirect()
	code := `
打印(1.1);
打印(1+2.2);
打印(1.1+2);
打印(1.1+2.2);
打印(0.1+0.222);

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "1.1\n")
	assert.Contains(t, result, "3.2\n")
	assert.Contains(t, result, "3.1\n")
	assert.Contains(t, result, "3.3000000000000003\n")
	assert.Contains(t, result, "0.322\n")
}
