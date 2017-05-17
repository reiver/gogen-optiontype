package shared

var (
	TypeScanInt64Imports = map[string]string{
		"fmt":     "",
		"strconv": "",
	}
)

const (
	TypeScanInt64Tmpl =
`
func (receiver *Type) Scan(src interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	switch t := src.(type) {
	case Type:
		*receiver = t
		return nil
	case int64:
		*receiver = Some(t)
		return nil
	case string:
		i64, err := strconv.ParseInt(t, 10, 64)
		if nil != err {
			return err
		}
		*receiver = Some(i64)
		return nil
	case []byte:
		s := string(t)
		i64, err := strconv.ParseInt(s, 10, 64)
		if nil != err {
			return err
		}
		*receiver = Some(i64)
		return nil
	default:
		return fmt.Errorf("Cannot scan something of type %T into an %T.", src, *receiver)
	}
}
`
)
