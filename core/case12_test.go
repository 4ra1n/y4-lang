package core

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase12(t *testing.T) {
	code := `
def test(a,b,c) {
	return a+b+c;
}
a=1;
b=2.2;
c=3;
print(test(a,b,c));

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()
}
