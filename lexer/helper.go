package lexer

func isLetter(c byte) bool {
	return ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') || c == '.' || c == '_' ||
		// CHINESE
		(c >= 0xe0 && c <= 0xef) || (c >= 0x80 && c <= 0xbf)
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func isSpace(c byte) bool {
	return c == ' ' || c == '\t'
}

func isCR(c byte) bool {
	return c == '\r'
}

func isLF(c byte) bool {
	return c == '\n'
}

func isCalc(c byte) bool {
	return c == '+' || c == '-' || c == '*' || c == '/' || c == '%'
}

func isBracket(c byte) bool {
	return c == '{' || c == '}' || c == '(' || c == ')' || c == '[' || c == ']'
}

func isComma(c byte) bool {
	return c == ','
}

func isPound(c byte) bool {
	return c == '#'
}

func isSem(c byte) bool {
	return c == ';'
}

func isQuota(c byte) bool {
	return c == '"'
}

func isSlash(c byte) bool {
	return c == '/'
}

func isDot(c byte) bool {
	return c == '.'
}
