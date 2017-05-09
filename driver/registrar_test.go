package gendriver

import (
	"testing"
)

func TestRegistrarInternalRegistrar(t *testing.T) {

	var x Registrar = &internalRegistrar{} // THIS IS THE LINE THAT ACTUALLY MATTERS.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}
