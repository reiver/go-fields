package fields_test

import (
	"testing"

	"fmt"
	"reflect"

	"github.com/reiver/go-fields"
)

func TestFlat_Unset_string(t *testing.T) {

	genValue := func(name string) string {
		return fmt.Sprintf("the value for key %q", name)
	}

	tests := []struct{
		Keys []string
		Unset []string
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
			Keys: []string{"once"},
			Unset: []string{"once"},
			Expected: []string{},
		},
		{
			Keys: []string{"once"},
			Unset: []string{"twice"},
			Expected: []string{"once"},
		},
		{
			Keys: []string{"once"},
			Unset: []string{"thrice"},
			Expected: []string{"once"},
		},
		{
			Keys: []string{"once"},
			Unset: []string{"fource"},
			Expected: []string{"once"},
		},
		{
			Keys: []string{"once"},
			Unset: []string{"once","twice","thrice","fource"},
			Expected: []string{},
		},



		{
			Keys: []string{"once","twice"},
			Expected: []string{"once","twice"},
		},
		{
			Keys: []string{"once","twice"},
			Unset: []string{"once"},
			Expected: []string{"twice"},
		},
		{
			Keys: []string{"once","twice"},
			Unset: []string{"twice",},
			Expected: []string{"once"},
		},
		{
			Keys: []string{"once","twice"},
			Unset: []string{"thrice"},
			Expected: []string{"once","twice"},
		},
		{
			Keys: []string{"once","twice"},
			Unset: []string{"fource"},
			Expected: []string{"once","twice"},
		},
		{
			Keys: []string{"once","twice"},
			Unset: []string{"once","twice","thrice","fource"},
			Expected: []string{},
		},



		{
			Keys: []string{"once","twice","thrice"},
			Expected: []string{"once","thrice","twice"},
		},
		{
			Keys: []string{"once","twice","thrice"},
			Unset: []string{"once"},
			Expected: []string{"thrice","twice"},
		},
		{
			Keys: []string{"once","twice","thrice"},
			Unset: []string{"twice"},
			Expected: []string{"once","thrice"},
		},
		{
			Keys: []string{"once","twice","thrice"},
			Unset: []string{"thrice"},
			Expected: []string{"once","twice"},
		},
		{
			Keys: []string{"once","twice","thrice"},
			Unset: []string{"fource"},
			Expected: []string{"once","thrice","twice"},
		},
		{
			Keys: []string{"once","twice","thrice"},
			Unset: []string{"once","twice","thrice","fource"},
			Expected: []string{},
		},



		{
			Keys: []string{"once","twice","thrice","fource"},
			Expected: []string{"fource","once","thrice","twice"},
		},
		{
			Keys: []string{"once","twice","thrice","fource"},
			Unset: []string{"once"},
			Expected: []string{"fource","thrice","twice"},
		},
		{
			Keys: []string{"once","twice","thrice","fource"},
			Unset: []string{"twice"},
			Expected: []string{"fource","once","thrice"},
		},
		{
			Keys: []string{"once","twice","thrice","fource"},
			Unset: []string{"thrice"},
			Expected: []string{"fource","once","twice"},
		},
		{
			Keys: []string{"once","twice","thrice","fource"},
			Unset: []string{"fource"},
			Expected: []string{"once","thrice","twice"},
		},
		{
			Keys: []string{"once","twice","thrice","fource"},
			Unset: []string{"once","twice","thrice","fource"},
			Expected: []string{},
		},
	}

	for testNumber, test := range tests {

		var actualFlat fields.Flat[string]

		for _, key := range test.Keys {
			var value string = genValue(key)

			actualFlat.Set(value, key)
		}

		for _, key := range test.Unset {
			actualFlat.Unset(key)
		}

		actual := actualFlat.Keys()
		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual flat-unset it not what was expected", testNumber)
			t.Logf("EXPECTED: (%d)\n%#v", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%#v", len(actual), actual)
			t.Logf("KEYS:       (%d)\n%#v", len(test.Keys), test.Keys)
			t.Logf("UNSET-KEYS: (%d)\n%#v", len(test.Unset), test.Unset)
			continue
		}
	}
}
