package ast

import (
	"errors"

	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/log"
)

type ParameterList struct {
	astList *ASTList
	offsets []int
}

func NewParameterList(l *base.List[ASTree]) *ParameterList {
	return &ParameterList{astList: NewASTList(l)}
}

func (p *ParameterList) Child(i int) (ASTree, error) {
	return p.astList.Child(i)
}

func (p *ParameterList) NumChildren() int {
	return p.astList.NumChildren()
}

func (p *ParameterList) Children() *base.List[ASTree] {
	return p.astList.Children()
}

func (p *ParameterList) Location() string {
	return p.astList.Location()
}

func (p *ParameterList) GetString() (string, error) {
	pa, err := p.astList.GetString()
	return "<parameter-list:" + pa + ">", err
}

func (p *ParameterList) Eval(env envir.Environment) (interface{}, error) {
	return nil, nil
}

func (p *ParameterList) Lookup(sym *envir.Symbols) {
	s := p.Size()
	p.offsets = make([]int, s)
	for i := 0; i < s; i++ {
		v, err := p.Name(i)
		if err != nil {
			log.Error(err)
			continue
		}
		p.offsets[i] = sym.PutNew(v)
	}
}

func (p *ParameterList) Name(i int) (string, error) {
	a, err := p.astList.Child(i)
	if err != nil {
		return "", err
	}
	al, ok := a.(*ASTLeaf)
	if !ok {
		return "", errors.New("parameter list cast error")
	}
	return al.GetToken().GetText()
}

func (p *ParameterList) Size() int {
	return p.astList.NumChildren()
}

func (p *ParameterList) Eval0(env envir.Environment, index int, val interface{}) {
	env.PutNest(0, p.offsets[index], val)
}
