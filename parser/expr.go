package parser

import (
	"github.com/4ra1n/y4-lang/ast"
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

type Expr struct {
	typ    string
	ops    *Operators
	factor *Parser
}

func NewExpr(typ string, exp *Parser, m *Operators) *Expr {
	return &Expr{
		typ:    typ,
		ops:    m,
		factor: exp,
	}
}

func (e *Expr) Parse(lex *lexer.Lexer, res *base.List[ast.ASTree]) error {
	right := e.factor.Parse(lex)
	var pre *Precedence
	for {
		pre = e.nextOperator(lex)
		if pre == nil {
			break
		}
		right = e.doShift(lex, right, pre.value)
	}
	res.Add(right)
	return nil
}

func (e *Expr) nextOperator(lex *lexer.Lexer) *Precedence {
	t, err := lex.Peek(0)
	if err != nil {
		log.Error(err)
		return nil
	}
	if t.IsIdentifier() {
		text, err := t.GetText()
		if err != nil {
			log.Error(err)
			return nil
		}
		v, ok := e.ops.Get(text)
		if ok {
			return v
		}
	}
	return nil
}

func (e *Expr) doShift(lex *lexer.Lexer, left ast.ASTree, pre int) ast.ASTree {
	list := base.NewList[ast.ASTree]()
	list.Add(left)
	l, err := lex.Read()
	if err != nil {
		log.Error(err)
		return nil
	}
	list.Add(ast.NewASTLeaf(l))
	right := e.factor.Parse(lex)
	var next *Precedence
	for {
		next = e.nextOperator(lex)
		if next == nil {
			break
		}
		if !rightIsExpr(pre, next) {
			break
		}
		right = e.doShift(lex, right, next.value)
	}
	list.Add(right)
	return makeAstList(e.typ, list)
}

func rightIsExpr(pre int, nextPre *Precedence) bool {
	if nextPre.leftAssoc {
		return pre < nextPre.value
	} else {
		return pre <= nextPre.value
	}
}

func (e *Expr) Match(lex *lexer.Lexer) bool {
	return e.factor.Match(lex)
}
