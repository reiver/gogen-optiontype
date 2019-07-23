package shared

var (
	NullableTypeStringImports = map[string]string{}
)

const (
	NullableTypeStringTmpl =
`
func (receiver NullableType) String() (string, error) {
	if NothingNullable() == receiver {
		return "", errNothingNullable
	}
	if Null() == receiver {
		return "", errNull
	}

	return receiver.value, nil
}
`
)
