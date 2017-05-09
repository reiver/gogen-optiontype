package gendriver

// Iterator is an iterator on Renderer.
type Iterator interface {
	Renderer() (Renderer, error)
	Err() error
	Next() bool
}
