package ast

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
)

type BlockStmt struct {
	astList *ASTList
}

func NewBlockStmt(l *base.List[ASTree]) *BlockStmt {
	return &BlockStmt{astList: NewASTList(l)}
}

func (b *BlockStmt) Child(i int) (ASTree, error) {
	return b.astList.Child(i)
}

func (b *BlockStmt) NumChildren() int {
	return b.astList.NumChildren()
}

func (b *BlockStmt) Children() *base.List[ASTree] {
	return b.astList.Children()
}

func (b *BlockStmt) Location() string {
	return b.astList.Location()
}

func (b *BlockStmt) GetString() (string, error) {
	s, err := b.astList.GetString()
	return "<block-stmt:" + s + ">", err
}

func (b *BlockStmt) Lookup(sym *envir.Symbols) {
	b.astList.Lookup(sym)
}

func (b *BlockStmt) Eval(env envir.Environment) (interface{}, error) {
	var result interface{}
	for _, a := range b.Children().Items() {
		_, isRt := a.(*ReturnStmt)
		if isRt {
			v, err := a.Eval(env)
			if err != nil {
				return nil, err
			}
			return &ReturnSignal{
				Val: v,
			}, nil
		}
		_, isNull := a.(*NullStmt)
		if !isNull {
			var err error
			result, err = a.Eval(env)
			if err != nil {
				return nil, err
			}
			rs, isRs := result.(*ReturnSignal)
			if isRs {
				return rs, nil
			}
			br, isBr := result.(*BreakSignal)
			if isBr {
				return br, nil
			}
			co, isCo := result.(*ContinueSignal)
			if isCo {
				return co, nil
			}
		}
	}
	return result, nil
}
