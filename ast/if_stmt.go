package ast

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
)

type IfStmt struct {
	astLst *ASTList
}

func NewIfStmt(l *base.List[ASTree]) *IfStmt {
	return &IfStmt{astLst: NewASTList(l)}
}

func (is *IfStmt) Child(i int) (ASTree, error) {
	return is.astLst.Child(i)
}

func (is *IfStmt) NumChildren() int {
	return is.astLst.NumChildren()
}

func (is *IfStmt) Children() *base.List[ASTree] {
	return is.astLst.Children()
}

func (is *IfStmt) Location() string {
	return is.astLst.Location()
}

func (is *IfStmt) GetString() (string, error) {
	ifs, err := is.astLst.GetString()
	return "<if-stmt:" + ifs + ">", err
}

func (is *IfStmt) Eval(en envir.Environment) (interface{}, error) {
	c, err := is.Condition()
	if err != nil {
		return nil, err
	}
	cv, err := c.Eval(en)
	if err != nil {
		return nil, err
	}
	ci, isI := cv.(int)
	if isI && ci != envir.FALSE {
		tb, err := is.ThenBlock()
		if err != nil {
			return nil, err
		}
		tbv, err := tb.Eval(en)
		if err != nil {
			return nil, err
		}
		return tbv, nil
	} else {
		b, err := is.ElseBlock()
		if err != nil {
			return nil, err
		}
		if b == nil {
			return 0, nil
		} else {
			return b.Eval(en)
		}
	}
}

func (is *IfStmt) Lookup(sym *envir.Symbols) {
	is.astLst.Lookup(sym)
}

func (is *IfStmt) Condition() (ASTree, error) {
	return is.Child(0)
}

func (is *IfStmt) ThenBlock() (ASTree, error) {
	return is.Child(1)
}

func (is *IfStmt) ElseBlock() (ASTree, error) {
	if is.NumChildren() > 2 {
		return is.Child(2)
	} else {
		return nil, nil
	}
}
