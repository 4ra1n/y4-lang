package ast

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/log"
)

type ArrayLiteral struct {
	astList *ASTList
}

func NewArrayLiteral(l *base.List[ASTree]) *ArrayLiteral {
	return &ArrayLiteral{astList: NewASTList(l)}
}

func (a *ArrayLiteral) Child(i int) (ASTree, error) {
	return a.astList.Child(i)
}

func (a *ArrayLiteral) NumChildren() int {
	return a.astList.NumChildren()
}

func (a *ArrayLiteral) Children() *base.List[ASTree] {
	return a.astList.Children()
}

func (a *ArrayLiteral) Location() string {
	return a.astList.Location()
}

func (a *ArrayLiteral) GetString() (string, error) {
	s, err := a.astList.GetString()
	return "<array-literal:" + s + ">", err
}

func (a *ArrayLiteral) Lookup(sym *envir.Symbols) {
	a.astList.Lookup(sym)
}

func (a *ArrayLiteral) Eval(env envir.Environment) (interface{}, error) {
	s := a.NumChildren()
	res := make([]interface{}, s)
	i := 0
	for _, t := range a.astList.Children().Items() {
		o, err := t.Eval(env)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		res[i] = o
		i++
	}
	return res, nil
}

func (a *ArrayLiteral) Size() int {
	return a.astList.NumChildren()
}
