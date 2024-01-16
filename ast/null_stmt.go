package ast

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
)

type NullStmt struct {
	astList *ASTList
}

func NewNullStmt(l *base.List[ASTree]) *NullStmt {
	return &NullStmt{astList: NewASTList(l)}
}

func (n *NullStmt) Child(i int) (ASTree, error) {
	return n.astList.Child(i)
}

func (n *NullStmt) NumChildren() int {
	return n.astList.NumChildren()
}

func (n *NullStmt) Children() *base.List[ASTree] {
	return n.astList.Children()
}

func (n *NullStmt) Location() string {
	return n.astList.Location()
}

func (n *NullStmt) GetString() (string, error) {
	s, err := n.astList.GetString()
	return "<null-stmt:" + s + ">", err
}

func (n *NullStmt) Eval(env envir.Environment) (interface{}, error) {
	return n.astList.Eval(env)
}

func (n *NullStmt) Lookup(sym *envir.Symbols) {
	n.astList.Lookup(sym)
}
