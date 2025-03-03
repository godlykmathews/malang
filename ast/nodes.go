package ast

type ASTNode interface{}

type Program struct {
	Statements []ASTNode
}

type ParayuStatement struct {
	Expression ASTNode
}

type KelkStatement struct {
	Identifier string // Variable to store input
}

type AssignmentStatement struct {
	Identifier string
	Expression ASTNode
}

type IfStatement struct {
	Condition ASTNode
	Body      []ASTNode
	ElseBody  []ASTNode // Optional else block
}

type WhileStatement struct {
	Condition ASTNode
	Body      []ASTNode
}

type ForStatement struct {
	Identifier string
	Start      ASTNode // Start of range
	End        ASTNode // End of range
	Body       []ASTNode
}

type BinaryExpression struct {
	Left     ASTNode
	Operator string
	Right    ASTNode
}

type StringLiteral struct {
	Value string
}

type Identifier struct {
	Name string
	Type string // Store the inferred type: "string" or "int" (or other types later)
}

type IntegerLiteral struct {
	Value int
}
