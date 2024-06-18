package token

import (
	"fmt"
	"strconv"
)

type NumberToken struct {
	lineNumber int
	value      int
}

func NewNumberToken(line int, v int) *NumberToken {
	return &NumberToken{
		lineNumber: line,
		value:      v,
	}
}

func (t *NumberToken) GetLineNumber() int {
	return t.lineNumber
}

func (t *NumberToken) IsIdentifier() bool {
	return false
}

func (t *NumberToken) IsNumber() bool {
	return true
}

func (t *NumberToken) IsString() bool {
	return false
}

func (t *NumberToken) GetNumber() (interface{}, error) {
	return t.value, nil
}

func (t *NumberToken) GetText() (string, error) {
	return strconv.Itoa(t.value), nil
}

func (t *NumberToken) String() string {
	return fmt.Sprintf("<number token><line:%d> value:%d", t.lineNumber, t.value)
}
