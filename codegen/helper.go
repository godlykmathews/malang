package codegen

import (
	"fmt"
	"regexp"
)

func operatorPrecedence(operator string) int {
	switch operator {
	case "==", "<": // Comparison operators
		return 1
	case "+": // Addition
		return 2
	// TODO: Add more operators and their precedence
	default:
		panic(fmt.Sprintf("Unknown operator: %s", operator))
	}
}

func isLeftAssociative(operator string) bool {
	switch operator {
	case "+", "==", "<": //all our operators are left associative
		return true
	default:
		return false //or panic if an unknown operator is encountered
	}
}

func RemoveComments(input string) string {
	commentRegex := regexp.MustCompile(`//.*`)
	return commentRegex.ReplaceAllString(input, "")
}
