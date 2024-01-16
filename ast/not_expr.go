package ast

import (
	"errors"

	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
)

type NotExpr struct {
	astList *ASTList
}

func NewNotExpr(l *base.List[ASTree]) *NotExpr {
	return &NotExpr{astList: NewASTList(l)}
}

func (n *NotExpr) Child(i int) (ASTree, error) {
	return n.astList.Child(i)
}

func (n *NotExpr) NumChildren() int {
	return n.astList.NumChildren()
}

func (n *NotExpr) Children() *base.List[ASTree] {
	return n.astList.Children()
}

func (n *NotExpr) Location() string {
	return n.astList.Location()
}

func (n *NotExpr) GetString() (string, error) {
	return "<not>", nil
}

func (n *NotExpr) Eval(env envir.Environment) (interface{}, error) {
	op, err := n.Child(0)
	if err != nil {
		return nil, err
	}
	v, err := op.Eval(env)
	if err != nil {
		return nil, err
	}
	vi, isInt := v.(int)
	if isInt {
		if vi == envir.FALSE {
			return envir.TRUE, nil
		} else if vi == envir.TRUE {
			return envir.FALSE, nil
		}
	}
	return nil, errors.New("not eval error")
}

func (n *NotExpr) Lookup(sym *envir.Symbols) {
	n.astList.Lookup(sym)
}
