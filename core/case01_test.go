package core

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase01(t *testing.T) {
	code := `
a=-1;
b=2;
c=a+b;
print(a+b);
print(c);
print(-a+b);
`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := NewInterpreter(lexer.NewLexer(r))
	i.Start()
}
