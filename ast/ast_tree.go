package ast

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
)

type ASTree interface {
	Child(i int) (ASTree, error)
	NumChildren() int
	Children() *base.List[ASTree]
	Location() string
	GetString() (string, error)
	Eval(env envir.Environment) (interface{}, error)
	Lookup(sym *envir.Symbols)
}
