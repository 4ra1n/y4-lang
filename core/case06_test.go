package core

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase06(t *testing.T) {
	code := `
def test(a) {
	while a<10 {
		a = a+1;
		if a > 5 {
			break;
		}
		print(a-1);
	}
}

test(1);

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()
}
