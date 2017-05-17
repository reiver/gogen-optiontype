package shared

var (
	TypeStringImports = map[string]string{}
)

const (
	TypeStringTmpl =
`
func (receiver Type) String() (string, error) {
	if None() == receiver {
		return "", errNone
	}

	return receiver.value, nil
}
`
)
