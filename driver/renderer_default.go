package gendriver

import (
	"github.com/reiver/go-tmpl"

	"io"
	"strings"
)

type DefaultRenderer struct {
	FileName string
	FileTmpl string
}

func (receiver DefaultRenderer) WriterTo(params RendererParams) (string, io.WriterTo, error) {

	fileContent := tmpl.Sprintt(receiver.FileTmpl, params)

	return receiver.FileName, strings.NewReader(fileContent), nil
}
