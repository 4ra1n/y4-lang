package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/assert"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase01(t *testing.T) {
	Redirect()

	code := `
a=[[1,2,3],[1,2]]
打印(a[0][1])
`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "2")
}
