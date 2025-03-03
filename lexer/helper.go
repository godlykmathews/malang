package lexer

func IsWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\r'
}

func IsLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func IsDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func IsAlpha(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}
