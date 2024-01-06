package ast

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
)

type ArrayRef struct {
	postfix *ASTList
}

func NewArrayRef(l *base.List[ASTree]) *ArrayRef {
	return &ArrayRef{postfix: NewASTList(l)}
}

func (a *ArrayRef) Child(i int) (ASTree, error) {
	return a.postfix.Child(i)
}

func (a *ArrayRef) NumChildren() int {
	return a.postfix.NumChildren()
}

func (a *ArrayRef) Children() *base.List[ASTree] {
	return a.postfix.Children()
}

func (a *ArrayRef) Location() string {
	return a.postfix.Location()
}

func (a *ArrayRef) Lookup(sym *envir.Symbols) {
	a.postfix.Lookup(sym)
}

func (a *ArrayRef) Eval(env envir.Environment) (interface{}, error) {
	return nil, nil
}

func (a *ArrayRef) Eval0(env envir.Environment, value interface{}) (interface{}, error) {
	valReflectValue := reflect.ValueOf(value)
	if valReflectValue.Kind() == reflect.Slice {
		pi, err := a.Index()
		if err != nil {
			return nil, err
		}
		r, err := pi.Eval(env)
		if err != nil {
			return nil, err
		}
		idx, ok := r.(int)
		if !ok {
			return nil, errors.New("index is not an integer")
		}
		if idx < 0 || idx >= valReflectValue.Len() {
			return nil, errors.New("index out of bounds")
		}
		return valReflectValue.Index(idx).Interface(), nil
	}
	return nil, fmt.Errorf("value is not a slice")
}

func (a *ArrayRef) Index() (ASTree, error) {
	return a.postfix.Child(0)
}

func (a *ArrayRef) GetString() (string, error) {
	ast, err := a.Index()
	if err != nil {
		return "", err
	}
	s, err := ast.GetString()
	if err != nil {
		return "", err
	}
	return "<array-ref:[" + s + "]>", nil
}
