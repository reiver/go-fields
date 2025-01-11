package fields_test

import (
	"testing"

	"fmt"
	"reflect"

	"github.com/reiver/go-fields"
)

func TestFlat_Get_string(t *testing.T) {

	genValue := func(name string) string {
		return fmt.Sprintf("the value for key %q", name)
	}

	tests := []struct{
		Keys []string
		GetKeys []string
		Expected []string
	}{
		{

		},
		{
			GetKeys: []string{"once","twice","thrice","fource","CHERRY","Banana","apple"},
		},



		{
			Keys: []string{"ONE"},
			GetKeys: []string{"once","twice","thrice","fource","CHERRY","Banana","apple"},
		},
		{
			Keys: []string{"ONE","TWO"},
			GetKeys: []string{"once","twice","thrice","fource","CHERRY","Banana","apple"},
		},
		{
			Keys: []string{"ONE","TWO","THREE"},
			GetKeys: []string{"once","twice","thrice","fource","CHERRY","Banana","apple"},
		},
		{
			Keys: []string{"ONE","TWO","THREE","FOUR"},
			GetKeys: []string{"once","twice","thrice","fource","CHERRY","Banana","apple"},
		},



		{
			Keys: []string{"apple"},
			GetKeys: []string{"once","twice","thrice","fource","CHERRY","Banana","apple"},
			Expected: []string{
				`the value for key "apple"`,
			},
		},
		{
			Keys: []string{"apple","Apple"},
			GetKeys: []string{"once","twice","thrice","fource","CHERRY","Banana","apple"},
			Expected: []string{
				`the value for key "apple"`,
			},
		},
		{
			Keys: []string{"apple","Apple","APPLE"},
			GetKeys: []string{"once","twice","thrice","fource","CHERRY","Banana","apple"},
			Expected: []string{
				`the value for key "apple"`,
			},
		},
		{
			Keys: []string{"apple","Apple","APPLE","banana"},
			GetKeys: []string{"once","twice","thrice","fource","CHERRY","Banana","apple"},
			Expected: []string{
				`the value for key "apple"`,
			},
		},
		{
			Keys: []string{"apple","Apple","APPLE","banana","Banana"},
			GetKeys: []string{"once","twice","thrice","fource","CHERRY","Banana","apple"},
			Expected: []string{
				`the value for key "Banana"`,
				`the value for key "apple"`,
			},
		},
		{
			Keys: []string{"apple","Apple","APPLE","banana","Banana","BANANA"},
			GetKeys: []string{"once","twice","thrice","fource","CHERRY","Banana","apple"},
			Expected: []string{
				`the value for key "Banana"`,
				`the value for key "apple"`,
			},
		},
		{
			Keys: []string{"apple","Apple","APPLE","banana","Banana","BANANA","cherry"},
			GetKeys: []string{"once","twice","thrice","fource","CHERRY","Banana","apple"},
			Expected: []string{
				`the value for key "Banana"`,
				`the value for key "apple"`,
			},
		},
		{
			Keys: []string{"apple","Apple","APPLE","banana","Banana","BANANA","cherry","Cherry"},
			GetKeys: []string{"once","twice","thrice","fource","CHERRY","Banana","apple"},
			Expected: []string{
				`the value for key "Banana"`,
				`the value for key "apple"`,
			},
		},
		{
			Keys: []string{"apple","Apple","APPLE","banana","Banana","BANANA","cherry","Cherry","CHERRY"},
			GetKeys: []string{"once","twice","thrice","fource","CHERRY","Banana","apple"},
			Expected: []string{
				`the value for key "CHERRY"`,
				`the value for key "Banana"`,
				`the value for key "apple"`,
			},
		},



/*
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
*/
	}

	for testNumber, test := range tests {

		var actualFlat fields.Flat[string]

		for _, key := range test.Keys {
			var value string = genValue(key)

			actualFlat.Set(value, key)
		}

		var actual []string
		for _, key := range test.GetKeys {
			value, found := actualFlat.Get(key)
			if found {
				actual = append(actual, value)
			}
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual flat-get it not what was expected", testNumber)
			t.Logf("EXPECTED: (%d)\n%#v", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%#v", len(actual), actual)
			t.Logf("KEYS:         (%d)\n%#v", len(test.Keys), test.Keys)
			t.Logf("GET-KEYS:     (%d)\n%#v", len(test.GetKeys), test.GetKeys)
			continue
		}
	}
}
