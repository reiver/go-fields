package fields_test

import (
	"testing"

	"fmt"

	"github.com/reiver/go-fields"
)

func TestFlat_UnsetAll_string(t *testing.T) {

	genValue := func(name string) string {
		return fmt.Sprintf("the value for key %q", name)
	}

	tests := []struct{
		Value *fields.Flat[string]
	}{
		{
			Value: nil,
		},
		{
			Value: new(fields.Flat[string]),
		},



		{
			Value: new(fields.Flat[string]).
				ChainSet(genValue("once"), "once"),
		},
		{
			Value: new(fields.Flat[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
		},
		{
			Value: new(fields.Flat[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
		},
		{
			Value: new(fields.Flat[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
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
