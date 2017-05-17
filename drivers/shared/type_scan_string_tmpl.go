package shared

var (
	TypeScanStringImports = map[string]string{
		"fmt":"",
	}
)

const (
	TypeScanStringTmpl =
`
func (receiver *Type) Scan(src interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	switch t := src.(type) {
	case Type:
		*receiver = t
		return nil
	case string:
		*receiver = Some(t)
		return nil
	case []byte:
		s := string(t)
		*receiver = Some(s)
		return nil
	default:
		return fmt.Errorf("Cannot scan something of type %T into an %T.", src, *receiver)
	}
}
`
)
