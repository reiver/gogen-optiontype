package gendriver

var (
	// Registry is where each Renderer is stored.
	//
	// Each Renderer is meant to create a single Go source code file.
	Registry Registrar
)

func init() {
	Registry = new(internalRegistrar)
}
