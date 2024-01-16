package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase24(t *testing.T) {
	code := `
#include "strings"

a = "hello world";
b = strings.toUpper(a);
print(b);

if strings.hasPrefix(a, "hello") {
	print("prefix check pass");
}

if strings.hasSuffix(a, "rld") {
	print("suffix check pass");
}

c = strings.split(a, " ");
lc = length(c);
print(c);
i = 0;
while i < lc {
	print(c[i]);
	i = i + 1;
}

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()
}
