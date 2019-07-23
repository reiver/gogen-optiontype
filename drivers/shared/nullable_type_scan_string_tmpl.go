package shared

var (
	NullableTypeScanStringImports = map[string]string{
		"fmt":"",
	}
)

const (
	NullableTypeScanStringTmpl =
`
func (receiver *NullableType) Scan(src interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == src {
		*receiver = Null()
		return nil
	}

	switch t := src.(type) {
	case NullableType:
		*receiver = t
		return nil
	case Type:
		switch t {
		case Nothing():
			*receiver = NothingNullable()
		default:
			datum, err := t.String()
			if nil != err {
				return fmt.Errorf("Problem unwrapping %T: (%T) %v", t, err, err)
			}
			*receiver = SomethingNullable(datum)
		}
		return nil
	case string:
		*receiver = SomethingNullable(t)
		return nil
	case []byte:
		s := string(t)
		*receiver = SomethingNullable(s)
		return nil
	default:
		return fmt.Errorf("Cannot scan something of type %T into an %T.", src, *receiver)
	}
}
`
)
