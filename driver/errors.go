package gendriver

import (
	"errors"
)

var (
	errInternalError = errors.New("Internal Error")
	errNilReceiver   = errors.New("Nil Receiver")
	errNotFound      = errors.New("Not Found")
)
