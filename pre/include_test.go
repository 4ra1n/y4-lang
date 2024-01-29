package pre

import (
	"fmt"
	"os"
	"testing"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
)

func TestInclude(t *testing.T) {
	code := `
a=1;
b=2;
打印(a+b);
`
	err := os.WriteFile("temp.y4", []byte(code), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	ip := NewIncludeProcessor("temp.y4")
	data := ip.Process()

	err = os.Remove("temp.y4")
	if err != nil {
		fmt.Println(err)
		return
	}

	l := lexer.NewLexer(data)
	// new interpreter
	i := core.NewInterpreter(l, nil)
	// start
	i.Start()
}
