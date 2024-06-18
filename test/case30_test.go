package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/conf"
	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase30(t *testing.T) {
	conf.TestConfig()
	code := `
函数 测试(数字) {
	打印(数字)
}

函数 主函数() {
	循环 数字=0; 数字<10; 数字=数字+1 {
		启动 测试(数字);
	}
}
`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()
}
