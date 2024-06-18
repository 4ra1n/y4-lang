package ast

import (
	"github.com/4ra1n/y4-lang/envir"
)

type Function struct {
	parameters *ParameterList
	body       *BlockStmt
	en         envir.Environment
}

func (f *Function) Body() *BlockStmt {
	return f.body
}

func (f *Function) Parameters() *ParameterList {
	return f.parameters
}

func NewFunction(parameters *ParameterList,
	body *BlockStmt, en envir.Environment) *Function {
	return &Function{parameters: parameters, body: body, en: en}
}
