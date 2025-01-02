package parser

import (
	"fmt"

	"github.com/Rohith04MVK/malang/ast"
	"github.com/Rohith04MVK/malang/lexer"
)

func Parse(tokens []lexer.Token) ([]ast.Node, error) {
	// Example (very simplified and incomplete):
	nodes := []ast.Node{}
	i := 0
	for i < len(tokens)-1 && tokens[i].Type != lexer.TokenTypeEOF {
		switch tokens[i].Type {
		case lexer.TokenTypeKeyword:
			switch tokens[i].Value {
			case "let": // Parse variable declaration
				if i+1 < len(tokens) && tokens[i+1].Type == lexer.TokenTypeIdentifier &&
					i+2 < len(tokens) && tokens[i+2].Type == lexer.TokenTypeEquals {
					identifier := tokens[i+1].Value
					// Parse the expression after the equals sign
					expr, steps, err := parseExpression(tokens[i+3:])
					if err != nil {
						return nil, err
					}
					nodes = append(nodes, ast.VariableDeclarationNode{Identifier: identifier, Value: expr})
					i += 3 + steps
				} else {
					return nil, fmt.Errorf("syntax error in variable declaration at line %d", tokens[i].Line)
				}

			case "parayada":
				if i+1 < len(tokens) && tokens[i+1].Type == lexer.TokenTypeString || tokens[i+1].Type == lexer.TokenTypeIdentifier || tokens[i+1].Type == lexer.TokenTypeNumber {
					nodes = append(nodes, ast.SayNode{Message: tokens[i+1].Value})
					i += 2
				} else if i+1 < len(tokens) && tokens[i+1].Type == lexer.TokenTypeOperator && tokens[i+1].Value == "+" {
					// Handle string concatenation (very basic)
					if i+2 < len(tokens) && tokens[i+2].Type == lexer.TokenTypeIdentifier {
						nodes = append(nodes, ast.SayNode{Message: fmt.Sprintf("\"\" + %s", tokens[i+2].Value)}) // Placeholder
						i += 3
					} else {
						return nil, fmt.Errorf("syntax error in parayada statement at line %d", tokens[i].Line)
					}
				} else {
					return nil, fmt.Errorf("syntax error in parayada statement at line %d", tokens[i].Line)
				}
			default:
				fmt.Printf("Unhandled keyword: %s\n", tokens[i].Value)
				i++
			}
		case lexer.TokenTypeNewline:
			i++
		default:
			fmt.Printf("Unexpected token: %+v\n", tokens[i])
			i++
		}
	}
	return nodes, nil
}

func parseExpression(tokens []lexer.Token) (ast.Node, int, error) {
	if len(tokens) == 0 {
		return nil, 0, fmt.Errorf("unexpected end of expression")
	}

	switch tokens[0].Type {
	case lexer.TokenTypeString:
		return ast.StringLiteralNode{Value: tokens[0].Value}, 1, nil
	case lexer.TokenTypeNumber:
		return ast.NumberLiteralNode{Value: tokens[0].Value}, 1, nil
	case lexer.TokenTypeIdentifier:
		return ast.IdentifierNode{Name: tokens[0].Value}, 1, nil
	default:
		return nil, 0, fmt.Errorf("unexpected token in expression: %+v", tokens[0])
	}
}
