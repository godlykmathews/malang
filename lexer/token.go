package lexer

type TokenType int

const (
	TokenTypeKeyword TokenType = iota
	TokenTypeIdentifier
	TokenTypeString
	TokenTypeEquals
	TokenTypeNumber
	TokenTypeOperator
	TokenTypeColon
	TokenTypeNewline
	TokenTypeEOF
)

type Token struct {
	Type   TokenType
	Value  string
	Line   int
	Column int
}
