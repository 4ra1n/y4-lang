package lexer

import (
	"io"

	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/token"
)

const (
	EMPTY = -1
)

type Lexer struct {
	reader     io.Reader
	lastChar   rune
	lineNumber int
	queue      *base.List[token.Token]
}

func NewLexer(r io.Reader) *Lexer {
	log.Debug("new lexer")
	return &Lexer{
		reader:     r,
		lastChar:   EMPTY,
		lineNumber: 1,
		queue:      base.NewList[token.Token](),
	}
}
