package fields

import (
	"sync"
)

// Hierarchical is a thread-safe set of hierarchical key-value / name-value pairs.
type Hierarchical[T any] struct {
	mutex  sync.Mutex
	values Flat[T]
	subs   map[string]*Hierarchical[T]
}

func (receiver *Hierarchical[T]) init() {
	if nil == receiver {
		return
	}

	if nil == receiver.subs {
		receiver.subs = map[string]*Hierarchical[T]{}
	}
}

func (receiver *Hierarchical[T]) ChainSet(value T, name ...string) *Hierarchical[T] {
	if nil == receiver {
		return nil
	}

	receiver.Set(value, name...)
	return receiver
}

func (receiver *Hierarchical[T]) Equal(other *Hierarchical[T]) bool {
	if receiver == other {
		return true
	}
	if nil == receiver {
		return false
	}
	if nil == other {
		return false
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	other.mutex.Lock()
	defer other.mutex.Unlock()

	return receiver.equal(other)
}

func (receiver *Hierarchical[T]) equal(other *Hierarchical[T]) bool {
	if receiver == other {
		return true
	}
	if nil == receiver {
		return false
	}
	if nil == other {
		return false
	}

	if !receiver.values.equal(&other.values) {
		return false
	}

	if len(receiver.subs) != len(other.subs) {
		return false
	}

	for key, value1 := range receiver.subs {
		value2, found := other.subs[key]
		if !found {
			return false
		}

		if !value1.equal(value2) {
			return false
		}
	}

	return true
}

func (receiver *Hierarchical[T]) Set(value T, name ...string) {
	if nil == receiver {
		return
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	receiver.set(value, name...)
}

func (receiver *Hierarchical[T]) set(value T, name ...string) {
	if nil == receiver {
		return
	}

	receiver.init()

	var numnames int = len(name)
	if numnames <= 0 {
		return
	}

	var name0 string = name[0]
	var rest []string = name[1:]

	switch numnames {
	case 1:
		receiver.setLocal(value, name0)
	default:
		var under *Hierarchical[T]
		var found bool

		under, found = receiver.subs[name0]
		if !found {
			under = new(Hierarchical[T])
			receiver.subs[name0] = under
		}

		under.set(value, rest...)
	}
}

func (receiver *Hierarchical[T]) setLocal(value T, name string) {
	if nil == receiver {
		return
	}

	receiver.values.set(value, name)
}
