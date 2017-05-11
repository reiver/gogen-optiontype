package shared

const (
	NullableTypeScanStringTmpl =
`package {{.Pkg}}

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

import (
	"fmt"
)

func (receiver *NullableType) Scan(src interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == src {
		*receiver = Null()
		return nil
	}

	switch t := src.(type) {
	case string:
		*receiver = SomeNullable(t)
		return nil
	case []byte:
		s := string(t)
		*receiver = SomeNullable(s)
		return nil
	default:
		return fmt.Errorf("Cannot scan something of type %T into an {{.Pkg}}.Type.", src)
	}
}
`
)
