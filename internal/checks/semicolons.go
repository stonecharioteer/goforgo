package checks

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"github.com/stonecharioteer/goforgo/internal/analysis"
)

func init() {
	analysis.Register(&noUnnecessarySemicolonCheck{})
}

type noUnnecessarySemicolonCheck struct{}

func (c *noUnnecessarySemicolonCheck) Name() string {
	return "no_unnecessary_semicolon"
}

func (c *noUnnecessarySemicolonCheck) Description() string {
	return "Checks if the file contains unnecessary semicolons."
}

func (c *noUnnecessarySemicolonCheck) Execute(filePath string) (bool, string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return false, "", fmt.Errorf("failed to read file: %w", err)
	}

	fset := token.NewFileSet()
	// Parse the file, but don't recover from errors to catch syntax issues
	node, err := parser.ParseFile(fset, filePath, content, 0)
	if err != nil {
		// If parsing fails, it might be due to an unnecessary semicolon.
		// We'll check the error message for a common semicolon error.
		errMsg := err.Error()
		if strings.Contains(errMsg, "unexpected semicolon") || strings.Contains(errMsg, "extra semicolon") {
			return false, fmt.Sprintf("❌ Unnecessary semicolon found: %s", errMsg), nil
		}
		return false, "", fmt.Errorf("failed to parse file: %w", err)
	}

	// If the file parses correctly, we need to inspect the AST for empty statements (which can be semicolons)
	var foundSemicolon bool
	ast.Inspect(node, func(n ast.Node) bool {
		if _, ok := n.(*ast.EmptyStmt); ok {
			// Check if the semicolon is actually necessary (e.g., in a for loop)
			// This is a simplified check; a more robust solution would involve deeper AST analysis
			pos := fset.Position(n.Pos())
			line := strings.Split(string(content), "\n")[pos.Line-1]
			if !strings.Contains(line, ";;") && !strings.Contains(line, "for") {
				foundSemicolon = true
				return false // Stop inspecting
			}
		}
		return true
	})

	if foundSemicolon {
		return false, "❌ Unnecessary semicolon found. Go's formatter usually handles this.", nil
	}

	return true, "✅ No unnecessary semicolons found.", nil
}
