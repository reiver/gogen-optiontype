package shared

var (
	NullableTypeString1Imports = map[string]string{
		"fmt":"",
	}
)

const (
	NullableTypeString1Tmpl =
`
func (receiver NullableType) String() string {
	if NothingNullable() == receiver {
		return "{{.Pkg}}.NothingNullable()"
	}
	if Null() == receiver {
		return "{{.Pkg}}.Null()"
	}

	return fmt.Sprintf("{{.Pkg}}.SomeNullable(%d)", receiver.value)
}
`
)
