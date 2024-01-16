package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase08(t *testing.T) {
	code := `
def test(a) {
	while a<10 {
		a = a+1;
		if a > 5 {
			print("continue");
			continue;
		}
		print(a);
	}
}

test(3);

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()
}
