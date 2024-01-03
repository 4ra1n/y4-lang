package ast

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/token"
)

const (
	UNKNOWN = -1
)

type Name struct {
	astLeaf *ASTLeaf
	nest    int
	index   int
}

func NewName(t token.Token) *Name {
	return &Name{
		astLeaf: NewASTLeaf(t),
		nest:    0,
		index:   UNKNOWN,
	}
}

func (n *Name) Child(i int) (ASTree, error) {
	return n.astLeaf.Child(i)
}

func (n *Name) NumChildren() int {
	return n.astLeaf.NumChildren()
}

func (n *Name) Children() *base.List[ASTree] {
	return n.astLeaf.Children()
}

func (n *Name) Location() string {
	return n.astLeaf.Location()
}

func (n *Name) GetString() (string, error) {
	na, err := n.astLeaf.GetString()
	return "<name:" + na + ">", err
}

func (n *Name) Token() token.Token {
	return n.astLeaf.Token
}

func (n *Name) Name() (string, error) {
	return n.Token().GetText()
}

func (n *Name) Eval(env envir.Environment) (interface{}, error) {
	nn, err := n.Name()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if n.index == UNKNOWN {
		return env.Get(nn), nil
	} else {
		return env.GetNest(n.nest, n.index), nil
	}
}

func (n *Name) Lookup(sym *envir.Symbols) {
	nn, err := n.Name()
	if err != nil {
		log.Error(err)
		return
	}
	loc := sym.Get(nn)
	if loc == nil {
		log.Errorf("undefined name: %s", nn)
	} else {
		n.nest = loc.Nest
		n.index = loc.Index
	}
}

func (n *Name) LookupForAssign(sym *envir.Symbols) {
	nn, err := n.Name()
	if err != nil {
		log.Error(err)
		return
	}
	loc := sym.Put(nn)
	n.nest = loc.Nest
	n.index = loc.Index
}

func (n *Name) EvalForAssign(env envir.Environment, value interface{}) {
	nn, err := n.Name()
	if err != nil {
		log.Error(err)
		return
	}
	if n.index == UNKNOWN {
		env.Put(nn, value)
	} else {
		env.PutNest(n.nest, n.index, value)
	}
}
