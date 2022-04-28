package table_test

import (
	"sort"
	"testing"

	. "github.com/cppforlife/go-cli-ui/ui/table"
	"github.com/stretchr/testify/assert"
)

func TestSorting(t *testing.T) {
	t.Run("sorts by single column in asc order", func(t *testing.T) {
		sortBy := []ColumnSort{{Column: 0, Asc: true}}
		rows := [][]Value{
			{ValueString{S: "b"}, ValueString{S: "x"}},
			{ValueString{S: "a"}, ValueString{S: "y"}},
		}

		sort.Sort(Sorting{SortBy: sortBy, Rows: rows})

		assert.Equal(t, rows, [][]Value{
			{ValueString{S: "a"}, ValueString{S: "y"}},
			{ValueString{S: "b"}, ValueString{S: "x"}},
		})
	})

	t.Run("sorts by single column in desc order", func(t *testing.T) {
		sortBy := []ColumnSort{{Column: 0, Asc: false}}
		rows := [][]Value{
			{ValueString{S: "a"}, ValueString{S: "y"}},
			{ValueString{S: "b"}, ValueString{S: "x"}},
		}

		sort.Sort(Sorting{SortBy: sortBy, Rows: rows})

		assert.Equal(t, rows, [][]Value{
			{ValueString{S: "b"}, ValueString{S: "x"}},
			{ValueString{S: "a"}, ValueString{S: "y"}},
		})
	})

	t.Run("sorts by multiple columns in asc order", func(t *testing.T) {
		sortBy := []ColumnSort{
			{Column: 0, Asc: true},
			{Column: 1, Asc: true},
		}

		rows := [][]Value{
			{ValueString{S: "b"}, ValueString{S: "x"}, ValueString{S: "2"}},
			{ValueString{S: "a"}, ValueString{S: "y"}, ValueString{S: "1"}},
			{ValueString{S: "b"}, ValueString{S: "z"}, ValueString{S: "2"}},
			{ValueString{S: "c"}, ValueString{S: "t"}, ValueString{S: "0"}},
		}

		sort.Sort(Sorting{SortBy: sortBy, Rows: rows})

		assert.Equal(t, rows, [][]Value{
			{ValueString{S: "a"}, ValueString{S: "y"}, ValueString{S: "1"}},
			{ValueString{S: "b"}, ValueString{S: "x"}, ValueString{S: "2"}},
			{ValueString{S: "b"}, ValueString{S: "z"}, ValueString{S: "2"}},
			{ValueString{S: "c"}, ValueString{S: "t"}, ValueString{S: "0"}},
		})
	})

	t.Run("sorts by multiple columns in asc and desc order", func(t *testing.T) {
		sortBy := []ColumnSort{
			{Column: 0, Asc: false},
			{Column: 1, Asc: true},
		}

		rows := [][]Value{
			{ValueString{S: "b"}, ValueString{S: "z"}, ValueString{S: "2"}},
			{ValueString{S: "a"}, ValueString{S: "x"}, ValueString{S: "1"}},
			{ValueString{S: "b"}, ValueString{S: "y"}, ValueString{S: "2"}},
			{ValueString{S: "c"}, ValueString{S: "t"}, ValueString{S: "0"}},
		}

		sort.Sort(Sorting{SortBy: sortBy, Rows: rows})

		assert.Equal(t, rows, [][]Value{
			{ValueString{S: "c"}, ValueString{S: "t"}, ValueString{S: "0"}},
			{ValueString{S: "b"}, ValueString{S: "y"}, ValueString{S: "2"}},
			{ValueString{S: "b"}, ValueString{S: "z"}, ValueString{S: "2"}},
			{ValueString{S: "a"}, ValueString{S: "x"}, ValueString{S: "1"}},
		})
	})

	t.Run("sorts real values (e.g. suffix does not count)", func(t *testing.T) {
		sortBy := []ColumnSort{
			{Column: 0, Asc: true},
			{Column: 1, Asc: true},
		}

		rows := [][]Value{
			{ValueSuffix{V: ValueString{S: "a"}, Suffix: "b"}, ValueString{S: "x"}},
			{ValueSuffix{V: ValueString{S: "a"}, Suffix: "a"}, ValueString{S: "y"}},
		}

		sort.Sort(Sorting{SortBy: sortBy, Rows: rows})

		assert.Equal(t, rows, [][]Value{
			{ValueSuffix{V: ValueString{S: "a"}, Suffix: "b"}, ValueString{S: "x"}},
			{ValueSuffix{V: ValueString{S: "a"}, Suffix: "a"}, ValueString{S: "y"}},
		})
	})
}
