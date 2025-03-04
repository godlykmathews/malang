# Parser (Syntax Analysis)

The parser, located in `internal/parser/parser.go`, performs **syntax analysis**.  It takes the token stream from the lexer and verifies that it conforms to the **grammar** of the Malang language. The output of the parser is an **Abstract Syntax Tree (AST)**.

**Theoretical Background:**

*   **Context-Free Grammars (CFGs):** Programming language syntax is typically defined using **context-free grammars (CFGs)**.  A CFG consists of:
    *   **Terminals:**  The basic symbols of the language (represented by tokens in the lexer).
    *   **Non-terminals:**  Symbols that represent syntactic categories (e.g., "statement", "expression").
    *   **Production Rules:**  Rules that specify how non-terminals can be expanded into sequences of terminals and non-terminals (e.g., `statement -> if_statement | while_statement | assignment_statement`).
    *   **Start Symbol:**  A designated non-terminal that represents the entire program.

*   **Recursive Descent Parsing:**  Malang uses a **recursive descent parser**, which is a common type of **top-down parser**.
    *   **Top-Down:**  It starts with the start symbol of the grammar (e.g., `Program`) and tries to derive the input string by applying production rules.
    *   **Recursive:**  Each non-terminal in the grammar is typically represented by a separate parsing function. These functions call each other recursively, mirroring the structure of the grammar.
    *   **Predictive Parsing:** Recursive descent parsers are often **predictive**, meaning they can determine which production rule to apply based on the current input token (the "lookahead" token).  This requires the grammar to be LL(1) - see below.

*   **LL(1) Grammars:** Ideally, a recursive descent parser works best with an **LL(1) grammar**.  This means:
    *   **L**eft-to-right scan of the input.
    *   **L**eftmost derivation (the leftmost non-terminal is always expanded first).
    *   **1** symbol of lookahead (the parser can decide which production rule to apply by looking at just the next token).

    While this compiler doesn't explicitly check for LL(1) properties, the grammar of Malang is designed to be simple enough to be parsed with a single token of lookahead. If you were to add more complex features, you might need to consider grammar transformations (like left-factoring or eliminating left recursion) to ensure it remains LL(1).

* **Operator Precedence and Associativity:** The parser correctly handles **operator precedence** (e.g., multiplication before addition) and **associativity** (e.g., left-associativity for subtraction: `a - b - c` is interpreted as `(a - b) - c`). This is achieved through the structure of the parsing functions (`parseComparison`, `parseTerm`, `parseFactor`) and the recursive calls between them.

**Key Components:**

**Key Components:**

*   **`Parser` struct:**

    ```go
    type Parser struct {
        tokens []lexer.Token
        pos    int
    }
    ```

    This struct holds the token stream and the current position within the stream.

*   **`NewParser(tokens []lexer.Token) *Parser`:** This constructor function creates a new `Parser` instance.

*   **`peek()`, `consume()`, `peekNext()` methods:** These helper methods are used to look at and consume tokens from the stream.

*   **Parsing Functions:**
    *   **`parse()`:** The top-level parsing function.
    *   **`parseStatement()`:** Parses a single statement.
    *   **`parseExpression()`:** Parses expressions, and importantly, correctly handles **operator precedence** and **associativity**. It achieves this through recursive calls to:
        *   **`parseComparison()`:** Handles comparison operators (`==`, `<`).
        *   **`parseTerm()`:** Handles addition and subtraction (`+`, `-`).
        *   **`parseFactor()`:** Handles multiplication and division (`*`, `/`).
        *    **`parsePrimary`:** Handles atomic expressions.

    *   **`parseBlock()`:** Parses a block of code enclosed in curly braces.

**Example (Parsing `parayu` statement):**

```go
func (p *Parser) parseParayuStatement() ast.ASTNode {
    p.consume(lexer.TokParayu) // Consumes "parayu"
    p.consume(lexer.TokLParen) // Consumes "("
    expression := p.parseExpression() // Parses the expression inside the parentheses
    p.consume(lexer.TokRParen) // Consumes ")"
    return ast.ParayuStatement{Expression: expression} // Creates an AST node
}
```

The parser enforces the syntax rules of Malang. For instance, it ensures that:

- `parayu` statements are followed by parentheses and an expression.
- `if` statements have a condition, `enkil`, a block, and optionally `alle` and another block.