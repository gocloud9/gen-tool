package generate

import (
	"embed"
	"github.com/gocloud9/gen-tool/pkg/parse"
	"html/template"
)

type Input struct {
	Results         *parse.Results
	Package         *parse.PackageInfo
	Struct          *parse.StructInfo
	StructField     *parse.FieldInfo
	StructMethod    *parse.FuncInfo
	Interface       *parse.InterfaceInfo
	InterfaceMethod *parse.FuncInfo
	Constant        *parse.ConstantInfo
	Var             *parse.VarInfo
	DefinedType     *parse.DefinedTypeInfo
	Alias           *parse.AliasTypeInfo
}

type Options struct {
	EmdedFS         []embed.FS // Optional embedded filesystem for templates
	Files           Files
	TemplateFuncMap template.FuncMap
}

func Execute(parseResults *parse.Results, opts Options) error {
	errs := errorGroup{}

	input := Input{
		Results:         parseResults,
		Package:         &parse.PackageInfo{},
		Struct:          &parse.StructInfo{},
		StructField:     &parse.FieldInfo{},
		StructMethod:    &parse.FuncInfo{},
		Interface:       &parse.InterfaceInfo{},
		InterfaceMethod: &parse.FuncInfo{},
		Constant:        &parse.ConstantInfo{},
		Var:             &parse.VarInfo{},
		DefinedType:     &parse.DefinedTypeInfo{},
	}

	errs.Add(opts.Files.Filter(Global).Generate(input, opts.TemplateFuncMap, opts.EmdedFS...))

	for i := range parseResults.Packages {
		input.Package = parseResults.Packages[i]

		errs.Add(opts.Files.Filter(PerPackage).Generate(input, opts.TemplateFuncMap, opts.EmdedFS...))

		for j := range parseResults.Packages[i].Structs {
			input.Struct = parseResults.Packages[i].Structs[j]

			errs.Add(opts.Files.Filter(PerStruct).Generate(input, opts.TemplateFuncMap, opts.EmdedFS...))

			for k := range parseResults.Packages[i].Structs[j].Fields {
				input.StructField = parseResults.Packages[i].Structs[j].Fields[k]

				errs.Add(opts.Files.Filter(PerStructField).Generate(input, opts.TemplateFuncMap, opts.EmdedFS...))

				input.StructField = &parse.FieldInfo{}
			}
			for k := range parseResults.Packages[i].Structs[j].Methods {
				input.StructMethod = parseResults.Packages[i].Structs[j].Methods[k]

				errs.Add(opts.Files.Filter(PerStructMethod).Generate(input, opts.TemplateFuncMap, opts.EmdedFS...))

				input.StructMethod = &parse.FuncInfo{}
			}

			input.Struct = &parse.StructInfo{}
		}

		for j := range parseResults.Packages[i].Interfaces {
			input.Interface = parseResults.Packages[i].Interfaces[j]

			errs.Add(opts.Files.Filter(PerInterface).Generate(input, opts.TemplateFuncMap, opts.EmdedFS...))
			for k := range parseResults.Packages[i].Interfaces[j].Methods {
				input.InterfaceMethod = parseResults.Packages[i].Interfaces[j].Methods[k]

				errs.Add(opts.Files.Filter(PerInterfaceMethod).Generate(input, opts.TemplateFuncMap, opts.EmdedFS...))

				input.InterfaceMethod = &parse.FuncInfo{}
			}

			input.Interface = &parse.InterfaceInfo{}
		}

		for j := range parseResults.Packages[i].Vars {
			input.Var = parseResults.Packages[i].Vars[j]

			errs.Add(opts.Files.Filter(PerVar).Generate(input, opts.TemplateFuncMap, opts.EmdedFS...))

			input.Var = &parse.VarInfo{}
		}

		for j := range parseResults.Packages[i].Constants {
			input.Constant = parseResults.Packages[i].Constants[j]

			errs.Add(opts.Files.Filter(PerConstant).Generate(input, opts.TemplateFuncMap, opts.EmdedFS...))

			input.Constant = &parse.ConstantInfo{}
		}

		for j := range parseResults.Packages[i].DefinedTypes {
			input.DefinedType = parseResults.Packages[i].DefinedTypes[j]

			errs.Add(opts.Files.Filter(PerDefinedType).Generate(input, opts.TemplateFuncMap, opts.EmdedFS...))

			input.DefinedType = &parse.DefinedTypeInfo{}
		}

		for j := range parseResults.Packages[i].Aliases {
			input.Alias = parseResults.Packages[i].Aliases[j]

			errs.Add(opts.Files.Filter(PerAlias).Generate(input, opts.TemplateFuncMap, opts.EmdedFS...))

			input.Alias = &parse.AliasTypeInfo{}
		}

		input.Package = &parse.PackageInfo{}
	}

	return errs.toError()
}
