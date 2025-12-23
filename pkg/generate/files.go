package generate

import (
	"embed"
	"text/template"
)

type Files []File

func (d Files) Generate(input interface{}, funcMap template.FuncMap, fs ...embed.FS) error {
	errs := errorGroup{}

	for i := range d {
		errs.Add(d[i].Generate(input, funcMap, fs...))
	}

	return errs.toError()
}

func (d Files) Filter(generationType Type) Files {
	filtered := Files{}

	for i := range d {
		if d[i].Type == generationType {
			filtered = append(filtered, d[i])
		}
	}

	return filtered
}
