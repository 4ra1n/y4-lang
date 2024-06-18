package native

import (
	"os"

	"github.com/4ra1n/y4-lang/envir"
)

const (
	nativeExistFileFunction  = "文件存在"
	nativeWriteFileFunction  = "写文件"
	nativeReadFileFunction   = "读文件"
	nativeDeleteFileFunction = "删除文件"
)

func y4ExistFile(fileName string) int {
	file, err := os.Open(fileName)
	defer func() {
		if err = file.Close(); err != nil {
			return
		}
	}()
	if err != nil {
		if os.IsNotExist(err) {
			return envir.FALSE
		}
	}
	return envir.TRUE
}

func y4ReadFile(fileName string) string {
	f, err := os.ReadFile(fileName)
	if err != nil {
		return "<空的>"
	}
	return string(f)
}

func y4WriteFile(fileName string, data string) int {
	err := os.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		return envir.FALSE
	} else {
		return envir.TRUE
	}
}

func y4DeleteFile(fileName string) int {
	err := os.Remove(fileName)
	if err != nil {
		return envir.FALSE
	} else {
		return envir.TRUE
	}
}
