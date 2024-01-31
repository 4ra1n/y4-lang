package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/assert"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase03(t *testing.T) {
	Finish()
	Redirect()
	code := `
函数 a(a,b,c,d,e){
	c=a+b+c+d+e;
    返回 c+1;
}
c = a(1,2,3,4,5);
打印(c)
打印(c-1)

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "16")
	assert.Contains(t, result, "15")
}
