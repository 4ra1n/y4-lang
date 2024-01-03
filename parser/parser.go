package parser

import (
	"github.com/4ra1n/y4-lang/ast"
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

type Parser struct {
	elements *base.List[Element]
	typ      string
}

func NewParser(typ string) *Parser {
	p := &Parser{typ: typ}
	p.ResetWithType(typ)
	return p
}

func NewParserWithParser(p *Parser) *Parser {
	return &Parser{
		elements: p.elements,
		typ:      p.typ,
	}
}

func (p *Parser) Parse(l *lexer.Lexer) ast.ASTree {
	results := base.NewList[ast.ASTree]()
	items := p.elements.Items()
	for _, e := range items {
		err := e.Parse(l, results)
		if err != nil {
			log.Error(err)
		}
	}
	return makeAstList(p.typ, results)
}

func (p *Parser) Match(l *lexer.Lexer) bool {
	if p.elements.Length() == 0 {
		return true
	} else {
		e, err := p.elements.Get(0)
		if err != nil {
			return false
		}
		return e.Match(l)
	}
}

func RuleWithType(typ string) *Parser {
	return NewParser(typ)
}

func RuleNoType() *Parser {
	return NewParser("")
}

func (p *Parser) ResetNoType() *Parser {
	p.elements = base.NewList[Element]()
	return p
}

func (p *Parser) ResetWithType(typ string) *Parser {
	p.elements = base.NewList[Element]()
	p.typ = typ
	return p
}

func (p *Parser) NumberNoType() *Parser {
	return p.NumberWithType("")
}

func (p *Parser) NumberWithType(typ string) *Parser {
	p.elements.Add(NewNumToken(typ))
	return p
}

func (p *Parser) IdentifierNoType(reserved *base.HashSet[string]) *Parser {
	return p.IdentifierWithType("", reserved)
}

func (p *Parser) IdentifierWithType(typ string, reserved *base.HashSet[string]) *Parser {
	p.elements.Add(NewIdToken(typ, reserved))
	return p
}

func (p *Parser) StringNoType() *Parser {
	return p.StringWithType("")
}

func (p *Parser) StringWithType(typ string) *Parser {
	p.elements.Add(NewStrToken(typ))
	return p
}

func (p *Parser) Token(pat ...string) *Parser {
	p.elements.Add(NewLeaf(pat))
	return p
}

func (p *Parser) Sep(pat ...string) *Parser {
	p.elements.Add(NewSkip(pat))
	return p
}

func (p *Parser) Ast(pa *Parser) *Parser {
	p.elements.Add(NewTree(pa))
	return p
}

func (p *Parser) Or(pa ...*Parser) *Parser {
	p.elements.Add(NewOrTree(pa))
	return p
}

func (p *Parser) Maybe(pa *Parser) *Parser {
	p2 := NewParserWithParser(pa)
	p2.ResetNoType()
	p.elements.Add(NewOrTree([]*Parser{pa, p2}))
	return p
}

func (p *Parser) Option(pa *Parser) *Parser {
	p.elements.Add(NewRepeat(pa, true))
	return p
}

func (p *Parser) Repeat(pa *Parser) *Parser {
	p.elements.Add(NewRepeat(pa, false))
	return p
}

func (p *Parser) ExpressionNoType(sub *Parser, operators *Operators) *Parser {
	p.elements.Add(NewExpr("", sub, operators))
	return p
}

func (p *Parser) ExpressionWithType(typ string, sub *Parser, operators *Operators) *Parser {
	p.elements.Add(NewExpr(typ, sub, operators))
	return p
}

func (p *Parser) InsertChoice(pa *Parser) *Parser {
	e, err := p.elements.Get(0)
	if err != nil {
		log.Error(e)
		return nil
	}
	ot, ok := e.(*OrTree)
	if ok {
		ot.Insert(pa)
	} else {
		otherwise := NewParserWithParser(p)
		p.ResetWithType("")
		p.Or(pa, otherwise)
	}
	return p
}
