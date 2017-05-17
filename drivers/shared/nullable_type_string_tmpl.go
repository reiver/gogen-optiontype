package shared

var (
	NullableTypeStringImports = map[string]string{}
)

const (
	NullableTypeStringTmpl =
`
func (receiver NullableType) String() (string, error) {
	if NoneNullable() == receiver {
		return "", errNoneNullable
	}
	if Null() == receiver {
		return "", errNull
	}

	return receiver.value, nil
}
`
)
