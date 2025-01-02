package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Rohith04MVK/malang/codegen"
	"github.com/Rohith04MVK/malang/lexer"
	"github.com/Rohith04MVK/malang/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run your_compiler.go <input_file.lang>")
		return
	}

	filename := os.Args[1]
	sourceCodeBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	sourceCode := string(sourceCodeBytes)

	tokens, err := lexer.Lex(sourceCode)
	debug := false
	if len(os.Args) == 3 && os.Args[2] == "debug" {
		debug = true
	}

	if err != nil {
		fmt.Println("Lexing Error:", err)
		return
	}
	if debug {
		fmt.Println("Tokens:")
		for _, token := range tokens {
			fmt.Printf("%+v\n", token)
		}
	}

	ast, err := parser.Parse(tokens)
	if err != nil {
		fmt.Println("Parsing Error:", err)
		return
	}
	if debug {
		fmt.Println("\nAST:")
		for _, node := range ast {
			fmt.Printf("%s\n", node.String())
		}
	}

	goCode := codegen.GenerateCode(ast)
	if debug {
		fmt.Println("\nGenerated Go Code:")
		fmt.Println(goCode)
	}

	// Write generated Go code to a temporary file
	tempGoFile := "temp.go"
	err = os.WriteFile(tempGoFile, []byte(goCode), 0644)
	if err != nil {
		fmt.Printf("Error writing temporary Go file: %v\n", err)
		return
	}
	defer os.Remove(tempGoFile) // Clean up the temporary file

	// Compile the generated Go code
	cmd := exec.Command("go", "run", tempGoFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Compilation or Execution Error:\n%s\n", output)
		return
	}

	// fmt.Printf("Compilation and Execution Time: %s\n", time.Since(startTime))
	fmt.Println(string(output))
}
