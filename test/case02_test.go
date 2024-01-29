package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
	"github.com/stretchr/testify/assert"
)

func TestCase02(t *testing.T) {
	Finish()
	Redirect()

	code := `
a = "hello world";
打印(a);
打印("test case2");
`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "hello world")
	assert.Contains(t, result, "test case2")
}
