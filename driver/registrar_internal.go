package gendriver

import (
	"sync"
)

type internalRegistrar struct {
	m map[string][]Renderer
	mutex sync.RWMutex
}

func (receiver *internalRegistrar) Iterator(params IteratorParams) (Iterator, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	m := receiver.m
	if nil == m {
		var iterator Iterator = new(internalIterator)
		return iterator, nil
	}
	s, ok := m[params.Type]
	if !ok {
		return nil, errNotFound
	}

	var it internalIterator
	if err := it.copyFrom(s); nil != err {
		return nil, errInternalError
	}

	var iterator Iterator = &it

	return iterator, nil
}

func (receiver *internalRegistrar) Register(typeName string, renderer Renderer) error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.m {
		receiver.m = map[string][]Renderer{}
	}
	if _, ok := receiver.m[typeName]; !ok {
		receiver.m[typeName] = []Renderer{}
	}

	receiver.m[typeName] = append(receiver.m[typeName], renderer)

	return nil
}
