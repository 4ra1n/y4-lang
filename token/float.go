package token

import (
	"fmt"
)

type FloatToken struct {
	lineNumber int
	value      float64
}

func NewFloatToken(line int, v float64) *FloatToken {
	return &FloatToken{
		lineNumber: line,
		value:      v,
	}
}

func (f *FloatToken) IsIdentifier() bool {
	return false
}

func (f *FloatToken) IsNumber() bool {
	return true
}

func (f *FloatToken) IsString() bool {
	return false
}

func (f *FloatToken) GetLineNumber() int {
	return f.lineNumber
}

func (f *FloatToken) GetNumber() (interface{}, error) {
	return f.value, nil
}

func (f *FloatToken) GetText() (string, error) {
	return fmt.Sprintf("%f", f.value), nil
}

func (f *FloatToken) String() string {
	return fmt.Sprintf("<float token><line:%d> value:%f", f.lineNumber, f.value)
}
