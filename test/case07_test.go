package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase07(t *testing.T) {
	code := `
def test(){
	i = 0;
	while i<10 {
		print(i);
		i=i+1;
	}
}

y4 test();
y4 test();
y4 test();
y4 test();

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()
}
