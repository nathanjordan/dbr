package dialect

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClickHouse(t *testing.T) {
	for _, test := range []struct {
		in   string
		want string
	}{
		{
			in:   "table.col",
			want: "`table`.`col`",
		},
		{
			in:   "col",
			want: "`col`",
		},
	} {
		require.Equal(t, test.want, ClickHouse.QuoteIdent(test.in))
	}
}

func TestArrayBoundaries(t *testing.T) {
	require.Equal(t, "[", ClickHouse.EncodeArrayBegin())
	require.Equal(t, "]", ClickHouse.EncodeArrayEnd())

	require.Equal(t, "(", MySQL.EncodeArrayBegin())
	require.Equal(t, ")", MySQL.EncodeArrayEnd())
}

func TestMySQL(t *testing.T) {
	for _, test := range []struct {
		in   string
		want string
	}{
		{
			in:   "table.col",
			want: "`table`.`col`",
		},
		{
			in:   "col",
			want: "`col`",
		},
	} {
		require.Equal(t, test.want, MySQL.QuoteIdent(test.in))
	}
}

func TestPostgreSQL(t *testing.T) {
	for _, test := range []struct {
		in   string
		want string
	}{
		{
			in:   "table.col",
			want: `"table"."col"`,
		},
		{
			in:   "col",
			want: `"col"`,
		},
	} {
		require.Equal(t, test.want, PostgreSQL.QuoteIdent(test.in))
	}
}

func TestSQLite3(t *testing.T) {
	for _, test := range []struct {
		in   string
		want string
	}{
		{
			in:   "table.col",
			want: `"table"."col"`,
		},
		{
			in:   "col",
			want: `"col"`,
		},
	} {
		require.Equal(t, test.want, SQLite3.QuoteIdent(test.in))
	}
}
