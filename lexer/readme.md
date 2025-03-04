# Lexer (Lexical Analysis)

The lexer, located in `malang/lexer/lexer.go`, performs the first phase of compilation: **lexical analysis**.  Its primary role is to convert the source code (a stream of characters) into a stream of **tokens**.

**Theoretical Background:**

*   **Regular Expressions:** Lexers are often built using **regular expressions**. While this lexer is hand-written, the patterns it matches (for identifiers, string literals, etc.) *could* be formally described using regular expressions.  Regular expressions provide a concise way to define the lexical structure of a language.
*   **Finite Automata:**  Regular expressions can be implemented using **finite automata** (specifically, Deterministic Finite Automata - DFAs).  A DFA is a mathematical model of computation that can efficiently recognize strings belonging to a regular language.  Essentially, the lexer acts as a DFA, transitioning between states based on the input characters.
*   **Tokenization:** The process of converting a character stream into a token stream is called **tokenization**. Each token represents a meaningful unit in the language.

**Key Components:**

*   **`Token` struct:**
```go
    type Token struct {
        Type  string
        Value string
        Line  int
        Col   int
    }
```

This struct represents a single token.  `Type` is the token type (e.g., `TokIdentifier`, `TokString`), `Value` is the actual text of the token (e.g., "name", `"Hello"`), and `Line` and `Col` store the token's position in the source code for error reporting.

*   **Token Types (Constants):**

    ```go
    const (
        TokParayu      = "PARAYU"
        TokKelk        = "KELK"
        // ... other token types ...
        TokEOF         = "EOF" // End-of-file marker
    )
    ```

These constants define all the possible token types in Malang.


*   **`Lex(input string) []Token` function:**

    This is the main function of the lexer. It takes the source code as a string (`input`) and returns a slice of `Token` structs.  It works by:

    1.  **Iterating through the input character by character.**
    2.  **Skipping whitespace and newlines.**
    3.  **Identifying different token types:**
        *   **Keywords:**  Uses a `map[string]string` to match keywords (e.g., "parayu", "kelk").
        *   **Identifiers:**  Matches sequences of letters, digits, and underscores.
        *   **String Literals:**  Matches text enclosed in double quotes.
        *   **Integer Literals:** Matches sequences of digits.
        *   **Operators:**  Matches operators like "+", "-", "*", "/", "==", "<", "=".
        *   **Parentheses and Braces:** Matches "(", ")", "{", "}".
        *   **Range Operator (".."):** Matches the two-dot sequence.
        * **Comments:** ignores comments.
    4.  **Creating `Token` structs for each identified token.**
    5.  **Appending the tokens to a slice.**
    6.  **Adding an `EOF` token at the end.**

    The lexer also handles basic error detection, such as unterminated string literals and unexpected characters.

**Example:**

Input:

```go
parayu("Hello" + name)
```

Output (Tokens):
```json
[
    {Type: "PARAYU", Value: "parayu", Line: 1, Col: 1},
    {Type: "LPAREN", Value: "(", Line: 1, Col: 7},
    {Type: "STRING", Value: "Hello", Line: 1, Col: 8},
    {Type: "OPERATOR", Value: "+", Line: 1, Col: 16},
    {Type: "IDENTIFIER", Value: "name", Line: 1, Col: 18},
    {Type: "RPAREN", Value: ")", Line: 1, Col: 22},
    {Type: "EOF", Value: "", Line: 1, Col: 23},
]
```