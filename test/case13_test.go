package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/assert"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase13(t *testing.T) {
	Finish()
	Redirect()
	code := `
函数 test(a,b,c) {
	a1=1;
	a2=1;
	a3=1;
	a4=1;
	a5=1;
	a6=1;
	a7=1;
	a8=1;
	a9=1;
	a10=1;
	a11=1;
	a12=1;
	a13=1;
	a1=1;
	a2=1;
	a3=1;
	a4=1;
	a5=1;
	a6=1;
	a7=1;
	a8=1;
	a9=1;
	返回 a1+a2+a+b+c;
}

打印(test(1,2,3));
`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "8\n")
}
