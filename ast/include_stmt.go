package ast

import (
	"errors"

	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/lib"
	"github.com/4ra1n/y4-lang/log"
)

type IncludeStmt struct {
	astList *ASTList
}

func NewIncludeStmt(l *base.List[ASTree]) *IncludeStmt {
	return &IncludeStmt{astList: NewASTList(l)}
}

func (ic *IncludeStmt) Child(i int) (ASTree, error) {
	return ic.astList.Child(i)
}

func (ic *IncludeStmt) NumChildren() int {
	return ic.astList.NumChildren()
}

func (ic *IncludeStmt) Children() *base.List[ASTree] {
	return ic.astList.Children()
}

func (ic *IncludeStmt) Location() string {
	return ic.astList.Location()
}

func (ic *IncludeStmt) GetString() (string, error) {
	ics, err := ic.astList.GetString()
	return "<include-stmt:" + ics + ">", err
}

func (ic *IncludeStmt) Eval(en envir.Environment) (interface{}, error) {
	i, err := ic.Child(0)
	if err != nil {
		return nil, err
	}
	iv, err := i.Eval(en)
	ivs, isS := iv.(string)
	if !isS {
		log.Error("not a string")
		return nil, errors.New("lib name not a string")
	}
	lib.AddLib(ivs, en)
	return nil, nil
}

func (ic *IncludeStmt) Lookup(sym *envir.Symbols) {
	ic.astList.Lookup(sym)
}
