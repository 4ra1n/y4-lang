package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
	"github.com/stretchr/testify/assert"
)

func TestCase09(t *testing.T) {
	Finish()
	Redirect()
	code := `
函数 test1(a) {
	返回 a+1;
}

函数 test2(b) {
	返回 test1(b)+test1(1);
}

a = test2(999);
打印(a);

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "1002\n")
}
