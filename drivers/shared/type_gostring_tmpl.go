package shared

var (
	TypeGoStringImports = map[string]string{
		"fmt":"",
	}
)

const (
	TypeGoStringTmpl =
`
func (receiver Type) GoString() string {
	if Nothing() == receiver {
		return "{{.Pkg}}.Nothing()"
	}

	return fmt.Sprintf("{{.Pkg}}.Something(%d)", receiver.value)
}
`
)
