package shared

var (
	NullableTypeImports = map[string]string{
		"bytes":               "",
		"database/sql/driver": "",
		"encoding/json":       "",
	}
)

const (
	NullableTypeTmpl =
`
type NullableType struct {
	loaded bool
	null   bool
	value  {{.Type}}
}

func (receiver NullableType) Map(fn func({{.Type}}){{.Type}}) NullableType {
	if NothingNullable() == receiver {
		return receiver
	}
	if Null() == receiver {
		return receiver
	}

	return SomethingNullable(fn(receiver.value))
}

func (receiver NullableType) MarshalJSON() ([]byte, error) {
	if NothingNullable() == receiver {
		return nil, errNothingNullable
	}
	if Null() == receiver {
		return json.Marshal(nil)
	}

	return json.Marshal(receiver.value)
}

func (receiver NullableType) Then(fn func({{.Type}})NullableType) NullableType {
	if NothingNullable() == receiver {
		return receiver
	}
	if Null() == receiver {
		return receiver
	}

	return fn(receiver.value)
}

func (receiver NullableType) Unwrap() ({{.Type}}, bool) {
	if NothingNullable() == receiver {
		var doesNotMatter {{.Type}}
		return doesNotMatter, false
	}
	if Null() == receiver {
		var doesNotMatter {{.Type}}
		return doesNotMatter, false
	}

	return receiver.value, true
}

func (receiver NullableType) WhenNothing(fn func()) {
	if NothingNullable() == receiver {
		fn()
	}
}

func (receiver NullableType) WhenNull(fn func()) {
	if Null() == receiver {
		fn()
	}
}

func (receiver NullableType) WhenSomething(fn func({{.Type}})) {
	if NothingNullable() != receiver && Null() != receiver {
		fn(receiver.value)
	}
}

func (receiver *NullableType) UnmarshalJSON(b []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if 0 == bytes.Compare(b, []byte{'n','u','l','l'}) {
		*receiver = Null()
		return nil
	}

	var target {{.Type}}

	if err := json.Unmarshal(b, &target); nil != err {
		return err
	}

	*receiver = SomethingNullable(target)

	return nil
}

func (receiver NullableType) Value() (driver.Value, error) {
	if NothingNullable() == receiver {
		return nil, errNothingNullable
	}
	if Null() == receiver {
		return nil, nil
	}

	return receiver.value, nil
}

func NothingNullable() NullableType {
	return NullableType{}
}

func SomethingNullable(value {{.Type}}) NullableType {
	return NullableType{
		value:  value,
		loaded: true,
	}
}

func Null() NullableType {
	return NullableType{
		null:   true,
		loaded: true,
	}
}
`
)
