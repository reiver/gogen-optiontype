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
	errNilReceiver     = errors.New("Nil Receiver")
	errNothing         = errors.New("{{.Pkg}}.Nothing()")
	errNothingNullable = errors.New("{{.Pkg}}.NothingNullable()")
	errNull            = errors.New("{{.Pkg}}.Null()")
)
`
)
