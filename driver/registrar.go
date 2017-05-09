package gendriver

type Registrar interface {
	Iterator(IteratorParams) (Iterator, error)
	Register(string, Renderer) error
}

type IteratorParams struct {
	Pkg  string
	Type string
}
