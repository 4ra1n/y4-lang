package token

import (
	"errors"
	"fmt"
)

type CommentToken struct {
	lineNumber int
	comment    string
}

func NewCommentToken(line int, comment string) *CommentToken {
	return &CommentToken{
		lineNumber: line,
		comment:    comment,
	}
}

func (t *CommentToken) GetLineNumber() int {
	return t.lineNumber
}

func (t *CommentToken) IsIdentifier() bool {
	return false
}

func (t *CommentToken) IsNumber() bool {
	return false
}

func (t *CommentToken) IsString() bool {
	return false
}

func (t *CommentToken) GetNumber() (interface{}, error) {
	return 0, errors.New("not number token")
}

func (t *CommentToken) GetText() (string, error) {
	return NULL, errors.New("not a text")
}

func (t *CommentToken) String() string {
	return fmt.Sprintf("<comment token><line:%d> value:%s", t.lineNumber, t.comment)
}
