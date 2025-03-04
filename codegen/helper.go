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
	case "==", "<": // Comparison
		return 1
	default:
		panic(fmt.Sprintf("Unknown operator: %s", operator))
	}
}
func isLeftAssociative(operator string) bool {
	switch operator {
	case "+", "-", "*", "/", "==", "<": //all our operators are left associative
		return true
	default:
		return false //or panic if an unknown operator is encountered
	}
}

func RemoveComments(input string) string {
	commentRegex := regexp.MustCompile(`//.*`)
	return commentRegex.ReplaceAllString(input, "")
}
