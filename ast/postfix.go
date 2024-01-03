package ast

import (
	"github.com/4ra1n/y4-lang/envir"
)

type Postfix interface {
	Eval0(env envir.Environment, value interface{}) (interface{}, error)
}
