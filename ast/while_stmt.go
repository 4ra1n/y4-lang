package ast

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/log"
)

type WhileStmt struct {
	astList *ASTList
}

func NewWhileStmt(l *base.List[ASTree]) *WhileStmt {
	return &WhileStmt{astList: NewASTList(l)}
}

func (w *WhileStmt) Condition() (ASTree, error) {
	return w.astList.Child(0)
}

func (w *WhileStmt) Body() (ASTree, error) {
	return w.astList.Child(1)
}

func (w *WhileStmt) Child(i int) (ASTree, error) {
	return w.astList.Child(i)
}

func (w *WhileStmt) NumChildren() int {
	return w.astList.NumChildren()
}

func (w *WhileStmt) Children() *base.List[ASTree] {
	return w.astList.Children()
}

func (w *WhileStmt) Location() string {
	return w.astList.Location()
}

func (w *WhileStmt) GetString() (string, error) {
	ws, err := w.astList.GetString()
	return "<while-stmt:" + ws + ">", err
}

func (w *WhileStmt) Eval(en envir.Environment) (interface{}, error) {
	var result interface{}
	for {
		c, err := w.Condition()
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
			ab, err := w.Body()
			if err != nil {
				log.Error(err)
				return nil, err
			}
			result, err = ab.Eval(en)
			if err != nil {
				log.Error(err)
				return nil, err
			}
			_, isBr := result.(*BreakSignal)
			if isBr {
				break
			}
			_, isCo := result.(*ContinueSignal)
			if isCo {
				continue
			}
		}
	}
	return nil, nil
}

func (w *WhileStmt) Lookup(sym *envir.Symbols) {
	w.astList.Lookup(sym)
}
