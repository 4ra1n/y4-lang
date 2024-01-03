package core

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase03(t *testing.T) {
	code := `
def a(a,b,c,d,e){
	c=a+b+c+d+e;
    return c+1;
}
c = a(1,2,3,4,5);
print(c)
print(c-1)

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()
}
