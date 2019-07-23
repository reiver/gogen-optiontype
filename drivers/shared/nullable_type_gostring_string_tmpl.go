package shared

var (
	NullableTypeGoStringStringImports = map[string]string{
		"fmt":"",
	}
)

const (
	NullableTypeGoStringStringTmpl =
`
func (receiver NullableType) GoString() string {
	if NothingNullable() == receiver {
		return "{{.Pkg}}.NothingNullable()"
	}
	if Null() == receiver {
		return "{{.Pkg}}.Null()"
	}

	return fmt.Sprintf("{{.Pkg}}.SomethingNullable(%q)", receiver.value)
}
`
)
