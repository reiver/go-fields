package fields_test

import (
	"testing"

	"fmt"
	"reflect"

	"github.com/reiver/go-fields"
)

func TestFlat_Equal_string(t *testing.T) {

	genValue := func(name string) string {
		return fmt.Sprintf("the value for key %q", name)
	}

	tests := []struct{
		First  *fields.Flat[string]
		Second *fields.Flat[string]
		Expected bool
	}{
		{
			Expected: true,
		},



		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once"),"once"),
			Second: new(fields.Flat[string]),
			Expected: false,
		},
		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once",),"once").
				ChainSet(genValue("twice",),"twice"),
			Second: new(fields.Flat[string]),
			Expected: false,
		},
		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once",),"once").
				ChainSet(genValue("twice",),"twice").
				ChainSet(genValue("thrice",),"thrice"),
			Second: new(fields.Flat[string]),
			Expected: false,
		},
		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once",),"once").
				ChainSet(genValue("twice",),"twice").
				ChainSet(genValue("thrice",),"thrice").
				ChainSet(genValue("fource",),"fource"),
			Second: new(fields.Flat[string]),
			Expected: false,
		},



		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once"),"once"),
			Second: new(fields.Flat[string]).
				ChainSet(genValue("once"),"once"),
			Expected: true,
		},
		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once",),"once").
				ChainSet(genValue("twice",),"twice"),
			Second: new(fields.Flat[string]).
				ChainSet(genValue("once"),"once"),
			Expected: false,
		},
		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once",),"once").
				ChainSet(genValue("twice",),"twice").
				ChainSet(genValue("thrice",),"thrice"),
			Second: new(fields.Flat[string]).
				ChainSet(genValue("once"),"once"),
			Expected: false,
		},
		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once",),"once").
				ChainSet(genValue("twice",),"twice").
				ChainSet(genValue("thrice",),"thrice").
				ChainSet(genValue("fource",),"fource"),
			Second: new(fields.Flat[string]).
				ChainSet(genValue("once"),"once"),
			Expected: false,
		},



		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once"),"once"),
			Second: new(fields.Flat[string]).
				ChainSet(genValue("once"),"once").
				ChainSet(genValue("twice"),"twice"),
			Expected: false,
		},
		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once",),"once").
				ChainSet(genValue("twice",),"twice"),
			Second: new(fields.Flat[string]).
				ChainSet(genValue("once"),"once").
				ChainSet(genValue("twice"),"twice"),
			Expected: true,
		},
		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once",),"once").
				ChainSet(genValue("twice",),"twice").
				ChainSet(genValue("thrice",),"thrice"),
			Second: new(fields.Flat[string]).
				ChainSet(genValue("once"),"once").
				ChainSet(genValue("twice"),"twice"),
			Expected: false,
		},
		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once",),"once").
				ChainSet(genValue("twice",),"twice").
				ChainSet(genValue("thrice",),"thrice").
				ChainSet(genValue("fource",),"fource"),
			Second: new(fields.Flat[string]).
				ChainSet(genValue("once"),"once").
				ChainSet(genValue("twice"),"twice"),
			Expected: false,
		},



		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once"),"once"),
			Second: new(fields.Flat[string]).
				ChainSet(genValue("once"),"once").
				ChainSet(genValue("twice"),"twice").
				ChainSet(genValue("thrice"),"thrice"),
			Expected: false,
		},
		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once",),"once").
				ChainSet(genValue("twice",),"twice"),
			Second: new(fields.Flat[string]).
				ChainSet(genValue("once"),"once").
				ChainSet(genValue("twice"),"twice").
				ChainSet(genValue("thrice"),"thrice"),
			Expected: false,
		},
		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once",),"once").
				ChainSet(genValue("twice",),"twice").
				ChainSet(genValue("thrice",),"thrice"),
			Second: new(fields.Flat[string]).
				ChainSet(genValue("once"),"once").
				ChainSet(genValue("twice"),"twice").
				ChainSet(genValue("thrice"),"thrice"),
			Expected: true,
		},
		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once",),"once").
				ChainSet(genValue("twice",),"twice").
				ChainSet(genValue("thrice",),"thrice").
				ChainSet(genValue("fource",),"fource"),
			Second: new(fields.Flat[string]).
				ChainSet(genValue("once"),"once").
				ChainSet(genValue("twice"),"twice").
				ChainSet(genValue("thrice"),"thrice"),
			Expected: false,
		},



		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once"),"once"),
			Second: new(fields.Flat[string]).
				ChainSet(genValue("once"),"once").
				ChainSet(genValue("twice"),"twice").
				ChainSet(genValue("thrice"),"thrice").
				ChainSet(genValue("fource"),"fource"),
			Expected: false,
		},
		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once",),"once").
				ChainSet(genValue("twice",),"twice"),
			Second: new(fields.Flat[string]).
				ChainSet(genValue("once"),"once").
				ChainSet(genValue("twice"),"twice").
				ChainSet(genValue("thrice"),"thrice").
				ChainSet(genValue("fource"),"fource"),
			Expected: false,
		},
		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once",),"once").
				ChainSet(genValue("twice",),"twice").
				ChainSet(genValue("thrice",),"thrice"),
			Second: new(fields.Flat[string]).
				ChainSet(genValue("once"),"once").
				ChainSet(genValue("twice"),"twice").
				ChainSet(genValue("thrice"),"thrice").
				ChainSet(genValue("fource"),"fource"),
			Expected: false,
		},
		{
			First:  new(fields.Flat[string]).
				ChainSet(genValue("once",),"once").
				ChainSet(genValue("twice",),"twice").
				ChainSet(genValue("thrice",),"thrice").
				ChainSet(genValue("fource",),"fource"),
			Second: new(fields.Flat[string]).
				ChainSet(genValue("once"),"once").
				ChainSet(genValue("twice"),"twice").
				ChainSet(genValue("thrice"),"thrice").
				ChainSet(genValue("fource"),"fource"),
			Expected: true,
		},
	}

	for testNumber, test := range tests {

		actual := test.First.Equal(test.Second)
		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual flat-equal it not what was expected", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}
	}
}
