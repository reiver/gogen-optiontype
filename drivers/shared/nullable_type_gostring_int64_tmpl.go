package shared

var (
	NullableTypeGoStringInt64Imports = map[string]string{
		"fmt":"",
	}
)

const (
	NullableTypeGoStringInt64Tmpl =
`
func (receiver NullableType) GoString() string {
	if NothingNullable() == receiver {
		return "{{.Pkg}}.NothingNullable()"
	}
	if Null() == receiver {
		return "{{.Pkg}}.Null()"
	}

	return fmt.Sprintf("{{.Pkg}}.SomethingNullable(%d)", receiver.value)
}
`
)
