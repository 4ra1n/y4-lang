package core

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase09(t *testing.T) {
	code := `
def test1(a) {
	return a+1;
}

def test2(b) {
	return test1(b)+test1(1);
}

a = test2(999);
print(a);

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()
}
