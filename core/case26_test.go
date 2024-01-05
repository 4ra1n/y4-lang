package core

import (
	"fmt"
	"os"
	"testing"

	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/pre"
)

func TestCase26(t *testing.T) {
	code := `
#include "http"

a = http.doGet("https://www.baidu.com");
print(a);

h = mapGet(a, "headers");
print(h);
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
