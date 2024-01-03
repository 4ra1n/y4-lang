package parser

import (
	"errors"
	"fmt"

	"github.com/4ra1n/y4-lang/ast"
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/token"
)

type Leaf struct {
	tokens []string
}

func NewLeaf(t []string) *Leaf {
	return &Leaf{tokens: t}
}

func (l *Leaf) Parse(lex *lexer.Lexer, res *base.List[ast.ASTree]) error {
	t, err := lex.Read()
	if err != nil {
		log.Error(err)
		return nil
	}
	if t.IsIdentifier() {
		for _, to := range l.tokens {
			text, err := t.GetText()
			if err != nil {
				log.Error(err)
				return err
			}
			if to == text {
				l.Find(res, t)
				return nil
			}
		}
	}
	if len(l.tokens) > 0 {
		return errors.New(fmt.Sprintf("%s expeted %s", l.tokens[0], t))
	} else {
		return errors.New(t.String())
	}
}

func (l *Leaf) Match(lex *lexer.Lexer) bool {
	t, err := lex.Peek(0)
	if err != nil {
		log.Error(err)
		return false
	}
	if t.IsIdentifier() {
		for _, to := range l.tokens {
			text, err := t.GetText()
			if err != nil {
				log.Error(err)
				return false
			}
			if to == text {
				return true
			}
		}
	}
	return false
}

func (l *Leaf) Find(res *base.List[ast.ASTree], t token.Token) {
	res.Add(ast.NewASTLeaf(t))
}
