# gen-tool

A small Go code-generation toolkit that parses Go packages and renders templates for packages, structs, fields, interfaces, constants, vars and types.

Core packages
- `pkg/parse` — parses Go source into structured metadata (`Results`, `PackageInfo`, `StructInfo`, etc.).
- `pkg/generate` — renders templates using parsed results and writes generated files.

Requirements
- Go 1.20+ (module-aware)
- Uses `golang.org/x/tools/go/packages`

Quick install
```bash
git clone git@github.com:gocloud9/gen-tool.git
cd gen-tool
go mod tidy
go test ./...
```

Basic usage (programmatic)

Brief: parse the current directory, then supply templates (including embedded templates) and `generate.Options` to `generate.Execute`.

```go
package main

import (
	"embed"
	"text/template"
	"log"
	"strings"

	"github.com/gocloud9/gen-tool/pkg/generate"
	"github.com/gocloud9/gen-tool/pkg/parse"
)

//go:embed _testdata/templates/*.tmpl
var templatesFS embed.FS

func main() {
	p := &parse.Parser{}
	results, err := p.ParseDirectory(".")
	if err != nil {
		log.Fatal(err)
	}

	opts := generate.Options{
		EmdedFS: []embed.FS{templatesFS}, // note: field name is `EmdedFS` in `generate.Options`
		Files: generate.Files{
			{
				DestinationPath: "gen/fields/{{.Package.Name}}_{{.Struct.Name | lower}}_{{.StructField.Name | lower}}.go",
				TemplatePath:    "_testdata/templates/field_template.tmpl",
				Type:            generate.PerStructField,
			},
			{
				DestinationPath: "gen/structs/{{.Package.Name}}_{{.Struct.Name | lower}}.go",
				TemplatePath:    "_testdata/templates/struct_template.tmpl",
				Type:            generate.PerStruct,
			},
			{
				DestinationPath: "gen/packages/{{.Package.Name}}.go",
				TemplatePath:    "_testdata/templates/package_template.tmpl",
				Type:            generate.PerPackage,
			},
			{
				DestinationPath: "gen/global/global.go",
				TemplatePath:    "_testdata/templates/global_template.tmpl",
				Type:            generate.Global,
			},
		},
		TemplateFuncMap: template.FuncMap{
			"lower": func(s string) string { return strings.ToLower(s) },
		},
	}

	if err := generate.Execute(results, opts); err != nil {
		log.Fatal(err)
	}
}
```

Templates
- Put templates in a folder such as `testdata/templates`.
- Supported template types: global, per-package, per-struct, per-struct-field, per-struct-method, per-interface, per-interface-method, per-var, per-constant, per-defined-type, per-alias.
- Template input struct is `generate.Input` with fields like `Package`, `Struct`, `StructField`, `Interface`, etc.
- You can embed templates using `//go:embed` and pass embedded `embed.FS` values via `Options.EmdedFS`.

Testing
- Several parser tests live under `pkg/parse/testdata` and `pkg/parse/parse_test.go`.
- Run all tests:
```bash
go test ./...
```

Development notes
- `pkg/parse` builds an AST-based model of Go code; markers and struct tags are preserved.
- `pkg/generate` iterates `parse.Results` and applies templates to generate files.
- Important field names used in examples match code: `Options.EmdedFS` and `Options.Files`.

Contributing
- Open issues and PRs on `https://github.com/gocloud9/gen-tool`.
- Follow existing code style and add tests for parser/generator changes.

License
- MIT (add a `LICENSE` file to the repository).

Files of interest
- `pkg/parse/parse.go`
- `pkg/generate/generate.go`
- `pkg/parse/parse_test.go`
- `pkg/generate/generate_test.go`
```