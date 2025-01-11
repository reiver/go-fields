package fields_test

import (
	"testing"

	"fmt"
	"reflect"

	"github.com/reiver/go-fields"
)

func TestFlat_Len_string(t *testing.T) {

	genValue := func(name string) string {
		return fmt.Sprintf("the value for key %q", name)
	}

	tests := []struct{
		Keys []string
		Expected int
	}{
		{
			Expected: 0,
		},



		{
			Keys: []string{"once"},
			Expected: 1,
		},
		{
			Keys: []string{"once","twice"},
			Expected: 2,
		},
		{
			Keys: []string{"once","twice","thrice"},
			Expected: 3,
		},
		{
			Keys: []string{"once","twice","thrice","fource"},
			Expected: 4,
		},
	}

	for testNumber, test := range tests {

		var actualFlat fields.Flat[string]

		for _, key := range test.Keys {
			var value string = genValue(key)

			actualFlat.Set(value, key)
		}

		actual := actualFlat.Len()
		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual flat-len it not what was expected", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("KEYS: (%d)\n%#v", len(test.Keys), test.Keys)
			continue
		}
	}
}
