package lexer

import (
	"fmt"
	"strings"
)

func Lex(sourceCode string) ([]Token, error) {
	lines := strings.Split(sourceCode, "\n")
	tokens := []Token{}
	for lineNum, line := range lines {
		chars := []rune(line)
		col := 0
		for col < len(chars) {
			switch {
			case IsWhitespace(chars[col]):
				col++
			case IsLetter(chars[col]):
				start := col
				for col < len(chars) && (IsLetter(chars[col]) || IsDigit(chars[col])) {
					col++
				}
				value := string(chars[start:col])
				tokenType := TokenTypeIdentifier // Assume identifier initially
				switch value {
				case "parayada", "let":
					tokenType = TokenTypeKeyword
				}
				tokens = append(tokens, Token{Type: tokenType, Value: value, Line: lineNum + 1, Column: start + 1})
			case chars[col] == '"':
				start := col + 1
				col++
				for col < len(chars) && chars[col] != '"' {
					col++
				}
				if col < len(chars) && chars[col] == '"' {
					tokens = append(tokens, Token{Type: TokenTypeString, Value: string(chars[start:col]), Line: lineNum + 1, Column: start + 1})
					col++
				} else {
					return nil, fmt.Errorf("syntax error: unterminated string at line %d, column %d", lineNum+1, start)
				}
			case IsDigit(chars[col]):
				start := col
				for col < len(chars) && IsDigit(chars[col]) {
					col++
				}
				tokens = append(tokens, Token{Type: TokenTypeNumber, Value: string(chars[start:col]), Line: lineNum + 1, Column: start + 1})
			case chars[col] == ':':
				tokens = append(tokens, Token{Type: TokenTypeColon, Value: ":", Line: lineNum + 1, Column: col + 1})
				col++
			case chars[col] == '+':
				tokens = append(tokens, Token{Type: TokenTypeOperator, Value: "+", Line: lineNum + 1, Column: col + 1})
				col++
			case chars[col] == '=': // Added equals sign
				tokens = append(tokens, Token{Type: TokenTypeEquals, Value: "=", Line: lineNum + 1, Column: col + 1})
				col++
			default:
				return nil, fmt.Errorf("lexical error: unexpected character '%c' at line %d, column %d", chars[col], lineNum+1, col+1)
			}
		}
		tokens = append(tokens, Token{Type: TokenTypeNewline, Value: "\n", Line: lineNum + 1, Column: len(chars) + 1})
	}
	tokens = append(tokens, Token{Type: TokenTypeEOF, Value: "", Line: len(lines) + 1, Column: 1})
	return tokens, nil
}
