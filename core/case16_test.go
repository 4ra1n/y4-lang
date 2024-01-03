package core

import (
	"fmt"
	"os"
	"testing"

	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/pre"
)

func TestCase16(t *testing.T) {
	code := `
a = "1";
if a == "1" {
    print("test ==");
}
if a != "2" {
	print("test !=2");
}
if a != "1" {
	print("test !=1");
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
	i := NewInterpreter(l)
	// start
	i.Start()

	err = os.Remove("temp.y4")
	if err != nil {
		fmt.Println(err)
		return
	}
}
