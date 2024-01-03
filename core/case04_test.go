package core

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase04(t *testing.T) {
	code := `
#include "strings"
a = "1"
if a == "" {
	print("a is null native")
} else {
	print("a is not null native")
}
if strings.isEmpty(a) {
	print("a is null strings lib")
} else {
	print("a is not null strings lib")
}

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := NewInterpreter(lexer.NewLexer(r))
	i.Start()
}
