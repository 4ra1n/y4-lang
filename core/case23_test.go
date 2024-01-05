package core

import (
	"fmt"
	"os"
	"testing"

	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/pre"
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
	err := os.WriteFile("temp.y4", []byte(code), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	// preprocessor
	ip := pre.NewIncludeProcessor("temp.y4")
	newReader := ip.Process()
	// new lexer
	l := lexer.NewLexer(newReader)
	// new interpreter
	i := NewInterpreter(l, nil)
	// start
	i.Start()

	err = os.Remove("temp.y4")
	if err != nil {
		fmt.Println(err)
		return
	}
}
