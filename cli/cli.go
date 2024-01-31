package cli

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/4ra1n/y4-lang/color"
	"github.com/4ra1n/y4-lang/conf"
	"github.com/4ra1n/y4-lang/log"
)

var (
	// y4lang 1.y4 2.y4 支持多个文件一起解释执行
	filePath []string
	// QuietFlag 这个是安静模式
	QuietFlag bool
	// 打印详细构建信息
	versionFlag bool
	// 帮助
	helpFlag bool
	// 初始化环境大小
	envSize int
	// 初始化协程池大小
	poolSize int
	// 日志级别
	logLevel string
)

func Start(cancel context.CancelFunc) {
	parseArgs()
	setLogLevel()
	start(cancel)
	// 正常执行完应该结束掉
	if cancel != nil {
		cancel()
	}
}

func setLogLevel() {
	logLevel = strings.TrimSpace(strings.ToLower(logLevel))
	// 如果开了安静模式不用打具体日志
	if QuietFlag {
		log.SetLevel(log.ErrorLevel)
		return
	}
	switch logLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
		break
	case "info":
		log.SetLevel(log.InfoLevel)
		break
	case "warn":
		log.SetLevel(log.WarnLevel)
		break
	case "error":
		log.SetLevel(log.ErrorLevel)
		break
	case "disabled":
		log.SetLevel(log.Disabled)
		break
	default:
		color.RedPrintln("error log level")
		os.Exit(0)
	}
}

func parseArgs() {
	// 基本 FLAG
	flag.IntVar(&envSize, "env-size", 0, "设置初始环境大小")
	flag.IntVar(&poolSize, "pool-size", 0, "设置线程池大小")
	flag.BoolVar(&versionFlag, "version", false, "打印具体版本")
	flag.BoolVar(&QuietFlag, "quiet", false, "安静模式")
	flag.BoolVar(&helpFlag, "h", false, "打印帮助信息")
	flag.StringVar(&logLevel, "log-level", "error",
		"打印日志级别 (debug, info, warn, error, disabled)")
	// 进阶 FLAG
	flag.BoolVar(&conf.ContinueWhenCastError, "xcwce", false, "解析AST报错继续")
	flag.BoolVar(&conf.BreakWhenCastError, "xbwce", true, "解析AST报错跳出")

	flag.BoolVar(&conf.ContinueWhenEOF, "xcwe", false, "遇到EOF报错继续")
	flag.BoolVar(&conf.BreakWhenEOF, "xbwe", true, "遇到EOF报错跳出")

	flag.BoolVar(&conf.ContinueWhenEvalError, "xcwee", false, "解析执行报错继续")
	flag.BoolVar(&conf.BreakWhenEvalError, "xbwee", true, "解析执行报错跳出")

	flag.BoolVar(&conf.ContinueWhenLexerError, "xcwle", false, "词法分析报错继续")
	flag.BoolVar(&conf.BreakWhenLexerError, "xbwle", true, "词法分析报错跳出")

	flag.BoolVar(&conf.ContinueWhenNullAST, "xcwna", false, "空AST继续")
	flag.BoolVar(&conf.BreakWhenNullAST, "xbwna", true, "空AST跳出")

	flag.BoolVar(&conf.ContinueWhenFirstError, "xcwfe", false, "取第一子元素错误继续")
	flag.BoolVar(&conf.BreakWhenFirstError, "xbwfe", true, "取第一子元素错误跳出")

	flag.BoolVar(&conf.ContinueWhenDebugError, "xcwde", true, "调试信息报错继续")
	flag.BoolVar(&conf.BreakWhenDebugError, "xbwde", false, "调试信息报错跳出")

	flag.BoolVar(&conf.ContinueNullStmt, "xcns", true, "空语句继续")
	flag.BoolVar(&conf.BreakNullStmt, "xbns", false, "空语句跳出")

	flag.BoolVar(&conf.DisableMainFunc, "xdmf", false, "是否禁用主函数")
	flag.BoolVar(&conf.DisableWaitForPool, "xdwfp", false, "禁用等待协程池任务完毕")
	flag.BoolVar(&conf.OnlyCheck, "xoc", false, "是否只检查不执行")
	flag.BoolVar(&conf.DisablePreProcess, "xdpp", false, "禁用预处理")

	flag.Usage = func() {
		fmt.Println("y4-lang 帮助信息:")
		flag.PrintDefaults()
	}

	flag.Parse()
	if helpFlag {
		PrintLogo()
		flag.Usage()
		os.Exit(0)
	}
	if versionFlag {
		PrintLogo()
		color.GreenPrintf("构建日期: %s\n", BuildTime)
		os.Exit(0)
	}
	if !QuietFlag {
		PrintLogo()
	}

	files := flag.Args()
	if len(files) > 0 {
		filePath = files
	} else {
		color.RedPrintln("没有输入文件")
		os.Exit(0)
	}
}

func PrintLogo() {
	color.GreenPrintln(Logo)
	color.YellowPrintf("%s %s\n%s\n项目地址: %s\n",
		Name, Version, Desc, ProjectUrl)
}
