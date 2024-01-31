package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/4ra1n/y4-lang/assert"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/pre"
)

func TestCase15(t *testing.T) {
	Finish()
	Redirect()
	code := `
#引入 "字符串"
#引入 "my_script.y4"

a = "1";
如果 !字符串.是空(a) {
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

	result := Read()
	assert.Contains(t, result, "from my_script.y4 file\n")
}
