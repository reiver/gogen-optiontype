package gendriver

import (
	"testing"
)

func TestIteratorInternalIterator(t *testing.T) {

	var x Iterator = &internalIterator{} // THIS IS THE LINE THAT ACTUALLY MATTERS

	if nil == x {
		t.Errorf("This should never happen.")
	}
}
