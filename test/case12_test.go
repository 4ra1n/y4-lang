package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/assert"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase12(t *testing.T) {
	Finish()
	Redirect()
	code := `
函数 test(a,b,c) {
	返回 a+b+c;
}
a=1;
b=2.2;
c=3;
打印(test(a,b,c));

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "6.2\n")
}
