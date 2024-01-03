package token

import (
	"errors"
	"fmt"
)

type BaseToken struct {
	lineNumber int
}

func NewBaseToken(line int) *BaseToken {
	return &BaseToken{
		lineNumber: line,
	}
}

func (t *BaseToken) GetLineNumber() int {
	return t.lineNumber
}

func (t *BaseToken) IsIdentifier() bool {
	return false
}

func (t *BaseToken) IsNumber() bool {
	return false
}

func (t *BaseToken) IsString() bool {
	return false
}

func (t *BaseToken) GetNumber() (interface{}, error) {
	return 0, errors.New("not number token")
}

func (t *BaseToken) GetText() (string, error) {
	return NULL, errors.New("not a text")
}

func (t *BaseToken) String() string {
	return fmt.Sprintf("<base token><line:%d> value:null", t.lineNumber)
}
