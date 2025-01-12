package fields_test

import (
	"testing"

	"fmt"
	"strings"

	"github.com/reiver/go-fields"
)

func TestHierarchical_IsEmpty_string(t *testing.T) {

	genValue := func(name ...string) string {
		return fmt.Sprintf("the value for key %q", strings.Join(name, "."))
	}

	tests := []struct{
		Value *fields.Hierarchical[string]
		Expected bool
	}{
		{
			Value: nil,
			Expected: true,
		},
		{
			Value: new(fields.Hierarchical[string]),
			Expected: true,
		},



		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once"),
			Expected: false,
		},
		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
			Expected: false,
		},
		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
			Expected: false,
		},
		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
			Expected: false,
		},



		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Expected: false,
		},
		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Expected: false,
		},
		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2").
				ChainSet(genValue("A","BB","CCC"), "A","BB","CCC"),
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
