package cli

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/4ra1n/y4-lang/color"
	"github.com/4ra1n/y4-lang/log"
)

var (
	QuietFlag   bool
	versionFlag bool
	helpFlag    bool
	envSize     int
	poolSize    int
	filePath    string
	logLevel    string
)

func Start(cancel context.CancelFunc) {
	parseArgs()
	setLogLevel()
	start(cancel)
	if cancel != nil {
		cancel()
	}
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
		os.Exit(-1)
	}
}

func parseArgs() {
	flag.StringVar(&filePath, "f", "", "specify the file path")
	flag.IntVar(&envSize, "env-size", 0, "set environment size")
	flag.IntVar(&poolSize, "pool-size", 0, "set threads pool size")
	flag.BoolVar(&versionFlag, "version", false, "print the version")
	flag.BoolVar(&QuietFlag, "quiet", false, "quiet mode (hide logo print)")
	flag.BoolVar(&helpFlag, "h", false, "print help information")
	flag.StringVar(&logLevel, "log-level", "error",
		"specify the log level ('debug', 'info', 'warn', 'error', 'disabled')")

	flag.Usage = func() {
		fmt.Println("y4-lang usage:")
		flag.PrintDefaults()
	}

	flag.Parse()
	if helpFlag {
		PrintLogo()
		flag.Usage()
		os.Exit(-1)
	}
	if versionFlag {
		PrintLogo()
		color.GreenPrintf("build time: %s\n", BuildTime)
		os.Exit(-1)
	}
	if !QuietFlag {
		PrintLogo()
	}
}

func PrintLogo() {
	color.GreenPrintln(Logo)
	color.YellowPrintf("%s %s\n%s\nproject: %s\n",
		Name, Version, Desc, ProjectUrl)
}
