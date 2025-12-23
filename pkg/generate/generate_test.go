package generate_test

import (
	"fmt"
	"github.com/gocloud9/gen-tool/pkg/generate"
	"github.com/gocloud9/gen-tool/pkg/parse"
	"html/template"
	"strings"
	"testing"
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
									Name:     "AField1",
									Markers:  map[string]string{"d": "e"},
									TypeInfo: &parse.TypeInfo{},
									Tags:     map[string][]string{"a": {"b", "c"}},
								},
								"AField2": {
									Name:     "AField2",
									Markers:  map[string]string{"i": "j"},
									TypeInfo: &parse.TypeInfo{},
									Tags:     map[string][]string{"f": {"g", "h"}},
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
					DestinationPath: "gen/fields/{{.Package.Name}}_{{.Struct.Name|lower}}_{{.StructField.Name|lower}}.go",
					TemplatePath:    "testdata/templates/field_template.tmpl",
					Type:            generate.PerStructField,
				},
				{
					DestinationPath: "gen/structs/{{.Package.Name}}_{{.Struct.Name|lower}}.go",
					TemplatePath:    "testdata/templates/struct_template.tmpl",
					Type:            generate.PerStruct,
				},
				{
					DestinationPath: "gen/packages/{{.Package.Name}}.go",
					TemplatePath:    "testdata/templates/package_template.tmpl",
					Type:            generate.PerPackage,
				},
				{
					DestinationPath: "gen/global/global.go",
					TemplatePath:    "testdata/templates/global_template.tmpl",
					Type:            generate.Global,
				},
			},
		})

	fmt.Println(err)
}
