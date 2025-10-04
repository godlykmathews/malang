package parser

import (
	"fmt"
	"strconv"

	"github.com/Rohith04MVK/malang/ast"
	"github.com/Rohith04MVK/malang/lexer"
)

type Parser struct {
	tokens []lexer.Token
	pos    int
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{tokens: tokens, pos: 0}
}

func (p *Parser) peek() lexer.Token {
	if p.pos >= len(p.tokens) {
		return lexer.Token{Type: lexer.TokEOF}
	}
	return p.tokens[p.pos]
}

func (p *Parser) consume(expectedType string) lexer.Token {
	token := p.peek()
	if token.Type != expectedType {
		panic(fmt.Sprintf("Expected token type %s, got %s at line %d, col %d", expectedType, token.Type, token.Line, token.Col))
	}
	p.pos++
	return token
}

func (p *Parser) Parse() ast.Program {
	program := ast.Program{Statements: []ast.ASTNode{}}
	for p.peek().Type != lexer.TokEOF {
		program.Statements = append(program.Statements, p.parseStatement())
	}
	return program
}

func (p *Parser) parseStatement() ast.ASTNode {
	switch p.peek().Type {
	case lexer.TokParayu:
		return p.parseParayuStatement()
	case lexer.TokKelk:
		return p.parseKelkStatement()
	case lexer.TokIdentifier: // Could be assignment or part of expression
		if p.peekNext().Type == lexer.TokOperator && p.peekNext().Value == "=" {
			return p.parseAssignmentStatement()
		}
		fallthrough // otherwise parse as an expression statement (for now, just expressions)
	case lexer.TokAadhyamayi:
		return p.parseIfStatement()
	case lexer.TokEllamSheriyano:
		return p.parseWhileStatement()
	case lexer.TokOnninuMumbu:
		return p.parseForStatement()
	default:
		return p.parseExpression()
	}
}

func (p *Parser) parseParayuStatement() ast.ASTNode {
	p.consume(lexer.TokParayu)
	p.consume(lexer.TokLParen)
	expression := p.parseExpression()
	p.consume(lexer.TokRParen)
	return ast.ParayuStatement{Expression: expression}
}

func (p *Parser) parseKelkStatement() ast.ASTNode {
	p.consume(lexer.TokKelk)
	p.consume(lexer.TokLParen)
	identifier := p.consume(lexer.TokIdentifier).Value
	p.consume(lexer.TokRParen)

	return ast.KelkStatement{Identifier: identifier}
}

func (p *Parser) parseAssignmentStatement() ast.ASTNode {
	identifier := p.consume(lexer.TokIdentifier).Value
	p.consume(lexer.TokOperator) // We already know it's an '='
	expression := p.parseExpression()
	return ast.AssignmentStatement{Identifier: identifier, Expression: expression}
}

func (p *Parser) parseIfStatement() ast.ASTNode {
	p.consume(lexer.TokAadhyamayi)
	p.consume(lexer.TokLParen)
	condition := p.parseExpression()
	p.consume(lexer.TokRParen)
	p.consume(lexer.TokAthengil)
	p.consume(lexer.TokLBrace)
	body := p.parseBlock()
	p.consume(lexer.TokRBrace)

	var elseBody []ast.ASTNode
	if p.peek().Type == lexer.TokIlla {
		p.consume(lexer.TokIlla)
		p.consume(lexer.TokLBrace)
		elseBody = p.parseBlock()
		p.consume(lexer.TokRBrace)
	}

	return ast.IfStatement{Condition: condition, Body: body, ElseBody: elseBody}
}

func (p *Parser) parseWhileStatement() ast.ASTNode {
	p.consume(lexer.TokEllamSheriyano)
	p.consume(lexer.TokLParen)
	condition := p.parseExpression()
	p.consume(lexer.TokRParen)
	p.consume(lexer.TokAthengil)
	p.consume(lexer.TokLBrace)
	body := p.parseBlock()
	p.consume(lexer.TokRBrace)

	return ast.WhileStatement{Condition: condition, Body: body}
}

func (p *Parser) parseForStatement() ast.ASTNode {
	p.consume(lexer.TokOnninuMumbu)
	identifier := p.consume(lexer.TokIdentifier).Value
	p.consume(lexer.TokEdukk)
	p.consume(lexer.TokLParen)
	start := p.parseExpression()
	p.consume(lexer.TokRange)
	end := p.parseExpression()
	p.consume(lexer.TokRParen)

	p.consume(lexer.TokLBrace)
	body := p.parseBlock()
	p.consume(lexer.TokRBrace)
	return ast.ForStatement{Identifier: identifier, Start: start, End: end, Body: body}
}

func (p *Parser) parseExpression() ast.ASTNode {
	return p.parseComparison()
}

func (p *Parser) parseComparison() ast.ASTNode {
	left := p.parseTerm() // Use parseTerm here
	for p.peek().Type == lexer.TokOperator &&
		(p.peek().Value == "==" || p.peek().Value == "<" ||
			p.peek().Value == ">" || p.peek().Value == "<=" ||
			p.peek().Value == ">=" || p.peek().Value == "!=") {
		operator := p.consume(lexer.TokOperator).Value
		right := p.parseTerm() // And here
		left = ast.BinaryExpression{Left: left, Operator: operator, Right: right}
	}
	return left
}
func (p *Parser) parseTerm() ast.ASTNode {
	left := p.parseFactor()

	for p.peek().Type == lexer.TokOperator && (p.peek().Value == "+" || p.peek().Value == "-") {
		operator := p.consume(lexer.TokOperator).Value
		right := p.parseFactor()
		left = ast.BinaryExpression{Left: left, Operator: operator, Right: right}
	}
	return left
}

func (p *Parser) parseFactor() ast.ASTNode {
	left := p.parsePrimary()
	for p.peek().Type == lexer.TokOperator && (p.peek().Value == "*" || p.peek().Value == "/") {
		operator := p.consume(lexer.TokOperator).Value
		right := p.parsePrimary()
		left = ast.BinaryExpression{Left: left, Operator: operator, Right: right}
	}
	return left
}

func (p *Parser) parsePrimary() ast.ASTNode {
	switch p.peek().Type {
	case lexer.TokInteger:
		value, _ := strconv.Atoi(p.consume(lexer.TokInteger).Value)
		return ast.IntegerLiteral{Value: value}
	case lexer.TokString:
		return ast.StringLiteral{Value: p.consume(lexer.TokString).Value}
	case lexer.TokIdentifier:
		name := p.consume(lexer.TokIdentifier).Value
		return ast.Identifier{Name: name, Type: ""}
	case lexer.TokLParen:
		p.consume(lexer.TokLParen)
		expression := p.parseExpression()
		p.consume(lexer.TokRParen)
		return expression
	default:
		panic(fmt.Sprintf("Unexpected token in expression: %s", p.peek().Type))
	}
}

func (p *Parser) peekNext() lexer.Token {
	if p.pos+1 >= len(p.tokens) {
		return lexer.Token{Type: lexer.TokEOF}
	}
	return p.tokens[p.pos+1]
}
