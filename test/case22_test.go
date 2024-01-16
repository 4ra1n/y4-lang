package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
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
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()
}
