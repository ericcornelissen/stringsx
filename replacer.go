package stringsx

import (
	"io"
	"strings"
)

// Replacer is a strings.Replacer.
type Replacer struct {
	r *strings.Replacer
}

// NewReplacer runs strings.NewReplacer.
func NewReplacer(oldnew ...string) *Replacer {
	return &Replacer{r: strings.NewReplacer(oldnew...)}
}

// Replace runs strings.Replacer.Replace.
func (replacer *Replacer) Replace(s string) string {
	return replacer.r.Replace(s)
}

// WriteString runs strings.Replacer.WriteString.
func (replacer *Replacer) WriteString(w io.Writer, s string) (n int, err error) {
	return replacer.r.WriteString(w, s)
}
