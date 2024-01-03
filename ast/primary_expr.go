package ast

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/log"
)

type PrimaryExpr struct {
	astList *ASTList
}

func NewPrimaryExpr(l *base.List[ASTree]) ASTree {
	if l.Length() == 1 {
		item, err := l.Get(0)
		if err != nil {
			log.Error(err)
			return nil
		}
		return item
	} else {
		return &PrimaryExpr{astList: NewASTList(l)}
	}
}

func (p *PrimaryExpr) Child(i int) (ASTree, error) {
	return p.astList.Child(i)
}

func (p *PrimaryExpr) NumChildren() int {
	return p.astList.NumChildren()
}

func (p *PrimaryExpr) Children() *base.List[ASTree] {
	return p.astList.Children()
}

func (p *PrimaryExpr) Location() string {
	return p.astList.Location()
}

func (p *PrimaryExpr) GetString() (string, error) {
	pr, err := p.astList.GetString()
	return "<primary-expr:" + pr + ">", err
}

func (p *PrimaryExpr) Lookup(sym *envir.Symbols) {
	p.astList.Lookup(sym)
}

func (p *PrimaryExpr) Operand() (ASTree, error) {
	return p.astList.Child(0)
}

func (p *PrimaryExpr) Postfix(nest int) Postfix {
	i := p.astList.NumChildren() - nest - 1
	a, err := p.astList.Child(i)

	ar, isAR := a.(*ArrayRef)
	if isAR {
		return ar
	}

	if err != nil {
		return nil
	}
	b, err := a.Child(1)
	if err != nil {
		return nil
	}
	po, ok := b.(Postfix)
	if !ok {
		return nil
	}
	return po
}

func (p *PrimaryExpr) HasPostfix(nest int) bool {
	return p.astList.NumChildren()-nest > 1
}

func (p *PrimaryExpr) Eval(env envir.Environment) (interface{}, error) {
	return p.evalSubExpr(env, 0)
}

func (p *PrimaryExpr) evalSubExpr(env envir.Environment, nest int) (interface{}, error) {
	if p.HasPostfix(nest) {
		target, err := p.evalSubExpr(env, nest+1)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return (p.Postfix(nest)).Eval0(env, target)
	} else {
		po, err := p.Operand()
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return po.Eval(env)
	}
}
