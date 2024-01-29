package cli

import (
	"context"
	"os"
	"strings"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/pre"
)

func start(cancel context.CancelFunc) {
	log.Debug("start y4-lang")

	var (
		err      error
		mainFile string
	)
	if len(filePath) > 1 {
		// find main method
		// if not found return error
		mainFile, err = pre.SearchMain(filePath)
		if err != nil {
			log.Error(err)
			return
		}
	} else if len(filePath) == 1 {
		mainFile = filePath[0]
	} else {
		log.Error("error file input")
		return
	}

	// check nil
	mainFile = strings.TrimSpace(mainFile)
	if mainFile == "" {
		log.Error("file name is null")
		return
	}

	// check extension name (allow *.y4)
	if !strings.HasSuffix(strings.ToLower(mainFile), ".y4") {
		log.Errorf("file extension must be y4")
		return
	}

	// check file exist
	file, err := os.Open(mainFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Errorf("file not exist: %s", filePath)
		} else {
			log.Errorf("open file error: %s", err.Error())
		}
		return
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Errorf("close file error: %s", err.Error())
		}
	}()

	// preprocessor
	ip := pre.NewIncludeProcessor(mainFile)
	newReader := ip.Process()

	// new lexer
	l := lexer.NewLexer(newReader)
	// new interpreter
	i := core.NewInterpreter(l, cancel)

	if envSize != 0 {
		log.Debugf("set env size %d", envSize)
		i.SetEnvSize(envSize)
	}

	if poolSize != 0 {
		log.Debugf("set pool size %d", poolSize)
		i.SetPoolSize(poolSize)
	}

	// start
	i.Start()
}
