package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/assert"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase23(t *testing.T) {
	Finish()
	Redirect()
	code := `
list = 新列表();
打印(list);

列表添加(list, "test1");
列表添加(list, "test2");
列表添加(list, "test3");
val = 列表获取(list, 0);
打印(val);

arr = 列表转数组(list);

i = 0;
当 i < 长度(arr) {
	打印(arr[i]);
	i = i + 1;
}

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "test1\n")
	assert.Contains(t, result, "test2\n")
	assert.Contains(t, result, "test3\n")
}
