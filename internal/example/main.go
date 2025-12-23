package main

import (
	"embed"
	"fmt"
	"github.com/gocloud9/gen-tool/pkg/generate"
	"github.com/gocloud9/gen-tool/pkg/parse"
)

//go:embed ../../pkg/generate/testdata/templates
var f embed.FS

func MustReadFile() {

}

func main() {
	p := parse.Parser{}
	r, err := p.ParseDirectory("testdata")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = generate.Execute(r, generate.Options{
		EmdedFS: []embed.FS{f},
	})

	if err != nil {
		panic(err)
	}
}
