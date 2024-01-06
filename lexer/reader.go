package lexer

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/token"
)

func (l *Lexer) Peek(i int) (token.Token, error) {
	b, err := l.fillQueue(i)
	if err != nil {
		return nil, err
	}
	if b {
		return l.queue.Get(i)
	} else {
		return token.EOF, nil
	}
}

func (l *Lexer) Read() (token.Token, error) {
	b, err := l.fillQueue(0)
	if err != nil {
		return nil, err
	}
	if b {
		var val token.Token
		val, err = l.queue.Get(0)
		if err != nil {
			return nil, err
		}
		err = l.queue.Remove(0)
		if err != nil {
			return nil, err
		}
		return val, nil
	} else {
		return token.EOF, nil
	}
}

func (l *Lexer) readRune() (int32, error) {
	buf := make([]byte, 1)
	n, err := l.reader.Read(buf)
	if err != nil {
		return 0, err
	}
	if n != 1 {
		return 0, fmt.Errorf("no byte read")
	}
	return rune(buf[0]), nil
}

func (l *Lexer) getChar() (rune, error) {
	if l.lastChar == EMPTY {
		val, err := l.readRune()
		if err != nil {
			log.Debug(err)
			return EMPTY, err
		}
		return val, nil
	} else {
		c := l.lastChar
		l.lastChar = EMPTY
		return c, nil
	}
}

func (l *Lexer) unGetChar(c rune) {
	l.lastChar = c
}

func (l *Lexer) fillQueue(i int) (bool, error) {
	for l.queue.Length() <= i {
		t, err := l.read0()
		if err != nil {
			return false, err
		}
		if t == nil {
			continue
		}
		_, isC := t.(*token.CommentToken)
		if isC {
			continue
		}
		if t == token.EOF {
			return false, nil
		}
		l.queue.Add(t)
	}
	return true, nil
}

func (l *Lexer) read0() (token.Token, error) {
	buf := bytes.Buffer{}
	c, err := l.getChar()
	if err != nil {
		return nil, err
	}
	for isSpace(c) {
		c, err = l.getChar()
		if err != nil {
			return nil, err
		}
	}
	if c < 0 {
		return token.EOF, nil
	} else if isPound(c) {
		buf.WriteRune(c)
		c, err = l.getChar()
		if err != nil {
			return nil, err
		}
		for isLetter(c) {
			buf.WriteRune(c)
			c, err = l.getChar()
			if err != nil {
				return nil, err
			}
		}
	} else if isSlash(c) {
		buf.WriteRune(c)
		c, err = l.getChar()
		if err != nil {
			return nil, err
		}
		if isSlash(c) {
			for !isLF(c) {
				buf.WriteRune(c)
				c, err = l.getChar()
				if err != nil {
					return nil, err
				}
			}
			l.lineNumber++
			return token.NewCommentToken(l.lineNumber-1, buf.String()), nil
		} else {
			l.unGetChar(c)
			return token.NewIdentifierToken(l.lineNumber, "/"), nil
		}
	} else if isDigit(c) {
		buf.WriteRune(c)
		c, err = l.getChar()
		if err != nil {
			return nil, err
		}
		dotOnce := false
		for isDigit(c) || isDot(c) {
			buf.WriteRune(c)
			c, err = l.getChar()
			if err != nil {
				return nil, err
			}
			if isDot(c) {
				if dotOnce {
					return nil, errors.New("invalid number")
				} else {
					dotOnce = true
					buf.WriteRune(c)
					c, err = l.getChar()
					if err != nil {
						return nil, err
					}
				}
			} else {
				continue
			}
		}
	} else if isLetter(c) {
		buf.WriteRune(c)
		c, err = l.getChar()
		if err != nil {
			return nil, err
		}
		for isLetter(c) || isDigit(c) {
			buf.WriteRune(c)
			c, err = l.getChar()
			if err != nil {
				return nil, err
			}
		}
	} else if isQuota(c) {
		buf.WriteRune(c)
		c, err = l.getChar()
		if err != nil {
			return nil, err
		}
		for !isQuota(c) {
			if c == '\\' {
				c, err = l.getChar()
				if err != nil {
					return nil, err
				}
				if isQuota(c) {
					buf.WriteRune(c)
					c, err = l.getChar()
					if err != nil {
						return nil, err
					}
				}
			} else {
				buf.WriteRune(c)
				c, err = l.getChar()
				if err != nil {
					return nil, err
				}
			}
		}
		c, err = l.getChar()
		if err != nil {
			return nil, err
		}
	} else if isCR(c) {
		c, err = l.getChar()
		if err != nil {
			return nil, err
		}
		if isLF(c) {
			l.lineNumber++
			return token.NewIdentifierToken(l.lineNumber-1, token.EOL), nil
		} else {
			return nil, errors.New("error token")
		}
	} else if isLF(c) {
		l.lineNumber++
		return token.NewIdentifierToken(l.lineNumber-1, token.EOL), nil
	} else if c == '=' {
		c, err = l.getChar()
		if err != nil {
			return nil, err
		}
		if c == '=' {
			return token.NewIdentifierToken(l.lineNumber, "=="), nil
		} else {
			l.unGetChar(c)
			return token.NewIdentifierToken(l.lineNumber, "="), nil
		}
	} else if c == '>' {
		c, err = l.getChar()
		if err != nil {
			return nil, err
		}
		if c == '=' {
			return token.NewIdentifierToken(l.lineNumber, ">="), nil
		} else {
			l.unGetChar(c)
			return token.NewIdentifierToken(l.lineNumber, ">"), nil
		}
	} else if c == '!' {
		c, err = l.getChar()
		if err != nil {
			return nil, err
		}
		if c == '=' {
			return token.NewIdentifierToken(l.lineNumber, "!="), nil
		} else {
			l.unGetChar(c)
			return token.NewIdentifierToken(l.lineNumber, "!"), nil
		}
	} else if c == '<' {
		c, err = l.getChar()
		if err != nil {
			return nil, err
		}
		if c == '=' {
			return token.NewIdentifierToken(l.lineNumber, "<="), nil
		} else {
			l.unGetChar(c)
			return token.NewIdentifierToken(l.lineNumber, "<"), nil
		}
	} else if isCalc(c) || isBracket(c) || isSem(c) || isComma(c) {
		return token.NewIdentifierToken(l.lineNumber, string(c)), nil
	} else {
		return nil, errors.New("error token")
	}

	if c > 0 {
		l.unGetChar(c)
	}
	temp := buf.String()
	var firstRune rune
	for _, r := range temp {
		firstRune = r
		break
	}
	if isLetter(firstRune) || isPound(firstRune) {
		if strings.ToLower(temp) == "true" {
			return token.NewNumberToken(l.lineNumber, 1), nil
		}
		if strings.ToLower(temp) == "false" {
			return token.NewNumberToken(l.lineNumber, 0), nil
		}
		if strings.ToLower(temp) == "null" {
			return token.NewNumberToken(l.lineNumber, 0), nil
		}
		return token.NewIdentifierToken(l.lineNumber, temp), nil
	} else if isDigit(firstRune) {
		if strings.Contains(temp, ".") {
			v, err := strconv.ParseFloat(temp, 64)
			if err != nil {
				return nil, err
			}
			return token.NewFloatToken(l.lineNumber, v), nil
		} else {
			v, err := strconv.Atoi(temp)
			if err != nil {
				return nil, err
			}
			return token.NewNumberToken(l.lineNumber, v), nil
		}
	} else if isQuota(firstRune) {
		temp = temp[1:]
		return token.NewStringToken(l.lineNumber, temp), nil
	} else {
		return nil, errors.New("error token")
	}
}
