package lexer

func isLetter(c rune) bool {
	return ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') || isDot(c)
}

func isDigit(c rune) bool {
	return '0' <= c && c <= '9'
}

func isSpace(c rune) bool {
	return c == ' ' || c == '\t'
}

func isCR(c rune) bool {
	return c == '\r'
}

func isLF(c rune) bool {
	return c == '\n'
}

func isCalc(c rune) bool {
	return c == '+' || c == '-' || c == '*' || c == '/' || c == '%'
}

func isBracket(c rune) bool {
	return c == '{' || c == '}' || c == '(' || c == ')' || c == '[' || c == ']'
}

func isComma(c rune) bool {
	return c == ','
}

func isPound(c rune) bool {
	return c == '#'
}

func isSem(c rune) bool {
	return c == ';'
}

func isQuota(c rune) bool {
	return c == '"'
}

func isSlash(c rune) bool {
	return c == '/'
}

func isDot(c rune) bool {
	return c == '.'
}
