package core

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestErrCase01(t *testing.T) {
	code := `
a=;
`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := NewInterpreter(lexer.NewLexer(r))
	i.Start()
}
