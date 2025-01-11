package fields_test

import (
	"testing"

	"fmt"
	"reflect"
	"strings"

	"github.com/reiver/go-fields"
)

func TestHierarchical_ForEach_string(t *testing.T) {

	genValue := func(name ...string) string {
		return fmt.Sprintf("the value for key %q", strings.Join(name,"."))
	}

	tests := []struct{
		Value *fields.Hierarchical[string]
		Expected []string
	}{
		{

		},



		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once"),
			Expected: []string{
				`once	=	the value for key "once"`,
			},
		},
		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
			Expected: []string{
				`once	=	the value for key "once"`,
				`twice	=	the value for key "twice"`,
			},
		},
		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
			Expected: []string{
				`once	=	the value for key "once"`,
				`thrice	=	the value for key "thrice"`,
				`twice	=	the value for key "twice"`,
			},
		},
		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fouce"), "fource"),
			Expected: []string{
				`fource	=	the value for key "fource"`,
				`once	=	the value for key "once"`,
				`thrice	=	the value for key "thrice"`,
				`twice	=	the value for key "twice"`,
			},
		},
		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fouce"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Expected: []string{
				`fource	=	the value for key "fource"`,
				`once	=	the value for key "once"`,
				`thrice	=	the value for key "thrice"`,
				`twice	=	the value for key "twice"`,
				`a.b1	=	the value for key "a.b1"`,
			},
		},
		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fouce"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Expected: []string{
				`fource	=	the value for key "fource"`,
				`once	=	the value for key "once"`,
				`thrice	=	the value for key "thrice"`,
				`twice	=	the value for key "twice"`,
				`a.b1	=	the value for key "a.b1"`,
				`a.b2	=	the value for key "a.b2"`,
			},
		},
		{
			Value: new(fields.Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fouce"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2").
				ChainSet(genValue("A","BB","CCC"), "A","BB","CCC"),
			Expected: []string{
				`fource	=	the value for key "fource"`,
				`once	=	the value for key "once"`,
				`thrice	=	the value for key "thrice"`,
				`twice	=	the value for key "twice"`,
				`A.BB.CCC	=	the value for key "A.BB.CCC"`,
				`a.b1	=	the value for key "a.b1"`,
				`a.b2	=	the value for key "a.b2"`,
			},
		},
	}

	for testNumber, test := range tests {

		var actual []string
		err := test.Value.ForEach(func(value string, name ...string)error{
			actual = append(actual, strings.Join(name, ".") +"\t=\t"+ genValue(name...))
			return nil
		})
		if nil != err {
			panic(err)
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual hierarchical-foreach it not what was expected", testNumber)
			t.Logf("EXPECTED: (%d)", len(expected))
			for _, x := range expected {
				t.Logf("\t- %s", x)
			}
			t.Logf("ACTUAL:   (%d)", len(actual))
			for _, x := range actual {
				t.Logf("\t- %s", x)
			}
			continue
		}
	}
}
