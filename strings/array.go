package strings

import (
	"slices"
	"strings"
)

// ToArray splits str at each instance of separator and returns the resulting
// substrings. Empty substrings are preserved.
func ToArray(str string, separator string) []string {
	return strings.Split(str, separator)
}

// ToArrayNonEmpty splits str at each instance of separator and returns only
// non-empty substrings.
func ToArrayNonEmpty(str string, separator string) []string {
	return slices.DeleteFunc(strings.Split(str, separator), func(s string) bool {
		return s == ""
	})
}

// ToArrayComma splits str on commas and preserves empty substrings.
func ToArrayComma(str string) []string {
	return ToArray(str, ",")
}

// ToArrayNonEmptyComma splits str on commas and discards empty substrings.
func ToArrayNonEmptyComma(str string) []string {
	return ToArrayNonEmpty(str, ",")
}
