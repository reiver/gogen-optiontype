package shared

var (
	TypeString1Imports = map[string]string{
		"fmt":"",
	}
)

const (
	TypeString1Tmpl =
`
func (receiver Type) String() string {
	if None() == receiver {
		return "{{.Pkg}}.None()"
	}

	return fmt.Sprintf("{{.Pkg}}.Some(%d)", receiver.value)
}
`
)
