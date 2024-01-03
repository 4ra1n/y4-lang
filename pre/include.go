package pre

import (
	"bytes"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/4ra1n/y4-lang/log"
)

type IncludePreprocessor struct {
	FileName string
	fileData string
}

func NewIncludeProcessor(fileName string) *IncludePreprocessor {
	file, err := os.Open(fileName)
	defer func() {
		if err = file.Close(); err != nil {
			log.Errorf("close file error: %s", err.Error())
		}
	}()
	if err != nil {
		if os.IsNotExist(err) {
			log.Debugf("file not exist: %s", fileName)
		} else {
			log.Errorf("open file error: %s", err.Error())
		}
		return nil
	} else {
		data, err := io.ReadAll(file)
		if err != nil {
			log.Errorf("read file error: %s", err.Error())
		}
		return &IncludePreprocessor{
			FileName: fileName,
			fileData: string(data),
		}
	}
}

func (ip *IncludePreprocessor) Process() io.Reader {
	re := regexp.MustCompile(`#include\s+"(.*?)"`)
	matches := re.FindAllStringSubmatch(ip.fileData, -1)

	for _, match := range matches {
		var item = match[1]
		if !strings.HasSuffix(item, ".y4") {
			item = item + ".y4"
		}

		file, err := os.Open(item)
		if err != nil {
			if os.IsNotExist(err) {
				log.Debugf("y4 file not found: %s", item)
			} else {
				log.Debugf("open file error: %s", err.Error())
			}
			// ignore native lib
			return bytes.NewReader([]byte(ip.fileData))
		}

		data, err := io.ReadAll(file)
		if err != nil {
			log.Error(err)
			return nil
		}
		dataStr := string(data)

		ip.fileData = strings.ReplaceAll(ip.fileData, match[0], dataStr)

		if err = file.Close(); err != nil {
			log.Errorf("close file error: %s", err.Error())
			return nil
		}
	}

	return bytes.NewReader([]byte(ip.fileData))
}
