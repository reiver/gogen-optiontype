package shared

const (
	ErrorsTmpl =
`package {{.Pkg}}

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

import (
	"errors"
)

var (
	errNilReceiver  = errors.New("Nil Receiver")
	errNone         = errors.New("{{.Pkg}}.None()")
	errNoneNullable = errors.New("{{.Pkg}}.NoneNullable()")
	errNull         = errors.New("{{.Pkg}}.Null()")
)
`
)
