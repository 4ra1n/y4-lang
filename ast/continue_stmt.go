package ast

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
)

type ContinueStmt struct {
	astList *ASTList
}

func NewContinueStmt(l *base.List[ASTree]) *ContinueStmt {
	return &ContinueStmt{astList: NewASTList(l)}
}

func (c *ContinueStmt) Child(i int) (ASTree, error) {
	return c.astList.Child(i)
}

func (c *ContinueStmt) NumChildren() int {
	return c.astList.NumChildren()
}

func (c *ContinueStmt) Children() *base.List[ASTree] {
	return c.astList.Children()
}

func (c *ContinueStmt) Location() string {
	return c.astList.Location()
}

func (c *ContinueStmt) GetString() (string, error) {
	return c.astList.GetString()
}

func (c *ContinueStmt) Eval(env envir.Environment) (interface{}, error) {
	return &ContinueSignal{Val: true}, nil
}

func (c *ContinueStmt) Lookup(sym *envir.Symbols) {
	c.astList.Lookup(sym)
}
