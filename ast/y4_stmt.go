package ast

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
)

type GoStmt struct {
	astList *ASTList
}

func NewGoStmt(l *base.List[ASTree]) *GoStmt {
	return &GoStmt{astList: NewASTList(l)}
}

func (g *GoStmt) Child(i int) (ASTree, error) {
	return g.astList.Child(i)
}

func (g *GoStmt) NumChildren() int {
	return g.astList.NumChildren()
}

func (g *GoStmt) Children() *base.List[ASTree] {
	return g.astList.Children()
}

func (g *GoStmt) Location() string {
	return g.astList.Location()
}

func (g *GoStmt) GetString() (string, error) {
	gos, err := g.astList.GetString()
	return "<go-stmt:" + gos + ">", err
}

func (g *GoStmt) Eval(env envir.Environment) (interface{}, error) {
	fun, err := g.Child(1)
	if err != nil {
		return nil, err
	}
	fn := func() {
		_, err = fun.Eval(env)
	}
	env.NewJob(fn)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (g *GoStmt) Lookup(sym *envir.Symbols) {
	g.astList.Lookup(sym)
}
