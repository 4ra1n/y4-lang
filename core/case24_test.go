package core

import (
	"fmt"
	"os"
	"testing"

	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/pre"
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
