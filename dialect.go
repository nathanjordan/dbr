package dbr

import "time"

// Dialect abstracts database driver differences in encoding
// types, and placeholders.
type Dialect interface {
	QuoteIdent(id string) string

	EncodeString(s string) string
	EncodeBool(b bool) string
	EncodeTime(t time.Time) string
	EncodeBytes(b []byte) string

	EncodeArrayBegin() string
	EncodeArrayEnd() string

	Placeholder(n int) string
}
