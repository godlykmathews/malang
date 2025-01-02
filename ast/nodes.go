package ast

// Node is the interface that all AST nodes must implement.
type Node interface {
	// String returns a string representation of the node.
	String() string
}

// SayNode represents a say statement.
type SayNode struct {
	// Say node is like the print statement in Python.
	Message string // The message to say. E.g. "Hello, world!"
	// MessageType string // The type of message to say. E.g. "string" // TODO: Implement this.
}

type VariableDeclarationNode struct {
	Identifier string
	Value      Node // Expression for the initial value
}

type VariableAssignmentNode struct {
	Identifier string
	Value      Node
}

type IdentifierNode struct { // To represent a variable being used
	Name string
}

// Expression Node Interface
type ExpressionNode interface {
	Node
	isExpression()
}

// Concrete Expression Node Types
type StringLiteralNode struct {
	Value string
}

func (StringLiteralNode) isExpression() {}

type NumberLiteralNode struct {
	Value string
}

func (NumberLiteralNode) isExpression() {}

func (IdentifierNode) isExpression() {} // Variables can be expressions

type BinaryOpNode struct {
	Left     Node
	Operator string
	Right    Node
}

func (BinaryOpNode) isExpression() {}
