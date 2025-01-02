package ast

import (
	"fmt"
)

func (n SayNode) String() string {
	return fmt.Sprintf("Say: %s", n.Message)
}

func (n VariableDeclarationNode) String() string {
	return fmt.Sprintf("VarDecl: %s = %v", n.Identifier, n.Value)
}

func (n VariableAssignmentNode) String() string {
	return fmt.Sprintf("VarAssign: %s = %v", n.Identifier, n.Value)
}

func (n IdentifierNode) String() string {
	return fmt.Sprintf("Identifier: %s", n.Name)
}

func (n StringLiteralNode) String() string {
	return fmt.Sprintf("StringLiteral: %q", n.Value)
}

func (n NumberLiteralNode) String() string {
	return fmt.Sprintf("NumberLiteral: %s", n.Value)
}
