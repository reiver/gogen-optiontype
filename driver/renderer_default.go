package gendriver

import (
	"github.com/reiver/go-tmpl"

	"bytes"
	"fmt"
	"io"
	"strings"
)

type DefaultRenderer struct {
	FileImports map[string]string
	FileName string
	FileTmpl string
}

func (receiver DefaultRenderer) WriterTo(params RendererParams) (string, io.WriterTo, error) {

	var buffer bytes.Buffer

	const headerTmpl =
`package {{.Pkg}}

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

`
	tmpl.Fprintt(&buffer, headerTmpl, params)

	if fileImports := receiver.FileImports; nil != fileImports && 0 < len(fileImports) {
		io.WriteString(&buffer, "import (\n")
		for path, name := range fileImports {
			switch name {
			case "":
				io.WriteString(&buffer, fmt.Sprintf("\t%q\n", path))
			case ".":
				io.WriteString(&buffer, fmt.Sprintf("\t. %q\n", path))
			case "_":
				io.WriteString(&buffer, fmt.Sprintf("\t_ %q\n", path))
			default:
				io.WriteString(&buffer, fmt.Sprintf("\t%s %q\n", name, path))
			}

		}
		io.WriteString(&buffer, ")\n")
	}

	tmpl.Fprintt(&buffer, receiver.FileTmpl, params)

	fileContent := buffer.String()

	return receiver.FileName, strings.NewReader(fileContent), nil
}
