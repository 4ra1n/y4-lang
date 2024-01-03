package ast

import (
	"errors"

	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/color"
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/native"
)

type Arguments struct {
	postfix *ASTList
}

func NewArguments(l *base.List[ASTree]) *Arguments {
	return &Arguments{postfix: NewASTList(l)}
}

func (a *Arguments) Child(i int) (ASTree, error) {
	return a.postfix.Child(i)
}

func (a *Arguments) NumChildren() int {
	return a.postfix.NumChildren()
}

func (a *Arguments) Children() *base.List[ASTree] {
	return a.postfix.Children()
}

func (a *Arguments) Location() string {
	return a.postfix.Location()
}

func (a *Arguments) GetString() (string, error) {
	s, err := a.postfix.GetString()
	return "<arguments:" + s + ">", err
}

func (a *Arguments) Lookup(sym *envir.Symbols) {
	a.postfix.Lookup(sym)
}

func (a *Arguments) Eval(env envir.Environment) (interface{}, error) {
	return nil, nil
}

func (a *Arguments) Eval0(en envir.Environment, value interface{}) (interface{}, error) {
	if value == nil {
		return nil, errors.New("value is null")
	}
	nf, isNativeFunc := value.(*native.NativeFunction)
	if !isNativeFunc {
		bf, isFunc := value.(*OptFunction)
		if !isFunc {
			return nil, errors.New("bad function")
		}
		params := bf.Parameters()
		if a.Size() != params.Size() {
			return nil, errors.New("function params not match")
		}
		newEnv := bf.MakeEnv()
		num := 0
		for _, tree := range a.Children().Items() {
			res, err := tree.Eval(en)
			if err != nil {
				return nil, err
			}
			params.Eval0(newEnv, num, res)
			num++
		}
		return bf.Body().Eval(newEnv)
	} else {
		paramsLen := nf.GetNumParam()
		if a.Size() != paramsLen {
			color.RedPrintf("except: %d actual: %d\n", paramsLen, a.Size())
			return nil, errors.New("native params not match")
		}
		args := make([]interface{}, paramsLen)
		num := 0
		var err error
		for _, item := range a.Children().Items() {
			args[num], err = item.Eval(en)
			num++
			if err != nil {
				return nil, err
			}
		}
		return nf.Invoke(args)
	}
}

func (a *Arguments) Size() int {
	return a.postfix.NumChildren()
}
