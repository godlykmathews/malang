package codegen

import (
	"fmt"
	"strings"

	"github.com/Rohith04MVK/malang/ast"
)

func GenerateCode(nodes []ast.Node) string {
	var goCode strings.Builder
	goCode.WriteString("package main\n\nimport \"fmt\"\n\n")
	generateExpressionCode := func(node ast.Node) string {
		switch expr := node.(type) {
		case ast.StringLiteralNode:
			return fmt.Sprintf("%q", expr.Value) // Use %q for quoted strings
		case ast.NumberLiteralNode:
			return expr.Value
		case ast.IdentifierNode:
			return expr.Name
		default:
			return fmt.Sprintf("/* Unsupported expression: %T */", node)
		}
	}

	// Helper function to generate code for a block of statements
	var generateBlock func(block []ast.Node, indent string)
	generateBlock = func(block []ast.Node, indent string) {
		for _, node := range block {
			switch n := node.(type) {
			case ast.SayNode:
				goCode.WriteString(fmt.Sprintf("%sfmt.Println(`%s`)\n", indent, n.Message))
			case ast.VariableDeclarationNode:
				// Basic declaration - inferring type as string for simplicity
				goCode.WriteString(fmt.Sprintf("%svar %s %s = %s\n", indent, n.Identifier, ,generateExpressionCode(n.Value)))
			case ast.VariableAssignmentNode:
				goCode.WriteString(fmt.Sprintf("%s%s = %s\n", indent, n.Identifier, generateExpressionCode(n.Value)))
			default:
				fmt.Printf("Unknown node type for code generation: %T\n", n)
			}
		}
	}

	goCode.WriteString("func main() {\n")
	generateBlock(nodes, "\t")
	goCode.WriteString("}\n")

	return goCode.String()
}
