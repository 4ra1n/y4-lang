package parser

import (
	"errors"

	"github.com/4ra1n/y4-lang/ast"
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/lexer"
)

type OrTree struct {
	parsers []*Parser
}

func NewOrTree(p []*Parser) *OrTree {
	return &OrTree{parsers: p}
}

func (o *OrTree) Parse(lex *lexer.Lexer, res *base.List[ast.ASTree]) error {
	p := o.Choose(lex)
	if p == nil {
		return errors.New("or tree parse error")
	} else {
		res.Add(p.Parse(lex))
	}
	return nil
}

func (o *OrTree) Match(lex *lexer.Lexer) bool {
	return o.Choose(lex) != nil
}

func (o *OrTree) Choose(l *lexer.Lexer) *Parser {
	for _, p := range o.parsers {
		if p.Match(l) {
			return p
		}
	}
	return nil
}

func (o *OrTree) Insert(p *Parser) {
	newParsers := make([]*Parser, len(o.parsers)+1)
	newParsers[0] = p
	copy(newParsers[1:], o.parsers)
	o.parsers = newParsers
}
