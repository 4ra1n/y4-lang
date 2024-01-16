package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase23(t *testing.T) {
	code := `
list = newList();
print(list);

listAdd(list, "test1");
listAdd(list, "test2");
listAdd(list, "test3");
val = listGet(list, 0);
print(val);

arr = listToArray(list);

i = 0;
while i < length(arr) {
	print(arr[i]);
	i = i + 1;
}

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()
}
