package core

import (
	"fmt"
	"os"
	"testing"

	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/pre"
)

func TestCase25(t *testing.T) {
	code := `
#include "base64"

a = base64.encode("4ra1n");
print(a);

b = base64.decode(a);
print(b);

`
	err := os.WriteFile("temp.y4", []byte(code), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	// preprocessor
	ip := pre.NewIncludeProcessor("temp.y4")
	newReader := ip.Process()
	// new lexer
	l := lexer.NewLexer(newReader)
	// new interpreter
	i := NewInterpreter(l, nil)
	// start
	i.Start()

	err = os.Remove("temp.y4")
	if err != nil {
		fmt.Println(err)
		return
	}
}