package stringsx

import (
	"strings"
	"unicode"
)

// Compare runs strings.Compare.
func Compare(a, b string) int {
	return strings.Compare(a, b)
}

// Contains runs strings.Contains.
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// ContainsAny runs strings.ContainsAny.
func ContainsAny(s, chars string) bool {
	return strings.ContainsAny(s, chars)
}

// ContainsRune runs strings.ContainsRune.
func ContainsRune(s string, r rune) bool {
	return strings.ContainsRune(s, r)
}

// Count runs strings.Count.
func Count(s, substr string) int {
	return strings.Count(s, substr)
}

// EqualFold runs strings.EqualFold.
func EqualFold(s, t string) bool {
	return strings.EqualFold(s, t)
}

// Fields runs strings.Fields.
func Fields(s string) []string {
	return strings.Fields(s)
}

// FieldsFunc runs strings.FieldsFunc.
func FieldsFunc(s string, f func(rune) bool) []string {
	return strings.FieldsFunc(s, f)
}

// HasPrefix runs strings.HasPrefix.
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffix runs strings.HasSuffix.
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// Index runs strings.Index.
func Index(s, substr string) int {
	return strings.Index(s, substr)
}

// IndexAny runs strings.IndexAny.
func IndexAny(s, chars string) int {
	return strings.IndexAny(s, chars)
}

// IndexByte runs strings.IndexByte.
func IndexByte(s string, c byte) int {
	return strings.IndexByte(s, c)
}

// IndexFunc runs strings.IndexFunc.
func IndexFunc(s string, f func(rune) bool) int {
	return strings.IndexFunc(s, f)
}

// IndexRune runs strings.IndexRune.
func IndexRune(s string, r rune) int {
	return strings.IndexRune(s, r)
}

// Join runs strings.Join.
func Join(elems []string, sep string) string {
	return strings.Join(elems, sep)
}

// LastIndex runs strings.LastIndex.
func LastIndex(s, substr string) int {
	return strings.LastIndex(s, substr)
}

// LastIndexAny runs strings.LastIndexAny.
func LastIndexAny(s, chars string) int {
	return strings.LastIndexAny(s, chars)
}

// LastIndexByte runs strings.LastIndexByte.
func LastIndexByte(s string, c byte) int {
	return strings.LastIndexByte(s, c)
}

// LastIndexFunc runs strings.LastIndexFunc.
func LastIndexFunc(s string, f func(rune) bool) int {
	return strings.LastIndexFunc(s, f)
}

// Map runs strings.Map.
func Map(mapping func(rune) rune, s string) string {
	return strings.Map(mapping, s)
}

// Repeat runs strings.Repeat.
func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// Replace runs strings.Replace.
func Replace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// ReplaceAll runs strings.ReplaceAll.
func ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// Split runs strings.Split.
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// SplitAfter runs strings.SplitAfter.
func SplitAfter(s, sep string) []string {
	return strings.SplitAfter(s, sep)
}

// SplitAfterN runs strings.SplitAfterN.
func SplitAfterN(s, sep string, n int) []string {
	return strings.SplitAfterN(s, sep, n)
}

// SplitN runs strings.SplitN.
func SplitN(s, sep string, n int) []string {
	return strings.SplitN(s, sep, n)
}

// Title runs strings.Title.
func Title(s string) string {
	return strings.Title(s)
}

// ToLower runs strings.ToLower.
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToLowerSpecial runs strings.ToLowerSpecial.
func ToLowerSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToLowerSpecial(c, s)
}

// ToTitle runs strings.ToTitle.
func ToTitle(s string) string {
	return strings.ToTitle(s)
}

// ToTitleSpecial runs strings.ToTitleSpecial.
func ToTitleSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToTitleSpecial(c, s)
}

// ToUpper runs strings.ToUpper.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToUpperSpecial runs strings.ToUpperSpecial.
func ToUpperSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToUpperSpecial(c, s)
}

// ToValidUTF8 runs strings.ToValidUTF8.
func ToValidUTF8(s, replacement string) string {
	return strings.ToValidUTF8(s, replacement)
}

// Trim runs strings.Trim.
func Trim(s string, cutset string) string {
	return strings.Trim(s, cutset)
}

// TrimFunc runs strings.TrimFunc.
func TrimFunc(s string, f func(rune) bool) string {
	return strings.TrimFunc(s, f)
}

// TrimLeft runs strings.TrimLeft.
func TrimLeft(s string, cutset string) string {
	return strings.TrimLeft(s, cutset)
}

// TrimLeftFunc runs strings.TrimLeftFunc.
func TrimLeftFunc(s string, f func(rune) bool) string {
	return strings.TrimLeftFunc(s, f)
}

// TrimPrefix runs strings.TrimPrefix.
func TrimPrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

// TrimRight runs strings.TrimRight.
func TrimRight(s string, cutset string) string {
	return strings.TrimRight(s, cutset)
}

// TrimRightFunc runs strings.TrimRightFunc.
func TrimRightFunc(s string, f func(rune) bool) string {
	return strings.TrimRightFunc(s, f)
}

// TrimSpace runs strings.TrimSpace.
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// TrimSuffix runs strings.TrimSuffix.
func TrimSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}
