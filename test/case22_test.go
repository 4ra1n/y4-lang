package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
	"github.com/stretchr/testify/assert"
)

func TestCase22(t *testing.T) {
	Finish()
	Redirect()
	code := `
如果 文件存在("../token.txt") {
	打印("token exist");
}
如果 文件存在("4ra1n.txt") {
	打印("4ra1n exist");
}

写文件("case22.txt", "test");
data = 读文件("case22.txt");
如果 data == "test" {
	打印("test success");
	删除文件("case22.txt");
}

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "test success\n")
}
