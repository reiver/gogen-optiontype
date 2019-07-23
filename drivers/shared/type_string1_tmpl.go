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
	if Nothing() == receiver {
		return "{{.Pkg}}.Nothing()"
	}

	return fmt.Sprintf("{{.Pkg}}.Something(%d)", receiver.value)
}
`
)
