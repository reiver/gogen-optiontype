package shared

const (
	NullableTypeScanInt64Tmpl =
`package {{.Pkg}}

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

import (
	"fmt"
	"strconv"
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
	case int64:
		*receiver = SomeNullable(t)
		return nil
	case string:
		i64, err := strconv.ParseInt(t, 10, 64)
		if nil != err {
			return err
		}
		*receiver = SomeNullable(i64)
		return nil
	case []byte:
		s := string(t)
		i64, err := strconv.ParseInt(s, 10, 64)
		if nil != err {
			return err
		}
		*receiver = SomeNullable(i64)
		return nil
	default:
		return fmt.Errorf("Cannot scan something of type %T into an {{.Pkg}}.Type.", src)
	}
}
`
)
