package fields_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-fields"
)

func TestFlat_MarshalJSON_string(t *testing.T) {

	type KeyValue struct{
		Key string
		Value string
	}

	tests := []struct{
		KeyValues []KeyValue
		Expected []byte
	}{
		{
			Expected: []byte("{}"),
		},



		{
			KeyValues: []KeyValue{
				KeyValue{Key:"once",Value:"ONE"},
			},
			Expected: []byte(`{"once":"ONE"}`),
		},
		{
			KeyValues: []KeyValue{
				KeyValue{Key:"once",Value:"ONE"},
				KeyValue{Key:"twice",Value:"TWO"},
			},
			Expected: []byte(`{"once":"ONE","twice":"TWO"}`),
		},
		{
			KeyValues: []KeyValue{
				KeyValue{Key:"once",Value:"ONE"},
				KeyValue{Key:"twice",Value:"TWO"},
				KeyValue{Key:"thrice",Value:"THREE"},
			},
			Expected: []byte(`{"once":"ONE","thrice":"THREE","twice":"TWO"}`),
		},
		{
			KeyValues: []KeyValue{
				KeyValue{Key:"once",Value:"ONE"},
				KeyValue{Key:"twice",Value:"TWO"},
				KeyValue{Key:"thrice",Value:"THREE"},
				KeyValue{Key:"fource",Value:"FOUR"},
			},
			Expected: []byte(`{"fource":"FOUR","once":"ONE","thrice":"THREE","twice":"TWO"}`),
		},
	}

	for testNumber, test := range tests {

		var actualFlat fields.Flat[string]
		for _, f := range test.KeyValues {
			actualFlat.Set(f.Value, f.Key)
		}

		var actual []byte
		var err error

		actual, err = actualFlat.MarshalJSON()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("KEY-VALUES: (%d) %#v", len(test.KeyValues), test.KeyValues)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual marshaled-json is not what was expected.", testNumber)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			t.Logf("KEY-VALUES: (%d) %#v", len(test.KeyValues), test.KeyValues)
			continue
		}
	}
}
