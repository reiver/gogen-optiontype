package shared

var (
	NullableTypeInt64Imports = map[string]string{}
)

const (
	NullableTypeInt64Tmpl =
`
func (receiver NullableType) Int64() (int64, error) {
	if NoneNullable() == receiver {
		return 0, errNoneNullable
	}
	if Null() == receiver {
		return 0, errNull
	}

	return receiver.value, nil
}
`
)
