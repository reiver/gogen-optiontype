package shared

var (
	NullableTypeScanInt64Imports = map[string]string{
		"fmt":     "",
		"strconv": "",
	}
)

const (
	NullableTypeScanInt64Tmpl =
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
			datum, err := t.Int64()
			if nil != err {
				return fmt.Errorf("Problem unwrapping %T: (%T) %v", t, err, err)
			}
			*receiver = SomeNullable(datum)
		}
		return nil
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
		return fmt.Errorf("Cannot scan something of type %T into an %T.", src, *receiver)
	}
}
`
)
