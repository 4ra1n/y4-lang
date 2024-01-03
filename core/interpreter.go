package core

import (
	"github.com/4ra1n/y4-lang/ast"
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/native"
	"github.com/4ra1n/y4-lang/token"
)

type Interpreter struct {
	lexer  *lexer.Lexer
	parser *CoreParser
}

func NewInterpreter(l *lexer.Lexer) *Interpreter {
	return &Interpreter{
		lexer:  l,
		parser: NewCoreParser(),
	}
}

func (i *Interpreter) Start() {
	en := native.NewNative(envir.NewResizableEnv()).Environment()
	for {
		v, err := i.lexer.Peek(0)
		if err != nil {
			log.Debug(err)
			break
		}
		if v == token.EOF {
			log.Debug("exit reason: eof")
			break
		}
		ct, ok := v.(*token.CommentToken)
		if ok {
			log.Debugf("line %d ignore comment", ct.GetLineNumber())
			_, _ = i.lexer.Read()
			continue
		}
		// core parse
		t := i.parser.Parse(i.lexer)
		if t == nil {
			log.Debug("ast is null")
			continue
		}
		// debug info
		s, err := t.GetString()
		if err != nil {
			log.Error(err)
			continue
		}
		// check ast list
		astList, ok := t.(ast.ASTree)
		if !ok {
			log.Error("parse ast list error")
			break
		}
		// if null stmt skip
		tree, err := astList.Children().Get(0)
		if err != nil {
			log.Error("check null stmt error")
			break
		}
		_, ok = tree.(*ast.NullStmt)
		if ok {
			log.Debug("ignore null stmt")
			continue
		}
		log.Infof("eval: %s", s)
		// eval ast
		t.Lookup(en.Symbols())
		_, err = t.Eval(en)
		if err != nil {
			log.Error(err)
		}
	}
	ok := en.WaitJob()
	if ok {
		log.Info("all threads finish")
	}
}