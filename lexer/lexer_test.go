package lexer

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/4ra1n/y4-lang/log"
)

func TestInputLexer(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := `#include "base64"`
	reader := bytes.NewReader([]byte(input))
	lexer := NewLexer(reader)
	tk, err := lexer.read0()
	if err != nil {
		panic(err)
	}
	fmt.Println(tk)
}
