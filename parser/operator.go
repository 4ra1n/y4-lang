package parser

import "github.com/4ra1n/y4-lang/base"

const (
	LEFT  = true
	RIGHT = false
)

type Operators struct {
	hashMap *base.Map[string, *Precedence]
}

func NewOperators() *Operators {
	return &Operators{
		hashMap: base.NewMap[string, *Precedence](),
	}
}

func (o *Operators) Add(name string, pre int, leftAssoc bool) {
	o.hashMap.Set(name, NewPrecedence(pre, leftAssoc))
}

func (o *Operators) Get(s string) (*Precedence, bool) {
	return o.hashMap.Get(s)
}
