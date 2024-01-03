package ast

import (
	"bytes"
	"errors"

	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/log"
)

type DefStmt struct {
	astList *ASTList
	size    int
	index   int
}

func NewDefStmt(l *base.List[ASTree]) *DefStmt {
	return &DefStmt{astList: NewASTList(l)}
}

func (d *DefStmt) Child(i int) (ASTree, error) {
	return d.astList.Child(i)
}

func (d *DefStmt) NumChildren() int {
	return d.astList.NumChildren()
}

func (d *DefStmt) Children() *base.List[ASTree] {
	return d.astList.Children()
}

func (d *DefStmt) Location() string {
	return d.astList.Location()
}

func (d *DefStmt) Name() (string, error) {
	a, err := d.astList.Child(1)
	if err != nil {
		return "", err
	}
	al, ok := a.(*ASTLeaf)
	if !ok {
		return "", errors.New("def stmt cast error")
	}
	return al.GetToken().GetText()
}

func (d *DefStmt) Parameters() (*ParameterList, error) {
	a, err := d.astList.Child(2)
	if err != nil {
		return nil, err
	}
	p, err := a.Child(1)
	if err != nil {
		return nil, err
	}
	pa, ok := p.(*ParameterList)
	if !ok {
		return nil, errors.New("def stmt cast error")
	}
	return pa, nil
}

func (d *DefStmt) Body() (*BlockStmt, error) {
	a, err := d.astList.Child(3)
	if err != nil {
		return nil, err
	}
	p, ok := a.(*BlockStmt)
	if !ok {
		return nil, errors.New("def stmt cast error")
	}
	return p, nil
}

func FunctionLookup(sym *envir.Symbols, params *ParameterList, body *BlockStmt) int {
	newSym := envir.NewSymbols(sym)
	params.Lookup(newSym)
	body.Lookup(newSym)
	return newSym.Size()
}

func (d *DefStmt) Lookup(sym *envir.Symbols) {
	dn, err := d.Name()
	if err != nil {
		log.Error(err)
		return
	}
	pl, err := d.Parameters()
	if err != nil {
		log.Error(err)
		return
	}
	pb, err := d.Body()
	if err != nil {
		log.Error(err)
		return
	}
	d.index = sym.PutNew(dn)
	d.size = FunctionLookup(sym, pl, pb)
}

func (d *DefStmt) Eval(en envir.Environment) (interface{}, error) {
	pa, err := d.Parameters()
	if err != nil {
		return nil, err
	}
	bo, err := d.Body()
	if err != nil {
		return nil, err
	}
	en.PutNest(0, d.index,
		NewOptFunction(pa, bo, en, d.size))
	nn, err := d.Name()
	if err != nil {
		return nil, err
	}
	return nn, nil
}

func (d *DefStmt) GetString() (string, error) {
	buf := &bytes.Buffer{}
	buf.WriteString("<def-stmt:")
	n, err := d.Name()
	if err != nil {
		return "", err
	}
	buf.WriteString(n)
	buf.WriteString(" ")
	p, err := d.Parameters()
	if err != nil {
		return "", err
	}
	pn, err := p.GetString()
	if err != nil {
		return "", err
	}
	buf.WriteString(pn)
	buf.WriteString(" ")
	pb, err := d.Body()
	if err != nil {
		return "", err
	}
	pbs, err := pb.GetString()
	if err != nil {
		return "", err
	}
	buf.WriteString(pbs)
	buf.WriteString(">")
	return buf.String(), nil
}
