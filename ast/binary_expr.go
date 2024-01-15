package ast

import (
	"errors"
	"fmt"
	"math"

	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/log"
)

type BinaryExpr struct {
	astList *ASTList
}

func NewBinaryExpr(l *base.List[ASTree]) *BinaryExpr {
	return &BinaryExpr{astList: NewASTList(l)}
}

func (b *BinaryExpr) Child(i int) (ASTree, error) {
	return b.astList.Child(i)
}

func (b *BinaryExpr) NumChildren() int {
	return b.astList.NumChildren()
}

func (b *BinaryExpr) Children() *base.List[ASTree] {
	return b.astList.Children()
}

func (b *BinaryExpr) Location() string {
	return b.astList.Location()
}

func (b *BinaryExpr) GetString() (string, error) {
	s, err := b.astList.GetString()
	return "<binary-expr:" + s + ">", err
}

func (b *BinaryExpr) left() (ASTree, error) {
	return b.astList.Child(0)
}

func (b *BinaryExpr) operator() (string, error) {
	a, err := b.astList.Child(1)
	if err != nil {
		return "", err
	}
	al, ok := a.(*ASTLeaf)
	if !ok {
		return "", errors.New("binary expr cast error")
	}
	return al.GetToken().GetText()
}

func (b *BinaryExpr) right() (ASTree, error) {
	return b.astList.Child(2)
}

func (b *BinaryExpr) Eval(env envir.Environment) (interface{}, error) {
	op, err := b.operator()
	if err != nil {
		return nil, err
	}
	if op == "=" {
		r, err := b.right()
		if err != nil {
			return nil, err
		}

		// check for array literal
		rl, isRl := r.(*ASTList)
		if isRl && rl.NumChildren() == 3 {
			ara, err := rl.Child(1)
			if err != nil {
				return nil, err
			}
			arr, isArr := ara.(*ArrayLiteral)
			if isArr {
				r = arr
			}
		}

		rv, err := r.Eval(env)
		if err != nil {
			return nil, err
		}
		return b.computeAssign(env, rv)
	} else {
		l, err := b.left()
		if err != nil {
			return nil, err
		}
		// c = (a+b) / 2
		if l.NumChildren() == 3 {
			lt, err := l.Child(1)
			if err != nil {
				return nil, err
			}
			lt, isBinary := lt.(*BinaryExpr)
			if isBinary {
				l = lt
			}
		}
		lv, err := l.Eval(env)
		if err != nil {
			return nil, err
		}
		r, err := b.right()
		if err != nil {
			return nil, err
		}
		// c = 50 / (a+b)
		if r.NumChildren() == 3 {
			rt, err := r.Child(1)
			if err != nil {
				return nil, err
			}
			rt, isBinary := rt.(*BinaryExpr)
			if isBinary {
				r = rt
			}
		}
		rv, err := r.Eval(env)
		if err != nil {
			return nil, err
		}
		return computeOp(lv, op, rv)
	}
}

func (b *BinaryExpr) Lookup(sym *envir.Symbols) {
	l, err := b.left()
	if err != nil {
		log.Error(err)
		return
	}
	op, err := b.operator()
	if err != nil {
		log.Error(err)
		return
	}
	if op == "=" {
		left, isName := l.(*Name)
		if isName {
			left.LookupForAssign(sym)
			r, err := b.right()
			if err != nil {
				log.Error(err)
				return
			}
			r.Lookup(sym)
		}
	}
	l.Lookup(sym)
	r, err := b.right()
	if err != nil {
		log.Error(err)
		return
	}
	r.Lookup(sym)
}

func (b *BinaryExpr) computeAssign(en envir.Environment, rvalue interface{}) (interface{}, error) {
	l, err := b.left()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	left, isName := l.(*Name)
	if isName {
		left.EvalForAssign(en, rvalue)
		return rvalue, nil
	} else {
		return b.computeAssign0(en, rvalue)
	}
}

func (b *BinaryExpr) computeAssign0(en envir.Environment, rvalue interface{}) (interface{}, error) {
	le, err := b.left()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	lep, isPri := le.(*PrimaryExpr)
	if isPri && lep.HasPostfix(0) {
		po := lep.Postfix(0)
		_, isArrayRef := po.(*ArrayRef)
		if isArrayRef {
			a, err := lep.evalSubExpr(en, 1)
			if err != nil {
				log.Error(err)
				return nil, err
			}
			ar, isIA := a.([]interface{})
			if isIA {
				ref, ok := lep.Postfix(0).(*ArrayRef)
				if !ok {
					log.Error("cast postfix to arrayref error")
					return nil, err
				}
				ri, err := ref.Index()
				if err != nil {
					log.Error(err)
					return nil, err
				}
				index, err := ri.Eval(en)
				if err != nil {
					log.Error(err)
					return nil, err
				}
				ii, isII := index.(int)
				if isII {
					ar[ii] = rvalue
					return rvalue, nil
				}
			}
			return nil, errors.New("bad array access")
		}
	}
	leN, isName := le.(*Name)
	if isName {
		leNN, err := leN.Name()
		if err != nil {
			log.Error(err)
			return nil, err
		}
		en.Put(leNN, rvalue)
		return rvalue, nil
	} else {
		return nil, errors.New("bad assignment")
	}
}

func computeOp(l interface{}, op string, r interface{}) (interface{}, error) {
	lei, isLI := l.(int)
	rgi, isRI := r.(int)
	leF, isLF := l.(float64)
	rgF, isRF := r.(float64)
	if isLI && isRI {
		return computeNumber(lei, op, rgi)
	} else if isLF && isRF {
		return computeFloat(leF, op, rgF)
	} else if isLI && isRF {
		leiNew := float64(lei)
		return computeFloat(leiNew, op, rgF)
	} else if isRI && isLF {
		rgiNew := float64(rgi)
		return computeFloat(leF, op, rgiNew)
	} else if op == "+" {
		ls := fmt.Sprintf("%v", l)
		rs := fmt.Sprintf("%v", r)
		return ls + rs, nil
	} else if op == "==" {
		if l == nil {
			if r == nil {
				return envir.TRUE, nil
			} else if r == 0 {
				// for null
				return envir.TRUE, nil
			} else {
				return envir.FALSE, nil
			}
		} else {
			if l == r {
				return envir.TRUE, nil
			} else {
				return envir.FALSE, nil
			}
		}
	} else if op == "!=" {
		if l == nil {
			if r == nil {
				return envir.FALSE, nil
			} else {
				return envir.TRUE, nil
			}
		} else {
			if l == r {
				return envir.FALSE, nil
			} else {
				return envir.TRUE, nil
			}
		}
	} else {
		return nil, errors.New("bad type")
	}
}

func computeNumber(lei int, op string, rgi int) (interface{}, error) {
	a := lei
	b := rgi
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	case "%":
		return a % b, nil
	case "==":
		if a == b {
			return envir.TRUE, nil
		} else {
			return envir.FALSE, nil
		}
	case "!=":
		if a != b {
			return envir.TRUE, nil
		} else {
			return envir.FALSE, nil
		}
	case ">":
		if a > b {
			return envir.TRUE, nil
		} else {
			return envir.FALSE, nil
		}
	case ">=":
		if a >= b {
			return envir.TRUE, nil
		} else {
			return envir.FALSE, nil
		}
	case "<":
		if a < b {
			return envir.TRUE, nil
		} else {
			return envir.FALSE, nil
		}
	case "<=":
		if a <= b {
			return envir.TRUE, nil
		} else {
			return envir.FALSE, nil
		}
	default:
		return nil, errors.New("bad operator")
	}
}

func computeFloat(lei float64, op string, rgi float64) (interface{}, error) {
	a := lei
	b := rgi
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	case "%":
		return math.Mod(lei, rgi), nil
	case "==":
		if a == b {
			return envir.TRUE, nil
		} else {
			return envir.FALSE, nil
		}
	case "!=":
		if a != b {
			return envir.TRUE, nil
		} else {
			return envir.FALSE, nil
		}
	case ">":
		if a > b {
			return envir.TRUE, nil
		} else {
			return envir.FALSE, nil
		}
	case ">=":
		if a >= b {
			return envir.TRUE, nil
		} else {
			return envir.FALSE, nil
		}
	case "<":
		if a < b {
			return envir.TRUE, nil
		} else {
			return envir.FALSE, nil
		}
	case "<=":
		if a <= b {
			return envir.TRUE, nil
		} else {
			return envir.FALSE, nil
		}
	default:
		return nil, errors.New("bad operator")
	}
}
