package parser

import (
	"github.com/Rohith04MVK/malang/ast"
	"github.com/Rohith04MVK/malang/lexer"
)

func (p *Parser) parseBlock() []ast.ASTNode {
	statements := []ast.ASTNode{}
	for p.peek().Type != lexer.TokRBrace && p.peek().Type != lexer.TokEOF {
		statements = append(statements, p.parseStatement())
	}
	return statements
}
