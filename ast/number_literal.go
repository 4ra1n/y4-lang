package ast

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/token"
)

type NumberLiteral struct {
	astLeaf *ASTLeaf
}

func NewNumberLiteral(t token.Token) *NumberLiteral {
	return &NumberLiteral{astLeaf: NewASTLeaf(t)}
}

func (n *NumberLiteral) Child(i int) (ASTree, error) {
	return n.astLeaf.Child(i)
}

func (n *NumberLiteral) NumChildren() int {
	return n.astLeaf.NumChildren()
}

func (n *NumberLiteral) Children() *base.List[ASTree] {
	return n.astLeaf.Children()
}

func (n *NumberLiteral) Location() string {
	return n.astLeaf.Location()
}

func (n *NumberLiteral) GetString() (string, error) {
	ns, err := n.astLeaf.GetString()
	return "<number-literal:" + ns + ">", err
}

func (n *NumberLiteral) Lookup(sym *envir.Symbols) {
	n.astLeaf.Lookup(sym)
}

func (n *NumberLiteral) Value() (interface{}, error) {
	return n.astLeaf.GetToken().GetNumber()
}

func (n *NumberLiteral) Eval(env envir.Environment) (interface{}, error) {
	return n.Value()
}

func (n *NumberLiteral) String() string {
	return n.astLeaf.Token.String()
}
