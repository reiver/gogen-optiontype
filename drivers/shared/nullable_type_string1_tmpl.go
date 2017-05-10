package shared

const (
	NullableTypeString1Tmpl =
`package {{.Pkg}}

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

import (
	"fmt"
)

func (receiver NullableType) String() string {
	if NoneNullable() == receiver {
		return "{{.Pkg}}.NoneNullable()"
	}
	if Null() == receiver {
		return "{{.Pkg}}.Null()"
	}

	return fmt.Sprintf("{{.Pkg}}.SomeNullable(%d)", receiver.value)
}
`
)
