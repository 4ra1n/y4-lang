package core

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase05(t *testing.T) {
	code := `
a = 1;
while a<10 {
	a = a+1;
	print(a-1);
}

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()
}
