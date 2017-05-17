package shared

var (
	TypeInt64Imports = map[string]string{}
)

const (
	TypeInt64Tmpl =
`
func (receiver Type) Int64() (int64, error) {
	if None() == receiver {
		return 0, errNone
	}

	return receiver.value, nil
}
`
)
