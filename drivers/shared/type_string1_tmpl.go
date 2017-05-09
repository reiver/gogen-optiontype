package shared

const (
	TypeString1Tmpl =
`package {{.Pkg}}

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

import (
	"fmt"
)

func (receiver Type) String() string {
	if None() == receiver {
		return "{{.Pkg}}.None()"
	}

	return fmt.Sprintf("{{.Pkg}}.Some(%d)", receiver.value)
}
`
)
