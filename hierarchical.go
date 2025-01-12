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

func (receiver *Hierarchical[T]) ForEach(fn func(T,...string)error) error {
        if nil == receiver {
                return nil
        }

        receiver.mutex.Lock()
        defer receiver.mutex.Unlock()

        return receiver.forEach(fn)
}

func (receiver *Hierarchical[T]) forEach(fn func(T,...string)error, prefix ...string) error {
        if nil == receiver {
                return nil
        }

	err := receiver.values.forEach(func(value T, key string)error{
                value, found := receiver.values.get(key)
                if !found {
			return nil
                }

		var a []string = append([]string(nil), prefix...)
		a = append(a, key)

                err := fn(value, a...)
                if nil != err {
                        return err
                }

		return nil
	})
	if nil != err {
		return err
	}

	var subkeys []string = receiver.subkeys()

	for _, subkey := range subkeys {
		sub, found := receiver.subs[subkey]
		if !found {
	/////////////// CONTINUE
			continue
		}

		var a []string = append([]string(nil), prefix...)
		a = append(a, subkey)

		sub.forEach(fn, a...)
	}

        return nil
}

func (receiver *Hierarchical[T]) IsEmpty() bool {
	if nil == receiver {
		return true
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	return receiver.isEmpty()
}

func (receiver *Hierarchical[T]) isEmpty() bool {
	if nil == receiver {
		return true
	}

	return receiver.values.isEmpty() && len(receiver.subs) <= 0
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

func (receiver *Hierarchical[T]) subkeys() []string {
	if nil == receiver {
		return nil
	}

	return sortedKeys(receiver.subs)
}

func (receiver *Hierarchical[T]) UnsetAll() {
	if nil == receiver {
		return
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	receiver.unsetAll()
}

func (receiver *Hierarchical[T]) unsetAll() {
	if nil == receiver {
		return
	}

	receiver.values.unsetAll()
	receiver.subs = nil
}
