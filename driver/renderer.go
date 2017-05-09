package gendriver

import (
	"io"
)

// Renderer represents a renderer. A renderer has a single WriterTo method
// that returns a file name and an io.WriterTo.
//
// The returned io.WriterTo is meant to create a single Go source code file.
// The returned file name is meant to be the name of that single Go source code file.
type Renderer interface {
	WriterTo(RendererParams) (string, io.WriterTo, error)
}

type RendererParams struct {
	Pkg  string
	Type string
}
