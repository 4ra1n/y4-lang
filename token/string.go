package token

import (
	"errors"
	"fmt"
)

type StringToken struct {
	lineNumber int
	literal    string
}

func NewStringToken(line int, s string) *StringToken {
	return &StringToken{
		lineNumber: line,
		literal:    s,
	}
}

func (t *StringToken) GetLineNumber() int {
	return t.lineNumber
}

func (t *StringToken) IsIdentifier() bool {
	return false
}

func (t *StringToken) IsNumber() bool {
	return false
}

func (t *StringToken) IsString() bool {
	return true
}

func (t *StringToken) GetNumber() (interface{}, error) {
	return 0, errors.New("not number token")
}

func (t *StringToken) GetText() (string, error) {
	return t.literal, nil
}

func (t *StringToken) String() string {
	return fmt.Sprintf("<string token><line:%d> value:%s", t.lineNumber, t.literal)
}
