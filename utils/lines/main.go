package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	var totalLines int
	root := "."
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			fileLines, err := countLines(path)
			if err != nil {
				return err
			}
			totalLines += fileLines
		}
		return nil
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("total lines in go files:", totalLines)
	_ = replaceInReadme(totalLines)
}

func replaceInReadme(totalLines int) error {
	filename := "README.md"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	re := regexp.MustCompile(`!\[]\(https://img.shields.io/badge/Code%20Lines-\d+-blue\)`)
	newContent := re.ReplaceAllString(string(content), fmt.Sprintf(
		"![Code Lines](https://img.shields.io/badge/Code%%20Lines-%d-blue)", totalLines))
	return os.WriteFile(filename, []byte(newContent), 0644)
}

func countLines(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}
	return lines, scanner.Err()
}
