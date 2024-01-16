package parser

import (
	"github.com/4ra1n/y4-lang/ast"
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/lexer"
)

type Tree struct {
	parser *Parser
}

func NewTree(p *Parser) *Tree {
	return &Tree{parser: p}
}

func (t *Tree) Parse(lex *lexer.Lexer, res *base.List[ast.ASTree]) error {
	res.Add(t.parser.Parse(lex))
	return nil
}

func (t *Tree) Match(lex *lexer.Lexer) bool {
	return t.parser.Match(lex)
}
