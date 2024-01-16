package parser

import (
	"github.com/4ra1n/y4-lang/ast"
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/token"
)

type Skip struct {
	leaf *Leaf
}

func NewSkip(t []string) *Skip {
	return &Skip{leaf: NewLeaf(t)}
}

func (s *Skip) Parse(lex *lexer.Lexer, res *base.List[ast.ASTree]) error {
	return s.leaf.Parse(lex, res)
}

func (s *Skip) Match(lex *lexer.Lexer) bool {
	return s.leaf.Match(lex)
}

func (s *Skip) Find(res *base.List[ast.ASTree], t token.Token) {
}
