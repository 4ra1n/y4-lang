package ast

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/token"
)

type StringLiteral struct {
	astLeaf *ASTLeaf
}

func NewStringLiteral(t token.Token) *StringLiteral {
	return &StringLiteral{astLeaf: NewASTLeaf(t)}
}

func (s *StringLiteral) Child(i int) (ASTree, error) {
	return s.astLeaf.Child(i)
}

func (s *StringLiteral) NumChildren() int {
	return s.astLeaf.NumChildren()
}

func (s *StringLiteral) Children() *base.List[ASTree] {
	return s.astLeaf.Children()
}

func (s *StringLiteral) Location() string {
	return s.astLeaf.Location()
}

func (s *StringLiteral) GetString() (string, error) {
	ss, err := s.astLeaf.GetString()
	return "<string-literal:" + ss + ">", err
}

func (s *StringLiteral) Lookup(sym *envir.Symbols) {
	s.astLeaf.Lookup(sym)
}

func (s *StringLiteral) Value() (string, error) {
	return s.astLeaf.GetToken().GetText()
}

func (s *StringLiteral) Eval(env envir.Environment) (interface{}, error) {
	return s.Value()
}
