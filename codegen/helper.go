package codegen

import (
	"fmt"
	"regexp"
)

func operatorPrecedence(operator string) int {
	switch operator {
	case "*", "/": // Multiplication and division have highest precedence
		return 3
	case "+", "-": // Addition and subtraction
		return 2
	case "==", "<", ">", "<=", ">=", "!=": // Comparison - add missing operators
		return 1
	default:
		// Match behavior with isLeftAssociative
		fmt.Printf("Warning: Unknown operator in precedence: %s, defaulting to lowest precedence\n", operator)
		return 0
	}
}

func isLeftAssociative(operator string) bool {
	switch operator {
	case "+", "-", "*", "/", "==", "<", ">", "<=", ">=", "!=": // Add missing operators
		return true
	default:
		fmt.Printf("Warning: Unknown operator in associativity check: %s, assuming left associative\n", operator)
		return true
	}
}

func RemoveComments(input string) string {
	commentRegex := regexp.MustCompile(`//.*`)
	return commentRegex.ReplaceAllString(input, "")
}
