package lexer

import (
	"io"

	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/token"
)

const (
	EMPTY = 0
)

type Lexer struct {
	reader     io.Reader
	lastChar   byte
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
