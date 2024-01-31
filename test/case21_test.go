package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/assert"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase21(t *testing.T) {
	Finish()
	Redirect()
	code := `
map = 新建哈希表();
打印(map);

放入哈希表(map, "key", "value");
val = 从哈希表取出(map, "key");
打印(val);

清空哈希表(map);
val = 从哈希表取出(map, "key");
打印(val);

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "&{{0 0} [] map[]}\n")
	assert.Contains(t, result, "value\n")
	assert.Contains(t, result, "<空的>\n")
}
