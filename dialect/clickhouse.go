package dialect

import (
	"time"
)

// clickhouse dialect is based on mysql dialect, but using
// [...] to represent array instead of (...).
type clickhouse struct{}

func (d clickhouse) QuoteIdent(s string) string {
	// Following the coding convention in this repo to call
	// methods on base object directly.
	return MySQL.QuoteIdent(s)
}

func (d clickhouse) EncodeString(s string) string {
	return MySQL.EncodeString(s)
}

func (d clickhouse) EncodeBool(b bool) string {
	return MySQL.EncodeBool(b)
}

func (d clickhouse) EncodeTime(t time.Time) string {
	return MySQL.EncodeTime(t)
}

func (d clickhouse) EncodeBytes(b []byte) string {
	return MySQL.EncodeBytes(b)
}

func (d clickhouse) EncodeArrayBegin() string {
	return "["
}

func (d clickhouse) EncodeArrayEnd() string {
	return "]"
}

func (d clickhouse) Placeholder(p int) string {
	return MySQL.Placeholder(p)
}
