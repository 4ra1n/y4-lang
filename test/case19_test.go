package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/assert"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase19(t *testing.T) {
	Finish()
	Redirect()
	code := `
arr = ["hello","world",1,2,3,1,2,3];
打印(长度(arr));
打印(长度(arr[0]));
`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "8\n")
	assert.Contains(t, result, "5\n")
}
