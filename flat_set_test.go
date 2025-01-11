package fields

import (
	"testing"

	"fmt"
	"reflect"
)

func TestFlat_Set_string(t *testing.T) {

	genValue := func(name string) string {
		return fmt.Sprintf("the value for key %q", name)
	}

	tests := []struct{
		Keys []string
		Expected map[string]string
	}{
		{

		},



		{
			Keys: []string{"once"},
			Expected: map[string]string{
				"once":genValue("once"),
			},
		},
		{
			Keys: []string{"once","twice"},
			Expected: map[string]string{
				"once":genValue("once"),
				"twice":genValue("twice"),
			},
		},
		{
			Keys: []string{"once","twice","thrice"},
			Expected: map[string]string{
				"once":genValue("once"),
				"twice":genValue("twice"),
				"thrice":genValue("thrice"),
			},
		},
		{
			Keys: []string{"once","twice","thrice","fource"},
			Expected: map[string]string{
				"once":genValue("once"),
				"twice":genValue("twice"),
				"thrice":genValue("thrice"),
				"fource":genValue("fource"),
			},
		},
	}

	for testNumber, test := range tests {

		var actualFlat Flat[string]

		for _, key := range test.Keys {
			var value string = genValue(key)

			actualFlat.Set(value, key)
		}

		actual := actualFlat.values
		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual flat-set it not what was expected", testNumber)
			t.Logf("EXPECTED: (%d)\n%#v", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%#v", len(actual), actual)
			t.Logf("KEYS: (%d)\n%#v", len(test.Keys), test.Keys)
			continue
		}
	}
}
