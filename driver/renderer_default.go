package gendriver

import (
	"github.com/reiver/go-tmpl"

	"bytes"
	"io"
	"strings"
)

type DefaultRenderer struct {
	FileImports map[string]string
	FileIsNotTest bool
	FileName string
	FileTmpl string
}

func (receiver DefaultRenderer) Imports() map[string]string {
	fileImports := receiver.FileImports

	imports := map[string]string{}
	for k,v := range fileImports {
		imports[k] = v
	}

	return imports
}

func(receiver DefaultRenderer) IsTest() bool {

	return !receiver.FileIsNotTest

}

func (receiver DefaultRenderer) WriterTo(params RendererParams) (string, io.WriterTo, error) {

	var buffer bytes.Buffer

	if !params.NoHeader {
		WriteHeader(&buffer, receiver.FileImports, params)
	}

	tmpl.Fprintt(&buffer, receiver.FileTmpl, params)

	fileContent := buffer.String()

	return receiver.FileName, strings.NewReader(fileContent), nil
}
