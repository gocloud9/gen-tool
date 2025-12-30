package generate_test

import (
	"fmt"
	"github.com/gocloud9/gen-tool/pkg/generate"
	"github.com/gocloud9/gen-tool/pkg/parse"
	"strings"
	"testing"
	"text/template"
)

func TestExecute(t *testing.T) {
	err := generate.Execute(
		&parse.Results{
			Packages: map[string]*parse.PackageInfo{
				"pkg1": {
					Name: "pkg1",
					Structs: map[string]*parse.StructInfo{
						"AStruct1": {
							Name: "AStruct1",
							Fields: map[string]*parse.FieldInfo{
								"AField1": {
									Name:    "AField1",
									Markers: map[string]string{"d": "e"},
									TypeInfo: &parse.TypeInfo{
										TypeName: "string",
									},
									Tags: map[string][]string{"a": {"b", "c"}},
								},
								"AField2": {
									Name:    "AField2",
									Markers: map[string]string{"i": "j"},
									TypeInfo: &parse.TypeInfo{
										TypeName: "int",
									},
									Tags: map[string][]string{"f": {"g", "h"}},
								},
							},
						},
					},
				},
			},
		},
		generate.Options{
			TemplateFuncMap: template.FuncMap{
				"lower": func(input string) string {
					return strings.ToLower(input)
				},
			},
			Files: generate.Files{
				{
					DestinationPath: "_testdata/expected/fields/{{.Package.Name}}_{{.Struct.Name|lower}}_{{.StructField.Name|lower}}.go",
					TemplatePath:    "_testdata/templates/field_template.tmpl",
					Type:            generate.PerStructField,
				},
				{
					DestinationPath: "_testdata/expected/structs/{{.Package.Name}}_{{.Struct.Name|lower}}.go",
					TemplatePath:    "_testdata/templates/struct_template.tmpl",
					Type:            generate.PerStruct,
				},
				{
					DestinationPath: "_testdata/expected/packages/{{.Package.Name}}.go",
					TemplatePath:    "_testdata/templates/package_template.tmpl",
					Type:            generate.PerPackage,
				},
				{
					DestinationPath: "_testdata/expected/global/global.go",
					TemplatePath:    "_testdata/templates/global_template.tmpl",
					Type:            generate.Global,
				},
			},
		})

	fmt.Println(err)
}
