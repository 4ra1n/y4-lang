package parser

import (
	"github.com/4ra1n/y4-lang/ast"
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/lexer"
)

type Element interface {
	Parse(lex *lexer.Lexer, res *base.List[ast.ASTree]) error
	Match(lex *lexer.Lexer) bool
}
