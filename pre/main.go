package pre

import (
	"errors"
	"os"
	"regexp"
)

// SearchMain
// files input file list
// return file name
func SearchMain(files []string) (string, error) {
	for _, file := range files {
		data, err := os.ReadFile(file)
		pattern := `def\s+main\s*\(\s*\)\s*{[^}]*}`
		re := regexp.MustCompile(pattern)
		found := re.MatchString(string(data))
		if found {
			return file, nil
		}
		if err != nil {
			return "", errors.New("search main error")
		}
	}
	return "", errors.New("search main error")
}
