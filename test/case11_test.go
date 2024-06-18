package test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase11(t *testing.T) {
	Finish()
	Redirect()
	code := `
打印(0.11+0.222+0.333+0.4444);

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "1.1094\n")
}
