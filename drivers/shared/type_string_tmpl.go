package shared

var (
	TypeStringImports = map[string]string{}
)

const (
	TypeStringTmpl =
`
func (receiver Type) String() (string, error) {
	if Nothing() == receiver {
		return "", errNothing
	}

	return receiver.value, nil
}
`
)
