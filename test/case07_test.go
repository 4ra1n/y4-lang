package test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase07(t *testing.T) {
	Finish()
	Redirect()
	code := `
函数 test(a) {
	当 a<10 {
		a = a+1;
		如果 a > 5 {
			跳出;
		}
		打印(a-1);
	}
}

test(1);

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "1\n")
	assert.Contains(t, result, "2\n")
	assert.Contains(t, result, "3\n")
	assert.Contains(t, result, "4\n")
}
