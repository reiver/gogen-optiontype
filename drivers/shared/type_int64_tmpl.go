package shared

var (
	TypeInt64Imports = map[string]string{}
)

const (
	TypeInt64Tmpl =
`
func (receiver Type) Int64() (int64, error) {
	if Nothing() == receiver {
		return 0, errNothing
	}

	return receiver.value, nil
}
`
)
