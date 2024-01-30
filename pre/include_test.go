package pre

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/test"
)

func TestInclude(t *testing.T) {
	test.Finish()
	test.Redirect()
	code := `
#引入 "others.y4"

测试1("test");
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
	assert.Contains(t, result, "-9")
	assert.Contains(t, result, "-8")
	assert.Contains(t, result, "-7")
	assert.Contains(t, result, "-6")
	assert.Contains(t, result, "-5")
	assert.Contains(t, result, "-4")
	assert.Contains(t, result, "-3")
	assert.Contains(t, result, "-2")
	assert.Contains(t, result, "-1")
	assert.Contains(t, result, "0")
	assert.Contains(t, result, "1")
}
