package fields_test

import (
	"testing"

	"fmt"
	"reflect"

	"github.com/reiver/go-fields"
)

func TestFlat_ForEach_string(t *testing.T) {

	genValue := func(name string) string {
		return fmt.Sprintf("the value for key %q", name)
	}

	tests := []struct{
		Keys []string
		Expected []string
	}{
		{

		},



		{
			Keys: []string{"once"},
			Expected: []string{
				`once	=	the value for key "once"`,
			},
		},
		{
			Keys: []string{"once","twice"},
			Expected: []string{
				`once	=	the value for key "once"`,
				`twice	=	the value for key "twice"`,
			},
		},
		{
			Keys: []string{"once","twice","thrice"},
			Expected: []string{
				`once	=	the value for key "once"`,
				`thrice	=	the value for key "thrice"`,
				`twice	=	the value for key "twice"`,
			},
		},
		{
			Keys: []string{"once","twice","thrice","fource"},
			Expected: []string{
				`fource	=	the value for key "fource"`,
				`once	=	the value for key "once"`,
				`thrice	=	the value for key "thrice"`,
				`twice	=	the value for key "twice"`,
			},
		},
	}

	for testNumber, test := range tests {

		var actualFlat fields.Flat[string]

		for _, key := range test.Keys {
			var value string = genValue(key)

			actualFlat.Set(value, key)
		}

		var actual []string
		for _, key := range actualFlat.Keys() {
			actual = append(actual, key +"\t=\t"+ genValue(key))
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual flat-foreach it not what was expected", testNumber)
			t.Logf("EXPECTED: (%d)\n%#v", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%#v", len(actual), actual)
			t.Logf("KEYS: (%d)\n%#v", len(test.Keys), test.Keys)
			continue
		}
	}
}
