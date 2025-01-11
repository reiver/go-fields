package fields

import (
	"sync"
)

// Flat is a thread-safe set key-value / name-value pairs.
//
// It is a thread-safe version of map[string]T
type Flat[T any] struct {
	mutex sync.Mutex
	values map[string]T
}

func (receiver *Flat[T]) init() {
	if nil == receiver {
		return
	}

	if nil == receiver.values {
		receiver.values = map[string]T{}
	}
}

func (receiver *Flat[T]) Get(name string) (T, bool) {
	if nil == receiver {
		var nada T
		return nada, false
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	return receiver.get(name)
}

func (receiver *Flat[T]) get(name string) (T, bool) {
	if nil == receiver {
		var nada T
		return nada, false
	}

	if len(receiver.values) <= 0 {
		var nada T
		return nada, false
	}

	value, found := receiver.values[name]
	return value, found
}

func (receiver *Flat[T]) ForEach(fn func(T,string)error) error {
	if nil == receiver {
		return nil
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	return receiver.forEach(fn)
}

func (receiver *Flat[T]) forEach(fn func(T,string)error) error {
	if nil == receiver {
		return nil
	}

	var keys []string = receiver.keys()

	for _, key := range keys {
		value, found := receiver.get(key)
		if !found {
	/////////////// CONTINUE
			continue
		}

		err := fn(value, key)
		if nil != err {
			return err
		}
	}

	return nil
}

func (receiver *Flat[T]) Keys() []string {
	if nil == receiver {
		return []string{}
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	return receiver.keys()
}

func (receiver *Flat[T]) keys() []string {
	if nil == receiver {
		return []string{}
	}

	return sortedKeys(receiver.values)
}

func (receiver *Flat[T]) Len() int {
	if nil == receiver {
		return 0
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	return receiver.len()
}

func (receiver *Flat[T]) len() int {
	if nil == receiver {
		return 0
	}

	return len(receiver.values)
}

func (receiver *Flat[T]) Set(value T, name string) {
	if nil == receiver {
		return
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	receiver.set(value, name)
}

func (receiver *Flat[T]) set(value T, name string) {
	if nil == receiver {
		return
	}

	receiver.init()

	receiver.values[name] = value
}

func (receiver *Flat[T]) Unset(name string) {
	if nil == receiver {
		return
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	receiver.unset(name)
}

func (receiver *Flat[T]) unset(name string) {
	if nil == receiver {
		return
	}

	receiver.init()

	delete(receiver.values, name)
}
