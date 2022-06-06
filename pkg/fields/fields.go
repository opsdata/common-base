package fields

import (
	"sort"
	"strings"
)

// Fields allows you to present fields independently from their storage.
type Fields interface {
	// Has returns whether the provided field exists.
	Has(field string) (exists bool)

	// Get returns the value for the provided field.
	Get(field string) (value string)
}

// Set
// - a map of field:value
// - it implements Fields
type Set map[string]string

// Set.String()
// - return all fields listed as a human readable string
// - and conveniently it is exactly the format that ParseSelector takes
// - 示例: "foo=bar", "baz=qup,foo=bar"
func (ls Set) String() string {
	selector := make([]string, 0, len(ls))
	for key, value := range ls {
		selector = append(selector, key+"="+value)
	}
	// Sort
	sort.StringSlice(selector).Sort()
	return strings.Join(selector, ",")
}

// Has returns whether the provided field exists in the map.
func (ls Set) Has(field string) bool {
	_, exists := ls[field]
	return exists
}

// Get returns the value in the map for the provided field.
func (ls Set) Get(field string) string {
	return ls[field]
}

// AsSelector converts fields into a selectors.
func (ls Set) AsSelector() Selector {
	return SelectorFromSet(ls)
}
