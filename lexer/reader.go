package lexer

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

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

func (l *Lexer) readByte() (byte, error) {
	buf := make([]byte, 1)
	n, err := l.reader.Read(buf)
	if err != nil {
		return 0, err
	}
	if n != 1 {
		return 0, fmt.Errorf("no byte read")
	}
	return buf[0], nil
}

func (l *Lexer) getChar() (byte, error) {
	if l.lastChar == EMPTY {
		val, err := l.readByte()
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

func (l *Lexer) unGetChar(c byte) {
	l.lastChar = c
}

func (l *Lexer) fillQueue(i int) (bool, error) {
	for l.queue.Length() <= i {
		t, err := l.read0()
		log.Debug(t)
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
		buf.WriteByte(c)
		c, err = l.getChar()
		if err != nil {
			return nil, err
		}
		for isLetter(c) {
			buf.WriteByte(c)
			c, err = l.getChar()
			if err != nil {
				return nil, err
			}
		}
	} else if isSlash(c) {
		buf.WriteByte(c)
		c, err = l.getChar()
		if err != nil {
			return nil, err
		}
		if isSlash(c) {
			for !isLF(c) {
				buf.WriteByte(c)
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
		buf.WriteByte(c)
		c, err = l.getChar()
		if err != nil {
			return nil, err
		}
		dotOnce := false
		for isDigit(c) || isDot(c) {
			buf.WriteByte(c)
			c, err = l.getChar()
			if err != nil {
				return nil, err
			}
			if isDot(c) {
				if dotOnce {
					return nil, errors.New("invalid number")
				} else {
					dotOnce = true
					buf.WriteByte(c)
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
		buf.WriteByte(c)
		c, err = l.getChar()
		if err != nil {
			return nil, err
		}
		for isLetter(c) || isDigit(c) {
			buf.WriteByte(c)
			c, err = l.getChar()
			if err != nil {
				return nil, err
			}
		}
	} else if isQuota(c) {
		buf.WriteByte(c)
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
					buf.WriteByte(c)
					c, err = l.getChar()
					if err != nil {
						return nil, err
					}
				} else {
					if c == 'n' {
						buf.WriteByte(byte('\n'))
						c, err = l.getChar()
						if err != nil {
							return nil, err
						}
					} else if c == 'r' {
						buf.WriteByte(byte('\r'))
						c, err = l.getChar()
						if err != nil {
							return nil, err
						}
					} else if c == 't' {
						buf.WriteByte(byte('\t'))
						c, err = l.getChar()
						if err != nil {
							return nil, err
						}
					}
				}
			} else {
				buf.WriteByte(c)
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

	var (
		firstRune  byte
		secondRune byte
		thirdRune  byte
	)
	if len(temp) == 1 {
		firstRune = temp[0]
		secondRune = 0
		thirdRune = 0
	} else if len(temp) == 2 {
		firstRune = temp[0]
		secondRune = temp[1]
		thirdRune = 0
	} else if len(temp) == 0 {
		firstRune = 0
		secondRune = 0
		thirdRune = 0
	} else {
		firstRune = temp[0]
		secondRune = temp[1]
		thirdRune = temp[2]
	}

	det := []byte{firstRune, secondRune, thirdRune}
	if utf8.Valid(det) {
		r, _ := utf8.DecodeRune(det)
		if isSimplifiedChinese(r) {
			if strings.ToLower(temp) == "真的" {
				return token.NewNumberToken(l.lineNumber, 1), nil
			} else if strings.ToLower(temp) == "假的" {
				return token.NewNumberToken(l.lineNumber, 0), nil
			} else if strings.ToLower(temp) == "空的" {
				return token.NewNumberToken(l.lineNumber, 0), nil
			}
			return token.NewIdentifierToken(l.lineNumber, temp), nil
		}
	}

	if isPound(firstRune) {
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
		return token.NewIdentifierToken(l.lineNumber, temp), nil
	}
}
