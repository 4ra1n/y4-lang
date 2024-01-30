package core

import (
	"context"

	"github.com/4ra1n/y4-lang/ast"
	"github.com/4ra1n/y4-lang/conf"
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
	// 根据设置先构建环境
	env := envir.NewResizableEnv(i.envSize, i.poolSize)
	// 包装内置库函数
	en := native.NewNative(env).Environment()
	for {
		// 读取单词
		v, err := i.lexer.Peek(0)
		if err != nil {
			if err.Error() == "EOF" {
				if conf.ContinueWhenEOF {
					continue
				}
				if conf.BreakWhenEOF {
					break
				}
			}
			log.Error("词法分析错误: ", err)
			if conf.ContinueWhenLexerError {
				continue
			}
			if conf.BreakWhenLexerError {
				break
			}
		}
		if v == token.EOF {
			log.Info("退出: EOF")
			if conf.ContinueWhenEOF {
				continue
			}
			if conf.BreakWhenEOF {
				break
			}
		}
		ct, ok := v.(*token.CommentToken)
		if ok {
			log.Debugf("第 %d 行是注释", ct.GetLineNumber())
			// 读掉注释继续执行
			_, _ = i.lexer.Read()
			continue
		}
		// 语法分析
		t := i.parser.Parse(i.lexer)
		// 得到的 AST 不应该是空
		if t == nil {
			log.Error("AST是空")
			// 这个地方需要给参数
			if conf.ContinueWhenNullAST {
				continue
			}
			if conf.BreakWhenNullAST {
				break
			}
		}
		// 调试用
		s, err := t.GetString()
		if err != nil {
			log.Error(err)
			if conf.ContinueWhenDebugError {
				continue
			}
			if conf.BreakWhenDebugError {
				break
			}
		}
		// 确保结果是 AST
		astList, ok := t.(ast.ASTree)
		if !ok {
			log.Error("解析AST错误")
			if conf.ContinueWhenCastError {
				continue
			}
			if conf.BreakWhenCastError {
				break
			}
		}
		// 取 AST 第一个元素
		tree, err := astList.Children().Get(0)
		if err != nil {
			log.Error("检查空语句错误")
			if conf.ContinueWhenFirstError {
				continue
			}
			if conf.BreakWhenFirstError {
				break
			}
		}
		// 第一个元素是空语句则跳出
		_, ok = tree.(*ast.NullStmt)
		if ok {
			log.Debug("忽略空语句")
			if conf.ContinueNullStmt {
				continue
			}
			if conf.BreakNullStmt {
				break
			}
		}
		log.Infof("执行AST: %s", s)
		// 处理符号
		t.Lookup(en.Symbols())
		// 执行 AST
		_, err = t.Eval(en)
		// 遇到错误
		if err != nil {
			log.Error(err)
			if conf.ContinueWhenEvalError {
				continue
			}
			if conf.BreakWhenEvalError {
				break
			}
		}
	}

	// 检查主函数
	if !conf.DisableMainFunc {
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
	}

	// 确保所有协程执行完毕
	if !conf.DisableWaitForPool {
		ok := en.WaitJob()
		if ok {
			log.Info("所有协程执行结束")
		}
	}

	if i.cancel != nil {
		i.cancel()
	} else {
		log.Info("测试完成")
	}
}
