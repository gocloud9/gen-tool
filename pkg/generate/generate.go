package generate

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
)

// FormatNode formats an AST node and returns the formatted source code.
func FormatNode(fset *token.FileSet, node ast.Node) ([]byte, error) {
	var buf bytes.Buffer
	err := format.Node(&buf, fset, node)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// FormatSource formats Go source code using go/format.
func FormatSource(src []byte) ([]byte, error) {
	return format.Source(src)
}
