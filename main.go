package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/Rohith04MVK/malang/codegen"
	"github.com/Rohith04MVK/malang/lexer"
	"github.com/Rohith04MVK/malang/parser"
)

func main() {
	// Define command-line flags
	debugTokens := flag.Bool("tokens", false, "Print tokens")
	debugAST := flag.Bool("ast", false, "Print AST")
	debugGoCode := flag.Bool("gocode", false, "Print generated Go code")
	flag.Parse()

	if flag.NArg() != 1 { //check if there is only one non flag argument.
		fmt.Println("Usage: mylang [options] <filename.mylang>")
		flag.PrintDefaults() //print all flags and their descriptions
		return
	}

	filename := flag.Arg(0) // Use flag.Arg to get positional arguments
	inputBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	input := string(inputBytes)
	input = codegen.RemoveComments(input)

	tokens := lexer.Lex(input)
	if *debugTokens {
		fmt.Println("Tokens:", tokens)
	}

	p := parser.NewParser(tokens)
	program := p.Parse()
	if *debugAST {
		fmt.Println("AST:", program)
	}

	generatedCode := codegen.GenerateCode(program)
	if *debugGoCode {
		fmt.Println("Generated Go Code:\n", generatedCode)
	}

	// --- Execute the generated code directly ---

	// 1. Create a temporary file.
	tmpFile, err := os.CreateTemp("", "mylang*.go")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	//defer os.Remove(tmpFile.Name()) // Clean up the temporary file

	// 2. Write the generated Go code to the temporary file.
	if _, err := tmpFile.WriteString(generatedCode); err != nil {
		fmt.Println("Error writing to temporary file:", err)
		return
	}
	if err := tmpFile.Close(); err != nil {
		fmt.Println("Error closing temporary file:", err)
		return
	}
	// 3. Use 'go run' to execute the code in the temporary file.
	cmd := exec.Command("go", "run", tmpFile.Name())
	cmd.Stdin = os.Stdin   // Connect standard input
	cmd.Stdout = os.Stdout // Connect standard output
	cmd.Stderr = os.Stderr // Connect standard error

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running generated code:", err)
		// Delete the temp file *after* we attempt to run it.  This is
		// crucial for debugging.  If the go run command fails, we *want*
		// the temp file to still exist so we can inspect it.
		fmt.Println("Temp file:", tmpFile.Name())
	} else {
		os.Remove(tmpFile.Name()) //only delete if runs.
	}
}
