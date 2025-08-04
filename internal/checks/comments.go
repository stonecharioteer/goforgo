package checks

import (
	"fmt"
	"go/parser"
	"go/token"

	"github.com/stonecharioteer/goforgo/internal/analysis"
)

func init() {
	analysis.Register(&hasLineCommentCheck{})
}

type hasLineCommentCheck struct{}

func (c *hasLineCommentCheck) Name() string {
	return "has_line_comment"
}

func (c *hasLineCommentCheck) Description() string {
	return "Checks if the file contains at least one line comment (//)."
}

func (c *hasLineCommentCheck) Execute(filePath string) (bool, string, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return false, "", fmt.Errorf("failed to parse file: %w", err)
	}

	for _, cg := range node.Comments {
		for _, c := range cg.List {
			if c.Text[:2] == "//" {
				return true, "✅ Found a line comment!", nil
			}
		}
	}

	return false, "❌ No line comment found. Add a comment starting with //", nil
}
