package core

import (
	"context"

	"github.com/4ra1n/y4-lang/ast"
	"github.com/4ra1n/y4-lang/envir"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/native"
	"github.com/4ra1n/y4-lang/token"
)

type Interpreter struct {
	lexer    *lexer.Lexer
	parser   *Parser
	cancel   context.CancelFunc
	envSize  int
	poolSize int
}

func NewInterpreter(l *lexer.Lexer, cancel context.CancelFunc) *Interpreter {
	return &Interpreter{
		lexer:    l,
		parser:   NewCoreParser(),
		cancel:   cancel,
		envSize:  envir.DefaultEnvSize,
		poolSize: envir.DefaultPoolSize,
	}
}

func (i *Interpreter) SetEnvSize(size int) {
	i.envSize = size
}

func (i *Interpreter) SetPoolSize(size int) {
	i.poolSize = size
}

func (i *Interpreter) Start() {
	env := envir.NewResizableEnv(i.envSize, i.poolSize)
	en := native.NewNative(env).Environment()
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

	// check main method
	main := en.Get("主函数")
	if main != nil {
		mainMethod, isOpt := main.(*ast.OptFunction)
		if isOpt {
			_, err := ast.EvalMain(mainMethod, en)
			if err != nil {
				log.Error(err)
			}
		}
	}

	ok := en.WaitJob()
	if ok {
		log.Info("all threads finish")
	}
	if i.cancel != nil {
		i.cancel()
	} else {
		log.Info("test finish")
	}
}
