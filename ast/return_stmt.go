package ast

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
)

type ReturnStmt struct {
	astList *ASTList
}

func NewReturnStmt(l *base.List[ASTree]) *ReturnStmt {
	return &ReturnStmt{astList: NewASTList(l)}
}

func (r *ReturnStmt) Child(i int) (ASTree, error) {
	return r.astList.Child(i)
}

func (r *ReturnStmt) NumChildren() int {
	return r.astList.NumChildren()
}

func (r *ReturnStmt) Children() *base.List[ASTree] {
	return r.astList.Children()
}

func (r *ReturnStmt) Location() string {
	return r.astList.Location()
}

func (r *ReturnStmt) GetString() (string, error) {
	ret, err := r.astList.GetString()
	return "<return-stmt:" + ret + ">", err
}

func (r *ReturnStmt) Eval(env envir.Environment) (interface{}, error) {
	a, err := r.Child(0)
	if err != nil {
		return nil, err
	}
	return a.Eval(env)
}

func (r *ReturnStmt) Lookup(sym *envir.Symbols) {
	r.astList.Lookup(sym)
}
