package parser

import (
	"errors"

	"github.com/4ra1n/y4-lang/ast"
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/token"
)

type IdToken struct {
	typ      string
	reserved *base.HashSet[string]
}

func NewIdToken(typ string, r *base.HashSet[string]) *IdToken {
	return &IdToken{
		typ:      typ,
		reserved: r,
	}
}

func (i *IdToken) Parse(lex *lexer.Lexer, res *base.List[ast.ASTree]) error {
	t, err := lex.Read()
	if err != nil {
		return err
	}
	if i.test(t) {
		leaf := makeLeaf(t, i.typ)
		if leaf != nil {
			res.Add(leaf)
		}
	} else {
		return errors.New("id token test error")
	}
	return nil
}

func (i *IdToken) Match(lex *lexer.Lexer) bool {
	t, err := lex.Peek(0)
	if err != nil {
		return false
	}
	return i.test(t)
}

func (i *IdToken) test(t token.Token) bool {
	text, err := t.GetText()
	if err != nil {
		log.Error(err)
		return false
	}
	return t.IsIdentifier() && !i.reserved.Contains(text)
}
