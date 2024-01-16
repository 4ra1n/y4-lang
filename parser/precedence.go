package parser

type Precedence struct {
	value     int
	leftAssoc bool
}

func NewPrecedence(v int, a bool) *Precedence {
	return &Precedence{
		value:     v,
		leftAssoc: a,
	}
}
