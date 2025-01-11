package fields

import (
	"testing"

	"fmt"
	"strings"
)

func TestHierarchical_Set_string(t *testing.T) {

	genValue := func(name ...string) string {
		return fmt.Sprintf("the value for key %q", strings.Join(name,"."))
	}

	tests := []struct{
		First  *Hierarchical[string]
		Second *Hierarchical[string]
		Expected bool
	}{
		{
			Expected: true,
		},



		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2").
				ChainSet(genValue("A","BB","CCC"), "A","BB","CCC"),
			Expected: false,
		},



		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once"),
			Expected: true,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2").
				ChainSet(genValue("A","BB","CCC"), "A","BB","CCC"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once"),
			Expected: false,
		},



		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
			Expected: true,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2").
				ChainSet(genValue("A","BB","CCC"), "A","BB","CCC"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
			Expected: false,
		},



		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
			Expected: true,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2").
				ChainSet(genValue("A","BB","CCC"), "A","BB","CCC"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
			Expected: false,
		},



		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
			Expected: true,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2").
				ChainSet(genValue("A","BB","CCC"), "A","BB","CCC"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
			Expected: false,
		},



		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Expected: true,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2").
				ChainSet(genValue("A","BB","CCC"), "A","BB","CCC"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Expected: false,
		},



		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Expected: true,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2").
				ChainSet(genValue("A","BB","CCC"), "A","BB","CCC"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Expected: false,
		},



		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Expected: true,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2").
				ChainSet(genValue("A","BB","CCC"), "A","BB","CCC"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Expected: false,
		},



		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2").
				ChainSet(genValue("A","BB","CCC"), "A","BB","CCC"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2").
				ChainSet(genValue("A","BB","CCC"), "A","BB","CCC"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2").
				ChainSet(genValue("A","BB","CCC"), "A","BB","CCC"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2").
				ChainSet(genValue("A","BB","CCC"), "A","BB","CCC"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2").
				ChainSet(genValue("A","BB","CCC"), "A","BB","CCC"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2").
				ChainSet(genValue("A","BB","CCC"), "A","BB","CCC"),
			Expected: false,
		},
		{
			First: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2").
				ChainSet(genValue("A","BB","CCC"), "A","BB","CCC"),
			Second: new(Hierarchical[string]).
				ChainSet(genValue("once"), "once").
				ChainSet(genValue("twice"), "twice").
				ChainSet(genValue("thrice"), "thrice").
				ChainSet(genValue("fource"), "fource").
				ChainSet(genValue("a","b1"), "a","b1").
				ChainSet(genValue("a","b2"), "a","b2").
				ChainSet(genValue("A","BB","CCC"), "A","BB","CCC"),
			Expected: true,
		},
	}

	for testNumber, test := range tests {

		actual := test.First.Equal(test.Second)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual hierarchical-set it not what was expected", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}
	}
}
