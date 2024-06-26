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
	DisablePreProcess      bool
	OnlyCheck              bool
)

// TestConfig
// 因为测试不走 flag 解析
// 这里需要手动初始化下
func TestConfig() {
	BreakWhenCastError = true
	BreakWhenEOF = true
	BreakWhenEvalError = true
	BreakWhenLexerError = true
	BreakWhenNullAST = true
	BreakWhenFirstError = true
	ContinueWhenDebugError = true
	ContinueNullStmt = true
}
