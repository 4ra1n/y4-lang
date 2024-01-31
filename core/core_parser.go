package core

import (
	"github.com/4ra1n/y4-lang/ast"
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/parser"
	"github.com/4ra1n/y4-lang/token"
)

// 核心语法规则
// 除非你明白自己在做什么
// 否则不要修改
var (
	reserved  = base.NewHashSet[string]()
	operators = parser.NewOperators()
	expr0     = parser.RuleNoType()

	// primary : "(" expr ")" | number | identifier | string
	primary = parser.RuleWithType("primary_expr").Or(
		parser.RuleNoType().Sep("(").Ast(expr0).Sep(")"),
		parser.RuleNoType().NumberWithType("number_literal"),
		parser.RuleNoType().IdentifierWithType("name", reserved),
		parser.RuleNoType().StringWithType("string_literal"))

	// factor : - primary | primary
	factor = parser.RuleNoType().Or(
		parser.RuleWithType("negative_expr").Sep("-").Ast(primary),
		parser.RuleWithType("not_expr").Sep("!").Ast(primary),
		primary)

	// expr : factor { operator factor }
	expr       = expr0.ExpressionWithType("binary_expr", factor, operators)
	statement0 = parser.RuleNoType()

	// block: "{" [ statement ] { ( ";" | EOL ) [statement] } "}"
	block = parser.RuleWithType("block_stmt").Sep("{").Option(statement0).Repeat(
		parser.RuleNoType().Sep(";", token.EOL).Option(statement0)).Sep("}")

	// simple: expr
	simple = parser.RuleWithType("primary_expr").Ast(expr)

	// continue_statement : "continue"
	continueStmt = parser.RuleWithType("continue_stmt").Sep("继续")

	// break_statement : "break"
	breakStmt = parser.RuleWithType("break_stmt").Sep("跳出")

	// statement : "if" expr block [ "else" block ]
	//             | "while" expr block | simple
	//			   | "for" expr ; expr ; expr block
	//             | "return" factor
	//             | "continue"
	// 			   | "break"
	//             | simple (expr)
	statement = statement0.Or(
		parser.RuleWithType("if_stmt").Sep(ifKey).Ast(expr).Ast(block).Option(
			parser.RuleNoType().Sep(elseKey).Ast(block)),
		parser.RuleWithType("while_stmt").Sep(whileKey).Ast(expr).Ast(block),
		parser.RuleWithType("for_stmt").Sep(forKey).
			Ast(expr).Sep(";"). // init
			Ast(expr).Sep(";"). // condition
			Ast(expr).          // iteration
			Ast(block),
		parser.RuleWithType("return_stmt").Sep(returnKey).Ast(expr),
		parser.RuleWithType("go_stmt").Sep(goKey).Ast(factor),
		continueStmt,
		breakStmt,
		simple,
	)

	// include : "#include" identifier
	include = parser.RuleWithType("include_stmt").
		Sep(includeKey).StringWithType("string_literal")

	// program : ( include | statement | null ) ( ";" | EOL )
	program = parser.RuleNoType().Or(
		include, statement, parser.RuleWithType("null_stmt")).Sep(";", token.EOL)

	param = parser.RuleNoType().IdentifierNoType(reserved)
	// params : param { ","  param }
	params = parser.RuleWithType("parameter_list").Ast(param).Repeat(
		parser.RuleNoType().Sep(",").Ast(param))

	// param_list : params "(" [ params ] ")"
	paramList = parser.RuleNoType().Sep("(").Maybe(params).Sep(")")

	// def : "def" identifier param_list block
	def = parser.RuleWithType("def_stmt").Sep(defKey).
		IdentifierNoType(reserved).Ast(paramList).Ast(block)

	// args : expr { "," expr }
	args = parser.RuleWithType("arguments").Ast(expr).
		Repeat(parser.RuleNoType().Sep(",").Ast(expr))

	// postfix : "(" [ args ] ")"
	postfix = parser.RuleNoType().Sep("(").Maybe(args).Sep(")")

	// element : expr { "," expr }
	elements = parser.RuleWithType("array_literal").Ast(expr).Repeat(
		parser.RuleNoType().Sep(",").Ast(expr))
)

type Parser struct {
}

func NewCoreParser() *Parser {
	addReserved()
	addOperators()
	arrayRule()
	functionRule()
	return &Parser{}
}

func functionRule() {
	// primary : ( "[" [element] "]" | "(" expr ")"
	//           | number | identifier | string ) { postfix }
	primary.Repeat(postfix)
	// simple : expr [args]
	simple.Option(args)
	// program : [ def | statement ] ( ";" | EOL )
	program.InsertChoice(def)
}

func arrayRule() {
	// primary : "[" [element] "]"
	//           | "(" expr ")" | number | identifier | string
	primary.InsertChoice(parser.RuleNoType().Sep("[").Maybe(elements).Sep("]"))
	// postfix : "(" [ args ] ")" | "[" expr "]"
	postfix.InsertChoice(parser.RuleWithType("array_ref").Sep("[").Ast(expr).Sep("]"))
}

func addOperators() {
	operators.Add("=", 1, parser.RIGHT)
	operators.Add("==", 2, parser.LEFT)
	operators.Add("!=", 2, parser.LEFT)
	operators.Add(">", 2, parser.LEFT)
	operators.Add(">=", 2, parser.LEFT)
	operators.Add("<", 2, parser.LEFT)
	operators.Add("<=", 2, parser.LEFT)
	operators.Add("+", 3, parser.LEFT)
	operators.Add("-", 3, parser.LEFT)
	operators.Add("*", 4, parser.LEFT)
	operators.Add("/", 4, parser.LEFT)
	operators.Add("%", 4, parser.LEFT)
}

func addReserved() {
	reserved.Add(";")
	reserved.Add("}")
	reserved.Add("]")
	reserved.Add(")")
	reserved.Add(token.EOL)
}

func (c *Parser) Parse(l *lexer.Lexer) ast.ASTree {
	return program.Parse(l)
}
