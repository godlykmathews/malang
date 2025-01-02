package lexer

func IsWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\r'
}

func IsLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func IsDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
