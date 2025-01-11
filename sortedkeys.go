package fields

import (
	"github.com/reiver/go-humane"
)

func sortedKeys[T any](m map[string]T) []string {
	if len(m) <= 0 {
		return []string{}
	}

	var keys []string
	for key, _ := range m {
		keys = append(keys, key)
	}

	humane.SortStrings(keys)

	return keys
}
