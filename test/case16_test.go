package test

import (
	"bytes"
	"testing"
	"time"

	"github.com/4ra1n/y4-lang/assert"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase16(t *testing.T) {
	Finish()
	Redirect()
	code := `
a = "1";
如果 a == "1" {
    打印("test ==");
}
如果 a != "2" {
	打印("test !=2");
}
如果 a != "1" {
	打印("test !=1");
}
`
	log.SetLevel(log.ErrorLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	time.Sleep(1000)

	result := Read()
	assert.Contains(t, result, "test ==\n")
	assert.Contains(t, result, "test !=2")
}
