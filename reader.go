package stringsx

import (
	"io"
	"strings"
)

// Reader is a strings.Reader.
type Reader struct {
	r *strings.Reader
}

// NewReader runs strings.NewReader.
func NewReader(s string) *Reader {
	return &Reader{r: strings.NewReader(s)}
}

// Len runs strings.Reader.Len.
func (reader *Reader) Len() int {
	return reader.r.Len()
}

// Read runs strings.Reader.Read.
func (reader *Reader) Read(b []byte) (n int, err error) {
	return reader.r.Read(b)
}

// ReadAt runs strings.Reader.ReadAt.
func (reader *Reader) ReadAt(b []byte, off int64) (n int, err error) {
	return reader.r.ReadAt(b, off)
}

// ReadByte runs strings.Reader.ReadByte.
func (reader *Reader) ReadByte() (byte, error) {
	return reader.r.ReadByte()
}

// ReadRune runs strings.Reader.ReadRune.
func (reader *Reader) ReadRune() (ch rune, size int, err error) {
	return reader.r.ReadRune()
}

// Reset runs strings.Reader.Reset.
func (reader *Reader) Reset(s string) {
	reader.r.Reset(s)
}

// Seek runs strings.Reader.Seek.
func (reader *Reader) Seek(offset int64, whence int) (int64, error) {
	return reader.r.Seek(offset, whence)
}

// Size runs strings.Reader.Size.
func (reader *Reader) Size() int64 {
	return reader.r.Size()
}

// UnreadByte runs strings.Reader.UnreadByte.
func (reader *Reader) UnreadByte() error {
	return reader.r.UnreadByte()
}

// UnreadRune runs strings.Reader.UnreadRune.
func (reader *Reader) UnreadRune() error {
	return reader.r.UnreadRune()
}

// WriteTo runs strings.Reader.WriteTo.
func (reader *Reader) WriteTo(w io.Writer) (n int64, err error) {
	return reader.r.WriteTo(w)
}
