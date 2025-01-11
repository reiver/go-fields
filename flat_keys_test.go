package fields_test

import (
	"testing"

	"fmt"
	"reflect"

	"github.com/reiver/go-fields"
)

func TestFlat_Keys_string(t *testing.T) {

	genValue := func(name string) string {
		return fmt.Sprintf("the value for key %q", name)
	}

	tests := []struct{
		Keys []string
		Expected []string
	}{
		{
			Expected: []string{},
		},



		{
			Keys: []string{"once"},
			Expected: []string{"once"},
		},
		{
			Keys: []string{"once","twice"},
			Expected: []string{"once","twice"},
		},
		{
			Keys: []string{"once","twice","thrice"},
			Expected: []string{"once","thrice","twice"},
		},
		{
			Keys: []string{"once","twice","thrice","fource"},
			Expected: []string{"fource","once","thrice","twice"},
		},
	}

	for testNumber, test := range tests {

		var actualFlat fields.Flat[string]

		for _, key := range test.Keys {
			var value string = genValue(key)

			actualFlat.Set(value, key)
		}

		actual := actualFlat.Keys()
		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual flat-keys it not what was expected", testNumber)
			t.Logf("EXPECTED: (%d)\n%#v", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%#v", len(actual), actual)
			t.Logf("KEYS: (%d)\n%#v", len(test.Keys), test.Keys)
			continue
		}
	}
}
