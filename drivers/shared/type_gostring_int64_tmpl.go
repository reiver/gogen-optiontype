package shared

var (
	TypeGoStringInt64Imports = map[string]string{
		"fmt":"",
	}
)

const (
	TypeGoStringInt64Tmpl =
`
func (receiver Type) GoString() string {
	if Nothing() == receiver {
		return "{{.Pkg}}.Nothing()"
	}

	return fmt.Sprintf("{{.Pkg}}.Something(%d)", receiver.value)
}
`
)
