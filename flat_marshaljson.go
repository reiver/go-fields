package fields

import (
	"github.com/reiver/go-erorr"
	"github.com/reiver/go-json"
)

func (receiver *Flat[T]) MarshalJSON() ([]byte, error) {
	if nil == receiver {
		return []byte("{}"), nil
	}

	var keys []string = receiver.Keys()

	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	bytes = append(bytes, '{')
	for index, key := range keys {
		if 0 < index {
			bytes = append(bytes, ',')
		}

		value, found := receiver.Get(key)
		if !found {
	/////////////// CONTINUE
			continue
		}

		bytes = append(bytes, json.MarshalString(key)...)
		bytes = append(bytes, ':')

		switch casted := any(value).(type) {
		case string:
			bytes = append(bytes, json.MarshalString(casted)...)
		default:
			p, err := json.Marshal(value)
			if nil != err {
				return nil, erorr.Errorf("fields: problem json-marshaling field %q: %w", key, err)
			}

			bytes = append(bytes, p...)
		}
	}
	bytes = append(bytes, '}')

	return bytes, nil
}
