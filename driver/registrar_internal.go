package gendriver

import (
	"sync"
)

type internalRegistrar struct {
	codes map[string][]Renderer
	tests map[string][]Renderer
	mutex sync.RWMutex
}

func (receiver *internalRegistrar) Iterator(params IteratorParams) (Iterator, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	codes := receiver.codes
	tests := receiver.tests
	if nil == codes && nil == tests {
		var iterator Iterator = new(internalIterator)
		return iterator, nil
	}
	c, cOK := codes[params.Type]
	t, tOK := tests[params.Type]
	if !cOK && !tOK {
		return nil, errNotFound
	}

	var it internalIterator
	if err := it.copyFrom(c); nil != err {
		return nil, errInternalError
	}
	if !params.NoTests {
		if err := it.copyFrom(t); nil != err {
			return nil, errInternalError
		}
	}

	var iterator Iterator = &it

	return iterator, nil
}

func (receiver *internalRegistrar) Register(typeName string, isTest bool, renderer Renderer) error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if !isTest {
		if nil == receiver.codes {
			receiver.codes = map[string][]Renderer{}
		}
		if _, ok := receiver.codes[typeName]; !ok {
			receiver.codes[typeName] = []Renderer{}
		}

		receiver.codes[typeName] = append(receiver.codes[typeName], renderer)
	} else {
		if nil == receiver.tests {
			receiver.tests = map[string][]Renderer{}
		}
		if _, ok := receiver.tests[typeName]; !ok {
			receiver.tests[typeName] = []Renderer{}
		}

		receiver.tests[typeName] = append(receiver.tests[typeName], renderer)
	}


	return nil
}
