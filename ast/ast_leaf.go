package ast

import (
	"errors"
	"strconv"

	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/token"
)

var empty = base.NewList[ASTree]()

type ASTLeaf struct {
	Token token.Token
}

func NewASTLeaf(token token.Token) *ASTLeaf {
	return &ASTLeaf{Token: token}
}

func (leaf *ASTLeaf) Child(i int) (ASTree, error) {
	return nil, errors.New("index out of bounds")
}

func (leaf *ASTLeaf) NumChildren() int {
	return 0
}

func (leaf *ASTLeaf) Children() *base.List[ASTree] {
	return empty
}

func (leaf *ASTLeaf) GetString() (string, error) {
	s, err := leaf.Token.GetText()
	return "<ast-leaf:" + s + ">", err
}

func (leaf *ASTLeaf) Location() string {
	return "at line " + strconv.Itoa(leaf.Token.GetLineNumber())
}

func (leaf *ASTLeaf) Eval(env envir.Environment) (interface{}, error) {
	return nil, nil
}

func (leaf *ASTLeaf) GetToken() token.Token {
	return leaf.Token
}

func (leaf *ASTLeaf) Lookup(sym *envir.Symbols) {
}
