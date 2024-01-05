package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
