package cli

import (
	"context"
	"flag"
	"os"
	"strings"

	"github.com/4ra1n/y4-lang/color"
	"github.com/4ra1n/y4-lang/log"
)

var (
	versionFlag bool
	helpFlag    bool
	filePath    string
	logLevel    string
)

const (
	ErrLogLevel = -1
	ErrFlag     = 0
)

func Start(cancel context.CancelFunc) {
	parseArgs()
	setLogLevel()
	start(cancel)
}

func setLogLevel() {
	logLevel = strings.TrimSpace(strings.ToLower(logLevel))
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
		os.Exit(ErrLogLevel)
	}
}

func parseArgs() {
	flag.StringVar(&filePath, "f", "", "specify the file path")
	flag.BoolVar(&versionFlag, "version", false, "print the version")
	flag.BoolVar(&versionFlag, "v", false, "print the version")
	flag.BoolVar(&helpFlag, "h", false, "print help information")
	flag.StringVar(&logLevel, "log-level", "error",
		"specify the log level ('debug', 'info', 'warn', 'error', 'disabled')")
	flag.Parse()
	if helpFlag {
		flag.Usage()
		os.Exit(ErrFlag)
	}
	if versionFlag {
		color.GreenPrintf("build time: %s\n", BuildTime)
		os.Exit(ErrFlag)
	}
}

func PrintLogo() {
	color.GreenPrintln(Logo)
	color.YellowPrintf("%s %s\n%s\nproject: %s\n",
		Name, Version, Desc, ProjectUrl)
}
