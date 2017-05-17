package shared

var (
	ErrorsImports = map[string]string{
		"errors":"",
	}
)

const (
	ErrorsTmpl =
`
var (
	errNilReceiver  = errors.New("Nil Receiver")
	errNone         = errors.New("{{.Pkg}}.None()")
	errNoneNullable = errors.New("{{.Pkg}}.NoneNullable()")
	errNull         = errors.New("{{.Pkg}}.Null()")
)
`
)
