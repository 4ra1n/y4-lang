package core

import (
	"fmt"
	"os"
	"testing"

	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/pre"
)

func TestCase22(t *testing.T) {
	code := `
if existFile("../token.txt") {
	print("token exist");
}
if existFile("4ra1n.txt") {
	print("4ra1n exist");
}

writeFile("case22.txt", "test");
data = readFile("case22.txt");
if data == "test" {
	print("test success");
	deleteFile("case22.txt");
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
	i := NewInterpreter(l, nil)
	// start
	i.Start()

	err = os.Remove("temp.y4")
	if err != nil {
		fmt.Println(err)
		return
	}
}
