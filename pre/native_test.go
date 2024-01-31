package pre

import (
	"fmt"
	"os"
	"testing"

	"github.com/4ra1n/y4-lang/assert"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/test"
)

func TestNative(t *testing.T) {
	test.Redirect()
	code := `
#引入 "others.y4"
#引入 "字符串"

a = "test";
如果 !字符串.是空(a) {
	打印("native not empty");
}
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

	result := test.Read()
	assert.Contains(t, result, "native not empty")
}
