package ast

import (
	"bytes"

	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/log"
)

type ASTList struct {
	children *base.List[ASTree]
}

func NewASTList(l *base.List[ASTree]) *ASTList {
	return &ASTList{children: l}
}

func (l *ASTList) Child(i int) (ASTree, error) {
	return l.children.Get(i)
}

func (l *ASTList) NumChildren() int {
	return l.children.Length()
}

func (l *ASTList) Children() *base.List[ASTree] {
	return l.children
}

func (l *ASTList) Location() string {
	for _, v := range l.children.Items() {
		s := v.Location()
		if s != "" {
			return s
		}
	}
	return ""
}

func (l *ASTList) GetString() (string, error) {
	buf := &bytes.Buffer{}
	buf.WriteString("[")
	sep := ""
	for _, v := range l.children.Items() {
		buf.WriteString(sep)
		sep = ", "
		s, err := v.GetString()
		if err != nil {
			return "", err
		}
		buf.WriteString(s)
	}
	buf.WriteString("]")
	return buf.String(), nil
}

func (l *ASTList) Eval(env envir.Environment) (interface{}, error) {
	for _, t := range l.Children().Items() {
		val, err := t.Eval(env)
		if err != nil {
			log.Error(err)
			return nil, err
		} else {
			return val, nil
		}
	}
	return nil, nil
}

func (l *ASTList) Lookup(sym *envir.Symbols) {
	for _, v := range l.children.Items() {
		v.Lookup(sym)
	}
}
