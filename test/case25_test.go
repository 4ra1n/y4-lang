package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
	"github.com/stretchr/testify/assert"
)

func TestCase25(t *testing.T) {
	Finish()
	Redirect()
	code := `
#引入 "贝斯六四"
#引入 "海克斯"

甲 = 贝斯六四.编码("4ra1n");
打印(甲);

乙 = 海克斯.编码("4ra1n");
打印(乙);

丙 = 贝斯六四.解码(甲);
打印(丙);

`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "NHJhMW4=\n")
	assert.Contains(t, result, "4ra1n\n")
	assert.Contains(t, result, "347261316e\n")
}
