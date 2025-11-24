package parse

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// ParseFile parses a Go source file and returns the AST.
func ParseFile(filename string, src interface{}) (*ast.File, *token.FileSet, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		return nil, nil, err
	}
	return file, fset, nil
}

// ParseDir parses all Go source files in a directory and returns a map of packages.
func ParseDir(fset *token.FileSet, path string) (map[string]*ast.Package, error) {
	pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return pkgs, nil
}
