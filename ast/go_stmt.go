package ast

import (
	"errors"

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
	// 这里需要克隆环境
	// 否则多线程传递的指针参数会被修改
	newEnv := env.Clone()
	fn := func() {
		_, err = fun.Eval(newEnv)
	}
	ok := env.NewJob(fn)
	if ok {
		return nil, nil
	} else {
		return nil, errors.New("创建新执行线程失败")
	}
}

func (g *GoStmt) Lookup(sym *envir.Symbols) {
	g.astList.Lookup(sym)
}
