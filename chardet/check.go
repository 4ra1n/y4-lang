package chardet

import "io/ioutil"

func DetectFileEncoding(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	detector := NewTextDetector()
	result, err := detector.DetectBest(data)
	if err != nil {
		return "", err
	}
	return result.Charset, nil
}
