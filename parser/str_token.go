package parser

import (
	"errors"

	"github.com/4ra1n/y4-lang/ast"
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/token"
)

type StrToken struct {
	typ string
}

func NewStrToken(typ string) *StrToken {
	idt := &StrToken{typ: typ}
	return idt
}

func (n *StrToken) Parse(lex *lexer.Lexer, res *base.List[ast.ASTree]) error {
	t, err := lex.Read()
	if err != nil {
		return err
	}
	if n.test(t) {
		leaf := makeLeaf(t, n.typ)
		res.Add(leaf)
	} else {
		return errors.New("str token test error")
	}
	return nil
}

func (n *StrToken) Match(lex *lexer.Lexer) bool {
	t, err := lex.Peek(0)
	if err != nil {
		return false
	}
	return n.test(t)
}

func (n *StrToken) test(t token.Token) bool {
	return t.IsString()
}
