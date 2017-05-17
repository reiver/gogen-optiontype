package gendriver

import (
	"sync"
)

type internalIterator struct {
	data []Renderer
	err   error
	indexPlusOne int
	mutex sync.RWMutex
}

func (receiver internalIterator) index() int {
	return receiver.indexPlusOne-1
}

func (receiver *internalIterator) copyFrom(p []Renderer) error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == p {
		return nil
	}

	receiver.data = append(receiver.data, p...)

	return nil
}

func (receiver internalIterator) Renderer() (Renderer, error) {

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	if err := receiver.err; nil != err {
		return nil, err
	}

	data := receiver.data
	if nil == data {
		return nil, errInternalError
	}

	index := receiver.index()
	if len(data) <= index {
		return nil, errNotFound
	}

	renderer := data[index]

	return renderer, nil
}

func (receiver *internalIterator) Err() error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	return receiver.err
}

func (receiver *internalIterator) Next() bool {
	if nil == receiver {
		return false
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if err := receiver.err; nil != err {
		return false
	}

	if data, index := receiver.data, receiver.index(); len(data) <= (1+index) {
		return false
	}

	receiver.indexPlusOne++


	return true
}
