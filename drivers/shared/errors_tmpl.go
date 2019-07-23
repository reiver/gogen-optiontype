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
	errNilReceiver     = errors.New("{{.Pkg}: Nil Receiver")
	errNothing         = errors.New("{{.Pkg}: {{.Pkg}}.Nothing()")
	errNothingNullable = errors.New("{{.Pkg}: {{.Pkg}}.NothingNullable()")
	errNull            = errors.New("{{.Pkg}: {{.Pkg}}.Null()")
)
`
)
