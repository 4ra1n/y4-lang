package test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase24(t *testing.T) {
	Finish()
	Redirect()
	code := `
#引入 "字符串"

a = "hello world";
b = 字符串.转大写(a);
打印(b);

如果 字符串.有前缀(a, "hello") {
	打印("prefix check pass");
}

如果 字符串.有后缀(a, "rld") {
	打印("suffix check pass");
}

c = 字符串.分割(a, " ");
lc = 长度(c);
打印(c);
i = 0;
当 i < lc {
	打印(c[i]);
	i = i + 1;
}

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "HELLO WORLD\n")
	assert.Contains(t, result, "prefix check pass\n")
	assert.Contains(t, result, "suffix check pass\n")
	assert.Contains(t, result, "[hello world]\n")
	assert.Contains(t, result, "hello\n")
	assert.Contains(t, result, "world\n")
}
