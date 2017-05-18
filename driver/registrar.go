package gendriver

type Registrar interface {
	Iterator(IteratorParams) (Iterator, error)
	Register(string, bool, Renderer) error
}

type IteratorParams struct {
	NoTests bool
	Pkg     string
	Type    string
}
