package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase21(t *testing.T) {
	code := `
map = newMap();
print(map);

mapPut(map, "key", "value");
val = mapGet(map, "key");
print(val);

mapClear(map);
val = mapGet(map, "key");
print(val);

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()
}
