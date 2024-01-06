package core

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase01(t *testing.T) {
	code := `
a=[[1,2,3],[1,2]]
print(a[0][1])
`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()
}
