package ast

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/log"
)

type ForStmt struct {
	astList *ASTList
}

func NewForStmt(l *base.List[ASTree]) *ForStmt {
	return &ForStmt{astList: NewASTList(l)}
}

func (w *ForStmt) Child(i int) (ASTree, error) {
	return w.astList.Child(i)
}

func (w *ForStmt) NumChildren() int {
	return w.astList.NumChildren()
}

func (w *ForStmt) Children() *base.List[ASTree] {
	return w.astList.Children()
}

func (w *ForStmt) Location() string {
	return w.astList.Location()
}

func (w *ForStmt) GetString() (string, error) {
	ws, err := w.astList.GetString()
	return "<for-stmt:" + ws + ">", err
}

func (w *ForStmt) Eval(en envir.Environment) (interface{}, error) {
	var result interface{}
	initExpr, err := w.Child(1)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	_, err = initExpr.Eval(en)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	for {
		c, err := w.Child(3)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		r, err := c.Eval(en)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		r, isInt := r.(int)
		if isInt && r == envir.FALSE {
			return result, nil
		} else {
			ab, err := w.Child(6)
			if err != nil {
				log.Error(err)
				return nil, err
			}
			result, err = ab.Eval(en)
			if err != nil {
				log.Error(err)
				return nil, err
			}
			rs, isRR := result.(*ReturnSignal)
			if isRR {
				return rs, nil
			}
			_, isBr := result.(*BreakSignal)
			if isBr {
				break
			}
			_, isCo := result.(*ContinueSignal)
			if isCo {
				continue
			}

			// exec iterator
			iter, err := w.Child(5)
			if err != nil {
				log.Error(err)
				return nil, err
			}
			_, err = iter.Eval(en)
			if err != nil {
				log.Error(err)
				return nil, err
			}
		}
	}
	return nil, nil
}

func (w *ForStmt) Lookup(sym *envir.Symbols) {
	w.astList.Lookup(sym)
}
