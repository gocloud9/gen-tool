package generate

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

type File struct {
	DestinationPath string
	TemplatePath    string
	FormatSource    bool
	Type            Type
}

func (d File) Generate(input interface{}, funcMap template.FuncMap, fs ...embed.FS) error {
	var tplBytes []byte
	errs := errorGroup{}

	for _, file := range fs {
		var err error
		tplBytes, err = file.ReadFile(d.TemplatePath)

		errs.Add(err)
	}

	if tplBytes == nil {
		var err error
		tplBytes, err = os.ReadFile(d.TemplatePath)
		if err != nil {
			errs.Add(err)

			return errs.toError()
		}
	}

	tmpl, err := template.New("body").Funcs(funcMap).Parse(string(tplBytes))
	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, input); err != nil {
		log.Fatalf("error executing template: %v", err)
	}

	var src []byte
	if d.FormatSource {
		src, err = format.Source(buf.Bytes())
		if err != nil {
			return err
		}
	} else {
		src = buf.Bytes()
	}

	var bufDestPath bytes.Buffer

	tmplDestPath, err := template.New("destination").Funcs(funcMap).Parse(d.DestinationPath)
	if err != nil {
		return err
	}

	if err := tmplDestPath.Execute(&bufDestPath, input); err != nil {
		log.Fatalf("error executing template: %v", err)
	}

	dir := filepath.Dir(bufDestPath.String())
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	if err := os.WriteFile(bufDestPath.String(), src, 0o644); err != nil {
		return err
	}

	return nil
}
