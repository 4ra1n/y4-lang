package token

import (
	"errors"
	"fmt"
)

type IdentifierToken struct {
	lineNumber int
	text       string
}

func NewIdentifierToken(line int, t string) *IdentifierToken {
	return &IdentifierToken{
		lineNumber: line,
		text:       t,
	}
}

func (t *IdentifierToken) GetLineNumber() int {
	return t.lineNumber
}

func (t *IdentifierToken) IsIdentifier() bool {
	return true
}

func (t *IdentifierToken) IsNumber() bool {
	return false
}

func (t *IdentifierToken) IsString() bool {
	return false
}

func (t *IdentifierToken) GetNumber() (interface{}, error) {
	return 0, errors.New("not number token")
}

func (t *IdentifierToken) GetText() (string, error) {
	return t.text, nil
}

func (t *IdentifierToken) String() string {
	return fmt.Sprintf("<identifier token><line:%d> value:%s", t.lineNumber, t.text)
}
