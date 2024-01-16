package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/pre"
)

func TestCase15(t *testing.T) {
	code := `
#include "strings"
#include "my_script.y4"

a = "1";
if !strings.isEmpty(a) {
    test();
}
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
	i := core.NewInterpreter(l, nil)
	// start
	i.Start()

	//err = os.Remove("temp.y4")
	if err != nil {
		fmt.Println(err)
		return
	}
}
