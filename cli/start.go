package cli

import (
	"os"
	"strings"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/pre"
)

func start() {
	log.Debug("start y4-lang")

	// check nil
	filePath = strings.TrimSpace(filePath)
	if filePath == "" {
		log.Error("file name is null")
		return
	}

	// check extension name
	if strings.HasSuffix(strings.ToLower(filePath), ".y4") {
		log.Errorf("file extension must be y4")
		return
	}

	// check file exist
	file, err := os.Open(filePath)
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
	ip := pre.NewIncludeProcessor(filePath)
	newReader := ip.Process()

	// new lexer
	l := lexer.NewLexer(newReader)
	// new interpreter
	i := core.NewInterpreter(l)
	// start
	i.Start()
}
