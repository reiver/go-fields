package fields_test

import (
	"testing"

	"fmt"
	"strings"

	"github.com/reiver/go-fields"
)

func TestHierarchical_UnsetAll_string(t *testing.T) {

	genValue := func(name ...string) string {
		return fmt.Sprintf("the value for key %q", strings.Join(name,"."))
	}

	tests := []struct{
		Value *fields.Hierarchical[string]
	}{
		{
			Value: nil,
		},
		{
			Value: new(fields.Hierarchical[string]),
		},



		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once"),
		},
		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
		},
		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
		},
		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
		},



		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
		},
		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
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
		},
	}

	for testNumber, test := range tests {

		test.Value.UnsetAll()

		if !test.Value.IsEmpty() {
			t.Errorf("For test #%d, expected it to be empty, but actually wasn't.", testNumber)
			continue
		}
	}
}
