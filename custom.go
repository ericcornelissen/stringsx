/*
Package stringsx is an extensions of Go's standard strings package with some
additional utility functions. Namely:

	• All(elemens []string condition func(string) bool) bool
	• Any(elems []string, condition func(string) bool) bool
	• IsEmpty(s string) bool
	• MapAll(elems []string, mapping func(string) string) (mapped []string)
*/
package stringsx

// All tests for every element of its first argument if the condition holds. If
// the condition holds for every string the return value will be true, otherwise
// the return value will be false.
func All(elems []string, condition func(string) bool) bool {
	for _, s := range elems {
		if !condition(s) {
			return false
		}
	}

	return true
}

// Any tests for every element of its first argument if the condition holds. If
// the condition holds for any string the return value will be true, otherwise
// the return value will be false.
func Any(elems []string, condition func(string) bool) bool {
	for _, s := range elems {
		if condition(s) {
			return true
		}
	}

	return false
}

// IsEmpty tests whether the string s is the empty string.
func IsEmpty(s string) bool {
	return s == ""
}

// MapAll maps every element of its first argument according to the provided
// mapping function.
func MapAll(elems []string, mapping func(string) string) (mapped []string) {
	mapped = make([]string, len(elems))
	for i, s := range elems {
		mapped[i] = mapping(s)
	}

	return mapped
}
