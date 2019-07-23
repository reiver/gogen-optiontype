package shared

var (
	NullableTypeGoStringImports = map[string]string{
		"fmt":"",
	}
)

const (
	NullableTypeGoStringTmpl =
`
func (receiver NullableType) GoString() string {
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
