package ast

import (
	"github.com/4ra1n/y4-lang/envir"
)

type Function struct {
	parameters *ParameterList
	body       *BlockStmt
	return0    interface{}
	en         envir.Environment
}

func (f *Function) Body() *BlockStmt {
	return f.body
}

func (f *Function) Parameters() *ParameterList {
	return f.parameters
}

func (f *Function) Return0() interface{} {
	return f.return0
}

func (f *Function) SetReturn0(r interface{}) {
	f.return0 = r
}

func (f *Function) MakeEnv() envir.Environment {
	return envir.NewBasicEnvWithEnv(f.en)
}

func NewFunction(parameters *ParameterList,
	body *BlockStmt, en envir.Environment) *Function {
	return &Function{parameters: parameters, body: body, en: en}
}
