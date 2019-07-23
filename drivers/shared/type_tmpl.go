package shared

var (
	TypeImports = map[string]string{
		"bytes":               "",
		"database/sql/driver": "",
		"encoding/json":       "",
		"fmt":                 "",
	}
)

const (
	TypeTmpl =
`
type Type struct {
	loaded bool
	value  {{.Type}}
}

func (receiver Type) MarshalJSON() ([]byte, error) {
	if Nothing() == receiver {
		return nil, errNothing
	}

	return json.Marshal(receiver.value)
}

func (receiver *Type) UnmarshalJSON(b []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if 0 == bytes.Compare(b, []byte{'n','u','l','l'}) {
		return fmt.Errorf("Cannot unmarshal %q into %T.", string(b), *receiver)
	}

	var target {{.Type}}

	if err := json.Unmarshal(b, &target); nil != err {
		return err
	}

	*receiver = Some(target)

	return nil
}

func (receiver Type) WhenNothing(fn func()) {
	if Nothing() == receiver {
		fn()
	}
}

func (receiver Type) WhenSome(fn func({{.Type}})) {
	if Nothing() != receiver {
		fn(receiver.value)
	}
}

func (receiver Type) Value() (driver.Value, error) {
	if Nothing() == receiver {
		return nil, errNothing
	}

	return receiver.value, nil
}

func Nothing() Type {
	return Type{}
}

func Some(value {{.Type}}) Type {
	return Type{
		value:  value,
		loaded: true,
	}
}
`
)
