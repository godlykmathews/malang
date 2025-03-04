package lexer

type Token struct {
	Type  string
	Value string
	Line  int
	Col   int
}

// Constants for token types
const (
	TokParayu         = "PARAYU"
	TokKelk           = "KELK"
	TokAadhyamayi     = "AADHYAMAYI"
	TokAthengil       = "ATHENGIL"
	TokIlla           = "ILLA"
	TokEllamSheriyano = "ELLAM_SHERIYANO"
	TokOnninuMumbu    = "ONNINU_MUMBU"
	TokEdukk          = "EDUKK"
	TokString         = "STRING"
	TokIdentifier     = "IDENTIFIER"
	TokInteger        = "INTEGER"
	TokOperator       = "OPERATOR"
	TokLParen         = "LPAREN"
	TokRParen         = "RPAREN"
	TokLBrace         = "LBRACE"
	TokRBrace         = "RBRACE"
	TokRange          = "RANGE"
	TokComma          = "COMMA"
	TokEOF            = "EOF"
	TokMinus          = "MINUS"
	TokMultiply       = "MULTIPLY"
	TokDivide         = "DIVIDE"
)
