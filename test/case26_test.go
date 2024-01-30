package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase26(t *testing.T) {
	Finish()
	Redirect()
	code := `
#引入 "超文本传输协议"

a = 超文本传输协议.盖特("https://www.baidu.com");
打印(a);

h = 从哈希表取出(a, "headers");
打印(h);
`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()
}
