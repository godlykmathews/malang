package codegen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Rohith04MVK/malang/ast"
)

// GenerateCode generates Go code from the AST.
func GenerateCode(program ast.Program) string {
	var code strings.Builder

	// Add package main
	code.WriteString("package main\n\n")

	// Add import statements (conditionally)
	needsFmt := false
	needsStrconv := false

	// Check if fmt or strconv are needed
	for _, statement := range program.Statements {
		if usesFmt(statement) {
			needsFmt = true
		}
		if usesStrconv(statement) {
			needsStrconv = true
		}
	}

	if needsFmt || needsStrconv {
		code.WriteString("import (\n")
		if needsFmt {
			code.WriteString("\t\"fmt\"\n")
		}
		if needsStrconv {
			code.WriteString("\t\"strconv\"\n")
		}
		code.WriteString(")\n\n")
	}

	// Add func main() {
	code.WriteString("func main() {\n")

	declaredVars := make(map[string]string)

	for _, statement := range program.Statements {
		code.WriteString(generateStatementCode(statement, declaredVars))
	}

	// Close func main() }
	code.WriteString("}\n")

	return code.String()
}

func generateStatementCode(statement ast.ASTNode, declaredVars map[string]string) string {
	switch s := statement.(type) {
	case ast.ParayuStatement:
		return fmt.Sprintf("\tfmt.Println(%s)\n", generateExpressionCodeForParayu(s.Expression, declaredVars))
	case ast.KelkStatement:
		if _, declared := declaredVars[s.Identifier]; declared {
			panic(fmt.Sprintf("Variable '%s' already declared", s.Identifier))
		}
		declaredVars[s.Identifier] = "string"
		return fmt.Sprintf("\tvar %s string\n\tfmt.Scanln(&%s)\n", s.Identifier, s.Identifier)
	case ast.AssignmentStatement:
		if _, declared := declaredVars[s.Identifier]; !declared {
			inferredType := inferType(s.Expression, declaredVars)
			declaredVars[s.Identifier] = inferredType
			return fmt.Sprintf("\t%s := %s\n", s.Identifier, generateExpressionCode(s.Expression, 0, declaredVars))
		} else {
			return fmt.Sprintf("\t%s = %s\n", s.Identifier, generateExpressionCode(s.Expression, 0, declaredVars))
		}
	case ast.IfStatement:
		code := fmt.Sprintf("\tif %s {\n%s\t}", generateExpressionCode(s.Condition, 0, declaredVars), generateBlockCode(s.Body, declaredVars))
		if s.ElseBody != nil {
			code += fmt.Sprintf(" else {\n%s\t}\n", generateBlockCode(s.ElseBody, declaredVars))
		}
		return code
	case ast.WhileStatement:
		return fmt.Sprintf("\tfor %s {\n%s\t}\n", generateExpressionCode(s.Condition, 0, declaredVars), generateBlockCode(s.Body, declaredVars))
	case ast.ForStatement:
		loopDeclaredVars := make(map[string]string)
		for k, v := range declaredVars {
			loopDeclaredVars[k] = v
		}
		loopDeclaredVars[s.Identifier] = "int"
		startCode := generateExpressionCode(s.Start, 0, declaredVars)
		endCode := generateExpressionCode(s.End, 0, declaredVars)
		return fmt.Sprintf("\tfor %s := %s; %s <= %s; %s++ {\n%s\t}\n", s.Identifier, startCode, s.Identifier, endCode, s.Identifier, generateBlockCode(s.Body, loopDeclaredVars))
	default:
		panic(fmt.Sprintf("Unexpected statement type: %T", statement))
	}
}

func generateBlockCode(statements []ast.ASTNode, declaredVars map[string]string) string {
	var code strings.Builder
	for _, stmt := range statements {
		//add tab to each statement.
		code.WriteString(generateStatementCode(stmt, declaredVars))
	}
	return code.String()
}

func generateExpressionCode(expression ast.ASTNode, parentPrecedence int, declaredVars map[string]string) string {
	return generateExpressionCodeWithPrecedence(expression, parentPrecedence, false, declaredVars)
}

func generateExpressionCodeForParayu(expression ast.ASTNode, declaredVars map[string]string) string {
	return generateExpressionCodeWithPrecedence(expression, 0, true, declaredVars)
}

func generateExpressionCodeWithPrecedence(expression ast.ASTNode, parentPrecedence int, isParayu bool, declaredVars map[string]string) string {
	switch e := expression.(type) {
	case ast.StringLiteral:
		return fmt.Sprintf("%q", e.Value)
	case ast.IntegerLiteral:
		if isParayu {
			return fmt.Sprintf("strconv.Itoa(%d)", e.Value)
		}
		return strconv.Itoa(e.Value)
	case ast.Identifier:
		if isParayu {
			if declaredVars[e.Name] == "string" {
				return e.Name
			} else {
				return fmt.Sprintf("strconv.Itoa(%s)", e.Name)
			}
		}
		return e.Name
	case ast.BinaryExpression:
		precedence := operatorPrecedence(e.Operator)

		if e.Operator == "+" && isParayu {
			leftCode := generateExpressionCodeWithPrecedence(e.Left, precedence, true, declaredVars)
			rightCode := generateExpressionCodeWithPrecedence(e.Right, precedence, true, declaredVars)
			return fmt.Sprintf("%s + %s", leftCode, rightCode)

		} else { // Handle other operators (including -, *, /)
			leftCode := generateExpressionCodeWithPrecedence(e.Left, precedence, isParayu, declaredVars)
			rightCode := generateExpressionCodeWithPrecedence(e.Right, precedence, isParayu, declaredVars)

			// Add parentheses based on precedence and associativity.
			if precedence < parentPrecedence || (precedence == parentPrecedence && isLeftAssociative(e.Operator)) {
				return fmt.Sprintf("(%s %s %s)", leftCode, e.Operator, rightCode)
			}
			return fmt.Sprintf("%s %s %s", leftCode, e.Operator, rightCode)
		}
	default:
		panic(fmt.Sprintf("Unexpected expression type: %T", expression))
	}
}

func inferType(expression ast.ASTNode, declaredVars map[string]string) string {
	switch e := expression.(type) {
	case ast.StringLiteral:
		return "string"
	case ast.IntegerLiteral:
		return "int"
	case ast.Identifier:
		if t, ok := declaredVars[e.Name]; ok {
			return t
		} else {
			panic(fmt.Sprintf("Undeclared identifier %s during type inference", e.Name))
		}
	case ast.BinaryExpression: // Now handles more operators
		if e.Operator == "+" {
			leftType := inferType(e.Left, declaredVars)
			rightType := inferType(e.Right, declaredVars)
			if leftType == "string" || rightType == "string" {
				return "string" // String concatenation
			}
		}
		// For -, *, /, <, ==, assume int (more robust type checking is needed in a real compiler)
		return "int"

	default:
		panic(fmt.Sprintf("Cannot infer type for expression: %T", expression))
	}
}

func usesFmt(statement ast.ASTNode) bool {
	switch s := statement.(type) {
	case ast.ParayuStatement:
		return true
	case ast.KelkStatement:
		return true
	case ast.IfStatement:
		return usesFmt(s.Condition) || blockUsesFmt(s.Body) || (s.ElseBody != nil && blockUsesFmt(s.ElseBody))
	case ast.WhileStatement:
		return usesFmt(s.Condition) || blockUsesFmt(s.Body)
	case ast.ForStatement:
		return usesFmt(s.Start) || usesFmt(s.End) || blockUsesFmt(s.Body)
	case ast.BinaryExpression:
		return usesFmt(s.Left) || usesFmt(s.Right)
	default:
		return false
	}
}
func blockUsesFmt(statements []ast.ASTNode) bool {
	for _, stmt := range statements {
		if usesFmt(stmt) {
			return true
		}
	}
	return false
}
func usesStrconv(statement ast.ASTNode) bool {
	switch s := statement.(type) {
	case ast.ParayuStatement:
		return expressionUsesStrconv(s.Expression) // Check expression within parayu
	case ast.IfStatement:
		return usesStrconv(s.Condition) || blockUsesStrconv(s.Body) || (s.ElseBody != nil && blockUsesStrconv(s.ElseBody))
	case ast.WhileStatement:
		return usesStrconv(s.Condition) || blockUsesStrconv(s.Body)
	case ast.ForStatement:
		return expressionUsesStrconv(s.Start) || expressionUsesStrconv(s.End) || blockUsesStrconv(s.Body)
	case ast.AssignmentStatement:
		return expressionUsesStrconv(s.Expression) //check if right side of assignment needs it
	default:
		return false
	}
}
func blockUsesStrconv(statements []ast.ASTNode) bool {
	for _, stmt := range statements {
		if usesStrconv(stmt) {
			return true
		}
	}
	return false
}

func expressionUsesStrconv(expression ast.ASTNode) bool {
	switch e := expression.(type) {
	case ast.IntegerLiteral:
		return true // Integer literals within Parayu always need strconv
	case ast.Identifier:
		// Check if we have type information, default to false if unknown
		if e.Type == "" {
			return false
		}
		return e.Type == "int" // needs conversion if it is an int
	case ast.BinaryExpression:
		return expressionUsesStrconv(e.Left) || expressionUsesStrconv(e.Right)
	default:
		return false
	}
}
