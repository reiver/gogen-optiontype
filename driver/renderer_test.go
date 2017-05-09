package gendriver

import (
	"testing"
)

func TestRendererDefaultRenderer(t *testing.T) {

	var x Renderer = DefaultRenderer{} // THIS IS THE LINE THAT ACTUALLY MATTERS.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}
