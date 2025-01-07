package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Token Types (for lexical analysis)
type TokenType int

const (
	// Keywords
	LetKeyword TokenType = iota
	SayKeyword
	IfKeyword
	OtherwiseKeyword
	RepeatKeyword
	TimesKeyword
	TakeKeyword
	InputKeyword
	AsKeyword

	// Operators
	PlusOperator
	MultiplyOperator
	GreaterThanOperator
	EqualsOperator

	// Literals
	NumberLiteral
	StringLiteral

	// Identifiers
	Identifier

	// Other
	Colon
	Newline

	EOF // End of File
)

type Token struct {
	Type   TokenType
	Value  string
	Line   int
	Column int
}

// Abstract Syntax Tree (AST) Nodes
type ASTNode interface {
	String() string // For debugging
}

type LetStatement struct {
	Identifier string
	Expression ASTNode
}

func (ls *LetStatement) String() string {
	return "LetStatement"
}

type SayStatement struct {
	Expression ASTNode
}

func (ss *SayStatement) String() string {
	return "SayStatement"
}

type IfStatement struct {
	Condition ASTNode
	ThenBlock []ASTNode
	ElseBlock []ASTNode
}

func (is *IfStatement) String() string {
	return "IfStatement"
}

type RepeatStatement struct {
	Count int
	Body  []ASTNode
}

func (rs *RepeatStatement) String() string {
	return "RepeatStatement"
}

type TakeInputStatement struct {
	Identifier string
}

func (tis *TakeInputStatement) String() string {
	return "TakeInputStatement"
}

// Expressions
type BinaryExpression struct {
	Left     ASTNode
	Operator TokenType
	Right    ASTNode
}

func (be *BinaryExpression) String() string {
	return "BinaryExpression"
}

type IdentifierExpression struct {
	Name string
}

func (ie *IdentifierExpression) String() string {
	return "IdentifierExpression"
}

type NumberLiteralExpression struct {
	Value int
}

func (nle *NumberLiteralExpression) String() string {
	return "NumberLiteralExpression"
}

type StringLiteralExpression struct {
	Value string
}

func (sle *StringLiteralExpression) String() string {
	return "StringLiteralExpression"
}

// Scanner (Lexical Analysis)
type Scanner struct {
	input   string
	start   int
	current int
	line    int
	column  int
	tokens  []Token
}

func NewScanner(input string) *Scanner {
	return &Scanner{input: input, line: 1, column: 1}
}

func (s *Scanner) ScanTokens() ([]Token, error) {
	for !s.isAtEnd() {
		s.start = s.current
		if err := s.scanToken(); err != nil {
			return nil, err
		}
	}
	s.tokens = append(s.tokens, Token{Type: EOF, Line: s.line, Column: s.column})
	return s.tokens, nil
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.input)
}

func (s *Scanner) advance() rune {
	s.current++
	s.column++
	return rune(s.input[s.current-1])
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return 0
	}
	return rune(s.input[s.current])
}

func (s *Scanner) addToken(tokenType TokenType) {
	s.addTokenWithValue(tokenType, s.input[s.start:s.current])
}

func (s *Scanner) addTokenWithValue(tokenType TokenType, value string) {
	s.tokens = append(s.tokens, Token{Type: tokenType, Value: value, Line: s.line, Column: s.start + 1})
}

func (s *Scanner) scanToken() error {
	c := s.advance()
	switch c {
	case ' ', '\r', '\t':
		// Ignore whitespace
	case '\n':
		s.addToken(Newline)
		s.line++
		s.column = 1
	case '+':
		s.addToken(PlusOperator)
	case '=':
		s.addToken(EqualsOperator)
	case '*':
		s.addToken(MultiplyOperator)
	case '>':
		s.addToken(GreaterThanOperator)
	case ':':
		s.addToken(Colon)
	case '"':
		return s.string()
	default:
		if isDigit(c) {
			return s.number()
		} else if isAlpha(c) {
			return s.identifier()
		}
		return fmt.Errorf("unexpected character '%c' at line %d, column %d", c, s.line, s.column-1)
	}
	return nil
}

func (s *Scanner) string() error {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
			s.column = 1
		}
		s.advance()
	}

	if s.isAtEnd() {
		return fmt.Errorf("unterminated string at line %d, column %d", s.line, s.start+1)
	}

	s.advance() // The closing "
	value := s.input[s.start+1 : s.current-1]
	s.addTokenWithValue(StringLiteral, value)
	return nil
}

func (s *Scanner) number() error {
	for isDigit(s.peek()) {
		s.advance()
	}
	s.addToken(NumberLiteral)
	return nil
}

func (s *Scanner) identifier() error {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := s.input[s.start:s.current]
	switch text {
	case "let":
		s.addToken(LetKeyword)
	case "say":
		s.addToken(SayKeyword)
	case "if":
		s.addToken(IfKeyword)
	case "otherwise":
		s.addToken(OtherwiseKeyword)
	case "repeat":
		s.addToken(RepeatKeyword)
	case "times":
		s.addToken(TimesKeyword)
	case "take":
		s.addToken(TakeKeyword)
	case "input":
		s.addToken(InputKeyword)
	case "as":
		s.addToken(AsKeyword)
	default:
		s.addToken(Identifier)
	}
	return nil
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isAlphaNumeric(c rune) bool {
	return isAlpha(c) || isDigit(c)
}

// Parser (Syntax Analysis)
type Parser struct {
	tokens  []Token
	current int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) Parse() ([]ASTNode, error) {
	statements := []ASTNode{}
	for !p.isAtEnd() {
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		statements = append(statements, stmt)
	}
	return statements, nil
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == EOF
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}

func (p *Parser) check(tokenType TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *Parser) match(tokenTypes ...TokenType) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) consume(tokenType TokenType, message string) (Token, error) {
	if p.check(tokenType) {
		return p.advance(), nil
	}
	return Token{}, fmt.Errorf("at line %d, column %d: %s", p.peek().Line, p.peek().Column, message)
}

func (p *Parser) parseStatement() (ASTNode, error) {
	if p.match(LetKeyword) {
		return p.parseLetStatement()
	}
	if p.match(SayKeyword) {
		return p.parseSayStatement()
	}
	if p.match(IfKeyword) {
		return p.parseIfStatement()
	}
	if p.match(RepeatKeyword) {
		return p.parseRepeatStatement()
	}
	if p.match(TakeKeyword) {
		return p.parseTakeInputStatement()
	}
	return nil, fmt.Errorf("at line %d, column %d: unexpected token '%s'", p.peek().Line, p.peek().Column, p.peek().Value)
}

func (p *Parser) parseLetStatement() (ASTNode, error) {
	identifierToken, err := p.consume(Identifier, "expected variable name after 'let'")
	if err != nil {
		return nil, err
	}
	_, err = p.consume(EqualsOperator, "expected '=' after variable name")
	if err != nil {
		return nil, err
	}
	expression, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(Newline, "expected newline after let statement")
	if err != nil {
		return nil, err
	}
	return &LetStatement{Identifier: identifierToken.Value, Expression: expression}, nil
}

func (p *Parser) parseSayStatement() (ASTNode, error) {
	expression, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(Newline, "expected newline after say statement")
	if err != nil {
		return nil, err
	}
	return &SayStatement{Expression: expression}, nil
}

func (p *Parser) parseIfStatement() (ASTNode, error) {
	condition, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(Colon, "expected ':' after if condition")
	if err != nil {
		return nil, err
	}

	thenBlock, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	var elseBlock []ASTNode
	if p.match(OtherwiseKeyword) {
		elseBlock, err = p.parseBlock()
		if err != nil {
			return nil, err
		}
	}
	return &IfStatement{Condition: condition, ThenBlock: thenBlock, ElseBlock: elseBlock}, nil
}

func (p *Parser) parseRepeatStatement() (ASTNode, error) {
	countToken, err := p.consume(NumberLiteral, "expected number after 'repeat'")
	if err != nil {
		return nil, err
	}
	count, err := strconv.Atoi(countToken.Value)
	if err != nil {
		return nil, fmt.Errorf("at line %d, column %d: invalid number '%s'", countToken.Line, countToken.Column, countToken.Value)
	}
	_, err = p.consume(TimesKeyword, "expected 'times' after repeat count")
	if err != nil {
		return nil, err
	}
	_, err = p.consume(Colon, "expected ':' after 'times'")
	if err != nil {
		return nil, err
	}
	body, err := p.parseBlock()
	if err != nil {
		return nil, err
	}
	return &RepeatStatement{Count: count, Body: body}, nil
}

func (p *Parser) parseTakeInputStatement() (ASTNode, error) {
	_, err := p.consume(InputKeyword, "expected 'input' after 'take'")
	if err != nil {
		return nil, err
	}
	_, err = p.consume(AsKeyword, "expected 'as' after 'input'")
	if err != nil {
		return nil, err
	}
	identifierToken, err := p.consume(Identifier, "expected variable name after 'as'")
	if err != nil {
		return nil, err
	}
	_, err = p.consume(Newline, "expected newline after take input statement")
	if err != nil {
		return nil, err
	}
	return &TakeInputStatement{Identifier: identifierToken.Value}, nil
}

func (p *Parser) parseBlock() ([]ASTNode, error) {
	statements := []ASTNode{}
	// Keep parsing statements until we encounter a keyword that signals the end of the block
	// or the end of the file. This is a simplified way to handle blocks without explicit
	// delimiters like braces.
	for !p.isAtEnd() && !p.check(OtherwiseKeyword) && !p.check(Newline) { // Adjust conditions as needed
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		statements = append(statements, stmt)
	}
	return statements, nil
}

func (p *Parser) parseExpression() (ASTNode, error) {
	return p.parseTerm() // For now, only terms
}

func (p *Parser) parseTerm() (ASTNode, error) {
	left, err := p.parseFactor()
	if err != nil {
		return nil, err
	}

	for p.match(PlusOperator, MultiplyOperator, GreaterThanOperator, EqualsOperator) {
		operator := p.previous().Type
		right, err := p.parseFactor()
		if err != nil {
			return nil, err
		}
		left = &BinaryExpression{Left: left, Operator: operator, Right: right}
	}
	return left, nil
}

func (p *Parser) parseFactor() (ASTNode, error) {
	if p.match(NumberLiteral) {
		num, err := strconv.Atoi(p.previous().Value)
		if err != nil {
			return nil, fmt.Errorf("at line %d, column %d: invalid number '%s'", p.previous().Line, p.previous().Column, p.previous().Value)
		}
		return &NumberLiteralExpression{Value: num}, nil
	}
	if p.match(StringLiteral) {
		return &StringLiteralExpression{Value: p.previous().Value}, nil
	}
	if p.match(Identifier) {
		return &IdentifierExpression{Name: p.previous().Value}, nil
	}
	return nil, fmt.Errorf("at line %d, column %d: expected expression, but found '%s'", p.peek().Line, p.peek().Column, p.peek().Value)
}

// Semantic Analysis and Code Generation
type Compiler struct {
	environment map[string]interface{} // Simple environment for variables
	output      strings.Builder
	reader      *bufio.Reader
}

func NewCompiler() *Compiler {
	return &Compiler{
		environment: make(map[string]interface{}),
		reader:      bufio.NewReader(os.Stdin),
	}
}

func (c *Compiler) Compile(sourceCode string) (string, error) {
	scanner := NewScanner(sourceCode)
	tokens, err := scanner.ScanTokens()
	if err != nil {
		return "", err
	}

	parser := NewParser(tokens)
	ast, err := parser.Parse()
	if err != nil {
		return "", err
	}

	c.output.WriteString("package main\n\nimport \"fmt\"\n\"strings\"\n\"bufio\"\n\"os\"\"\nfunc main() {\n")
	for _, node := range ast {
		if err := c.generateGoCode(node); err != nil {
			return "", err
		}
	}
	c.output.WriteString("}\n")
	return c.output.String(), nil
}

func (c *Compiler) generateGoCode(node ASTNode) error {
	switch stmt := node.(type) {
	case *LetStatement:
		c.environment[stmt.Identifier] = nil // Placeholder, type not enforced strictly
		c.output.WriteString(fmt.Sprintf("\t%s := ", stmt.Identifier))
		if err := c.generateGoCode(stmt.Expression); err != nil {
			return err
		}
		c.output.WriteString("\n")
	case *SayStatement:
		c.output.WriteString("\tfmt.Println(")
		if err := c.generateGoCode(stmt.Expression); err != nil {
			return err
		}
		c.output.WriteString(")\n")
	case *IfStatement:
		c.output.WriteString("\tif ")
		if err := c.generateGoCode(stmt.Condition); err != nil {
			return err
		}
		c.output.WriteString(" {\n")
		for _, thenStmt := range stmt.ThenBlock {
			if err := c.generateGoCode(thenStmt); err != nil {
				return err
			}
		}
		c.output.WriteString("\t}")
		if len(stmt.ElseBlock) > 0 {
			c.output.WriteString(" else {\n")
			for _, elseStmt := range stmt.ElseBlock {
				if err := c.generateGoCode(elseStmt); err != nil {
					return err
				}
			}
			c.output.WriteString("\t}\n")
		}
	case *RepeatStatement:
		c.output.WriteString(fmt.Sprintf("\tfor i := 0; i < %d; i++ {\n", stmt.Count))
		for _, bodyStmt := range stmt.Body {
			if err := c.generateGoCode(bodyStmt); err != nil {
				return err
			}
		}
		c.output.WriteString("\t}\n")
	case *TakeInputStatement:
		c.output.WriteString(fmt.Sprintf("\tfmt.Print(\"Enter input for %s: \")\n", stmt.Identifier))
		c.output.WriteString(fmt.Sprintf("\t%sStr, _ := c.reader.ReadString('\\n')\n", stmt.Identifier))
		c.output.WriteString(fmt.Sprintf("\t%s = strings.TrimSpace(%sStr)\n", stmt.Identifier, stmt.Identifier))
		c.environment[stmt.Identifier] = "" // Store as string initially
	case *BinaryExpression:
		if err := c.generateGoCode(stmt.Left); err != nil {
			return err
		}
		c.output.WriteString(" ")
		c.output.WriteString(c.translateOperator(stmt.Operator))
		c.output.WriteString(" ")
		if err := c.generateGoCode(stmt.Right); err != nil {
			return err
		}
	case *IdentifierExpression:
		c.output.WriteString(stmt.Name)
	case *NumberLiteralExpression:
		c.output.WriteString(strconv.Itoa(stmt.Value))
	case *StringLiteralExpression:
		c.output.WriteString(fmt.Sprintf("\"%s\"", stmt.Value))
	default:
		return fmt.Errorf("unknown AST node type: %T", node)
	}
	return nil
}

func (c *Compiler) translateOperator(op TokenType) string {
	switch op {
	case PlusOperator:
		return "+"
	case MultiplyOperator:
		return "*"
	case GreaterThanOperator:
		return ">"
	case EqualsOperator:
		return "=="
	default:
		return "" // Should not happen
	}
}

func main() {
	sourceCode := `take input as user_value
say "You entered: " + user_value
if user_value > "exit": say "Goodbye!"
otherwise: say "You didn't type 'exit'. Program continues."
`

	compiler := NewCompiler()
	goCode, err := compiler.Compile(sourceCode)
	if err != nil {
		fmt.Println("Compilation Error:", err)
		return
	}

	fmt.Println("Generated Go Code:\n", goCode)

}
