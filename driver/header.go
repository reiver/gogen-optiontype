package gendriver

import (
	"github.com/reiver/go-tmpl"

	"fmt"
	"io"
)

const (
	headerTmpl =
`package {{.Pkg}}

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

`
)

func WriteHeader(w io.Writer, imports map[string]string, params interface{}) {
	tmpl.Fprintt(w, headerTmpl, params)

	if nil != imports && 0 < len(imports) {
		io.WriteString(w, "import (\n")
		for path, name := range imports {
			switch name {
			case "":
				io.WriteString(w, fmt.Sprintf("\t%q\n", path))
			case ".":
				io.WriteString(w, fmt.Sprintf("\t. %q\n", path))
			case "_":
				io.WriteString(w, fmt.Sprintf("\t_ %q\n", path))
			default:
				io.WriteString(w, fmt.Sprintf("\t%s %q\n", name, path))
			}
		}
		io.WriteString(w, ")\n")
	}
}
