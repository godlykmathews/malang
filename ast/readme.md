#  Abstract Syntax Tree (AST)

The Abstract Syntax Tree (AST) is a tree-like representation of the parsed code, representing the *abstract* syntactic structure of the program.


* **Abstract vs. Concrete Syntax:** The AST is an *abstract* representation, meaning it omits details that are not essential to the program's meaning (like parentheses used solely for grouping, or specific keywords like `enkil` and `alle`). The *concrete* syntax is the exact sequence of characters in the source code. The parser transforms the concrete syntax into the abstract syntax.
* **Tree Structure:** The tree structure of the AST reflects the hierarchical relationships between different parts of the program. For example, an `IfStatement` node might have child nodes representing the condition, the "then" block, and the "else" block.



**Example (AST for x = 10 + 5 * 2):**

```json
AssignmentStatement{
    Identifier: "x",
    Expression: BinaryExpression{
        Left: IntegerLiteral{Value: 10},
        Operator: "+",
        Right: BinaryExpression{
            Left: IntegerLiteral{Value: 5},
            Operator: "*",
            Right: IntegerLiteral{Value: 2},
        },
    },
}
```