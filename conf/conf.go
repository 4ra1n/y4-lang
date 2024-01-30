package conf

// 解释器进阶参数配置
var (
	ContinueWhenLexerError bool
	BreakWhenLexerError    bool
	ContinueWhenEOF        bool
	BreakWhenEOF           bool
	ContinueWhenNullAST    bool
	BreakWhenNullAST       bool
	ContinueWhenDebugError bool
	BreakWhenDebugError    bool
	ContinueWhenCastError  bool
	BreakWhenCastError     bool
	ContinueWhenFirstError bool
	BreakWhenFirstError    bool
	ContinueWhenEvalError  bool
	BreakWhenEvalError     bool
	ContinueNullStmt       bool
	BreakNullStmt          bool
	DisableMainFunc        bool
	DisableWaitForPool     bool
)
