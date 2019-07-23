package shared

var (
	TypeGoStringStringImports = map[string]string{
		"fmt":"",
	}
)

const (
	TypeGoStringStringTmpl =
`
func (receiver Type) GoString() string {
	if Nothing() == receiver {
		return "{{.Pkg}}.Nothing()"
	}

	return fmt.Sprintf("{{.Pkg}}.Something(%q)", receiver.value)
}
`
)
