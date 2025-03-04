package lexer

import (
	"fmt"
	"strings"
)

func Lex(input string) []Token {
	tokens := []Token{}
	line := 1
	col := 1

	keywords := map[string]string{
		"parayu":          TokParayu,
		"kelk":            TokKelk,
		"ith_sheriyano":   TokAadhyamayi,
		"enkil":           TokAthengil,
		"alle":            TokIlla,
		"ellam_sheriyano": TokEllamSheriyano,
		"oron_ayi":        TokOnninuMumbu,
		"edukk":           TokEdukk,
	}

	operators := []string{"==", "<", "=", "+", "-", "*", "/"} // keep longer operators first

	for i := 0; i < len(input); {
		char := input[i]

		// Skip whitespace
		if char == ' ' || char == '\t' {
			i++
			col++
			continue
		}

		if char == '\n' {
			line++
			col = 1
			i++
			continue
		}

		// String literals
		if char == '"' {
			start := i + 1
			for i++; i < len(input) && input[i] != '"'; i++ {
				if input[i] == '\n' {
					line++
					col = 1
				} else {
					col++
				}
			}
			if i < len(input) && input[i] == '"' {
				tokens = append(tokens, Token{Type: TokString, Value: input[start:i], Line: line, Col: col})
				col += i - start + 2
				i++

			} else {
				panic(fmt.Sprintf("Unterminated string literal at line %d, col %d", line, col))
			}

			continue
		}

		// Identifiers and Keywords
		if IsAlpha(char) {
			start := i
			for i < len(input) && (IsAlpha(input[i]) || IsDigit(input[i]) || input[i] == '_') {
				i++
			}
			value := input[start:i]
			if tokenType, ok := keywords[value]; ok {
				tokens = append(tokens, Token{Type: tokenType, Value: value, Line: line, Col: col})
			} else {
				tokens = append(tokens, Token{Type: TokIdentifier, Value: value, Line: line, Col: col})
			}
			col += i - start
			continue
		}

		// Numbers
		if IsDigit(char) {
			start := i
			for i < len(input) && IsDigit(input[i]) {
				i++
			}
			tokens = append(tokens, Token{Type: TokInteger, Value: input[start:i], Line: line, Col: col})
			col += i - start
			continue
		}

		// Operators
		matchedOperator := false
		for _, op := range operators {
			if strings.HasPrefix(input[i:], op) {
				tokens = append(tokens, Token{Type: TokOperator, Value: op, Line: line, Col: col})
				i += len(op)
				col += len(op)
				matchedOperator = true
				break
			}
		}
		if matchedOperator {
			continue
		}

		// Range ..
		if i+1 < len(input) && input[i] == '.' && input[i+1] == '.' {
			tokens = append(tokens, Token{Type: TokRange, Value: "..", Line: line, Col: col})
			i += 2
			col += 2
			continue
		}

		// Parentheses and Braces, Comma
		switch char {
		case '(':
			tokens = append(tokens, Token{Type: TokLParen, Value: "(", Line: line, Col: col})
			i++
			col++
		case ')':
			tokens = append(tokens, Token{Type: TokRParen, Value: ")", Line: line, Col: col})
			i++
			col++
		case '{':
			tokens = append(tokens, Token{Type: TokLBrace, Value: "{", Line: line, Col: col})
			i++
			col++
		case '}':
			tokens = append(tokens, Token{Type: TokRBrace, Value: "}", Line: line, Col: col})
			i++
			col++
		case ',':
			tokens = append(tokens, Token{Type: TokComma, Value: ",", Line: line, Col: col})
			i++
			col++

		default:
			panic(fmt.Sprintf("Unexpected character '%c' at line %d, col %d", char, line, col))
		}
	}

	tokens = append(tokens, Token{Type: TokEOF, Value: "", Line: line, Col: col})
	return tokens
}
