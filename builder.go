package stringsx

import "strings"

// Builder is a strings.Builder.
type Builder struct {
	sb strings.Builder
}

// Cap runs strings.Builder.Cap.
func (builder *Builder) Cap() int {
	return builder.sb.Cap()
}

// Grow runs strings.Builder.Grow.
func (builder *Builder) Grow(n int) {
	builder.sb.Grow(n)
}

// Len runs strings.Builder.Len.
func (builder *Builder) Len() int {
	return builder.sb.Len()
}

// Reset runs strings.Builder.Reset.
func (builder *Builder) Reset() {
	builder.sb.Reset()
}

// String runs strings.Builder.String.
func (builder *Builder) String() string {
	return builder.sb.String()
}

// Write runs strings.Builder.Write.
func (builder *Builder) Write(p []byte) (n int, err error) {
	return builder.sb.Write(p)
}

// WriteByte runs strings.Builder.WriteByte.
func (builder *Builder) WriteByte(c byte) error {
	return builder.sb.WriteByte(c)
}

// WriteRune runs strings.Builder.WriteRune.
func (builder *Builder) WriteRune(r rune) (n int, err error) {
	return builder.sb.WriteRune(r)
}

// WriteString runs strings.Builder.WriteString.
func (builder *Builder) WriteString(s string) (n int, err error) {
	return builder.sb.WriteString(s)
}
