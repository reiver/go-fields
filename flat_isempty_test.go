package fields_test

import (
	"testing"

	"fmt"

	"github.com/reiver/go-fields"
)

func TestFlat_IsEmpty_string(t *testing.T) {

	genValue := func(name string) string {
		return fmt.Sprintf("the value for key %q", name)
	}

	tests := []struct{
		Value *fields.Flat[string]
		Expected bool
	}{
		{
			Value: nil,
			Expected: true,
		},
		{
			Value: new(fields.Flat[string]),
			Expected: true,
		},



		{
			Value: new(fields.Flat[string]).
				ChainSet(genValue("once"), "once"),
			Expected: false,
		},
		{
			Value: new(fields.Flat[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
			Expected: false,
		},
		{
			Value: new(fields.Flat[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
			Expected: false,
		},
		{
			Value: new(fields.Flat[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual := test.Value.IsEmpty()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual is-empty is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}
	}
}
