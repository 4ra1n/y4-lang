package ast

import (
	"errors"

	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
)

type NegativeExpr struct {
	astList *ASTList
}

func NewNegativeExpr(l *base.List[ASTree]) *NegativeExpr {
	return &NegativeExpr{astList: NewASTList(l)}
}

func (n *NegativeExpr) Child(i int) (ASTree, error) {
	return n.astList.Child(i)
}

func (n *NegativeExpr) NumChildren() int {
	return n.astList.NumChildren()
}

func (n *NegativeExpr) Children() *base.List[ASTree] {
	return n.astList.Children()
}

func (n *NegativeExpr) Location() string {
	return n.astList.Location()
}

func (n *NegativeExpr) operator() (ASTree, error) {
	return n.astList.Child(0)
}

func (n *NegativeExpr) GetString() (string, error) {
	o, err := n.operator()
	if err != nil {
		return "", err
	}
	os, err := o.GetString()
	if err != nil {
		return "", err
	}
	return "-" + os, nil
}

func (n *NegativeExpr) Eval(env envir.Environment) (interface{}, error) {
	op, err := n.operator()
	if err != nil {
		return nil, err
	}
	v, err := op.Eval(env)
	if err != nil {
		return nil, err
	}
	vi, isInt := v.(int)
	if isInt {
		return -vi, nil
	} else {
		return nil, errors.New("negative eval error")
	}
}

func (n *NegativeExpr) Lookup(sym *envir.Symbols) {
	n.astList.Lookup(sym)
}
