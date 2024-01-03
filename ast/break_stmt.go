package ast

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
)

type BreakStmt struct {
	astList *ASTList
}

func NewBreakStmt(l *base.List[ASTree]) *BreakStmt {
	return &BreakStmt{astList: NewASTList(l)}
}

func (b *BreakStmt) Child(i int) (ASTree, error) {
	return b.astList.Child(0)
}

func (b *BreakStmt) NumChildren() int {
	return b.astList.NumChildren()
}

func (b *BreakStmt) Children() *base.List[ASTree] {
	return b.astList.Children()
}

func (b *BreakStmt) Location() string {
	return b.astList.Location()
}

func (b *BreakStmt) GetString() (string, error) {
	return b.astList.GetString()
}

func (b *BreakStmt) Eval(env envir.Environment) (interface{}, error) {
	return &BreakSignal{true}, nil
}

func (b *BreakStmt) Lookup(sym *envir.Symbols) {
	b.astList.Lookup(sym)
}
