package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
	"github.com/stretchr/testify/assert"
)

func TestCase08(t *testing.T) {
	Finish()
	Redirect()
	code := `
函数 test(a) {
	当 a<10 {
		a = a+1;
		如果 a > 5 {
			打印("continue");
			继续;
		}
		打印(a);
	}
}

test(3);

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "4\n")
	assert.Contains(t, result, "5\n")
	assert.Contains(t, result, "continue\n")
}
