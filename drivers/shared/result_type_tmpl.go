package shared

var (
	ResultTypeImports = map[string]string{
		"bytes":               "",
		"database/sql/driver": "",
		"encoding/json":       "",
		"errors":              "",
		"fmt":                 "",
	}
)

const (
	ResultTypeTmpl =
`
type ResultType struct {
	loaded bool
	err    error
	value  {{.Type}}
}

func (receiver ResultType) errored() bool {
	return nil != receiver.err
}

func (receiver ResultType) Else(value {{.Type}}) ResultType {
	if NothingResult() == receiver {
		return SomethingResult(value)
	}
	if receiver.errored() {
		return SomethingResult(value)
	}

	return receiver
}

func (receiver ResultType) Map(fn func({{.Type}}){{.Type}}) ResultType {
	if NothingResult() == receiver {
		return receiver
	}
	if receiver.errored() {
		return receiver
	}

	return SomethingResult(fn(receiver.value))
}

func (receiver ResultType) Then(fn func({{.Type}})ResultType) ResultType {
	if NothingResult() == receiver {
		return receiver
	}
	if receiver.errored() {
		return receiver
	}

	return fn(receiver.value)
}

func (receiver ResultType) Unwrap() ({{.Type}}, bool) {
	if NothingResult() == receiver {
		var doesNotMatter {{.Type}}
		return doesNotMatter, false
	}
	if receiver.errored() {
		var doesNotMatter {{.Type}}
		return doesNotMatter, false
	}

	return receiver.value, true
}

func (receiver ResultType) WhenNothing(fn func()) {
	if NothingResult() == receiver {
		fn()
	}
}

func (receiver ResultType) WhenError(fn func(error)) {
	if receiver.errored() {
		fn(receiver.err)
	}
}

func (receiver ResultType) WhenSomething(fn func({{.Type}})) {
	if NothingResult() != receiver && !receiver.errored() {
		fn(receiver.value)
	}
}

func NothingResult() ResultType {
	return ResultType{}
}

func SomethingResult(value {{.Type}}) ResultType {
	return ResultType{
		value:  value,
		loaded: true,
	}
}

func Err(err error) ResultType {
	return ResultType{
		err:    err,
		loaded: true,
	}
}

func Error(errmsg string) ResultType {
	err := errors.New(errmsg)
	return Err(err)
}

func Errorf(format string, a ...interface{}) ResultType {
	err := fmt.Errorf(format, a...)
	return Err(err)
}
`
)
