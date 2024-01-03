package parser

import (
	"github.com/4ra1n/y4-lang/ast"
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/lexer"
)

type Repeat struct {
	parser   *Parser
	onlyOnce bool
}

func NewRepeat(p *Parser, once bool) *Repeat {
	return &Repeat{
		parser:   p,
		onlyOnce: once,
	}
}

func (r *Repeat) Parse(lex *lexer.Lexer, res *base.List[ast.ASTree]) error {
	for {
		ok := r.parser.Match(lex)
		if !ok {
			break
		}
		t := r.parser.Parse(lex)
		if t == nil {
			continue
		}
		_, ok = t.(*ast.ASTList)
		if ok || t.NumChildren() > 0 {
			res.Add(t)
		}
		if r.onlyOnce {
			break
		}
	}
	return nil
}

func (r *Repeat) Match(lex *lexer.Lexer) bool {
	return r.parser.Match(lex)
}
