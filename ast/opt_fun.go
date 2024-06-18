package ast

import (
	"github.com/4ra1n/y4-lang/envir"
)

type OptFunction struct {
	size int
	fun  *Function
}

func NewOptFunction(parameters *ParameterList,
	body *BlockStmt, en envir.Environment, memorySize int) *OptFunction {
	f := NewFunction(parameters, body, en)
	of := &OptFunction{
		fun:  f,
		size: memorySize,
	}
	return of
}

func (o *OptFunction) Parameters() *ParameterList {
	return o.fun.Parameters()
}

func (o *OptFunction) Body() *BlockStmt {
	return o.fun.Body()
}

func (o *OptFunction) MakeEnv() envir.Environment {
	return envir.NewArrayEnv(o.size, o.fun.en.GetPool(), o.fun.en)
}

func EvalMain(bf *OptFunction, en envir.Environment) (interface{}, error) {
	env := envir.NewArrayEnv(bf.size, bf.fun.en.GetPool(), en)
	v, err := bf.Body().Eval(env)
	if err != nil {
		return nil, err
	}
	return v, nil
}
