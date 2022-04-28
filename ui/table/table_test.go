package table_test

import (
	"bytes"
	"strings"
	"testing"

	. "github.com/cppforlife/go-cli-ui/ui/table"
	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	t.Run("Print", func(t *testing.T) {
		t.Run("prints a table in default formatting (borders, empties, etc.)", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			table := Table{
				Content: "things",
				Header: []Header{
					NewHeader("Header1"),
					NewHeader("Header2"),
				},

				Rows: [][]Value{
					{ValueString{S: "r1c1"}, ValueString{S: "r1c2"}},
					{ValueString{S: "r2c1"}, ValueString{S: "r2c2"}},
				},

				Notes: []string{"note1", "note2"},
			}
			table.Print(buf)
			assert.Equal(t, "\n"+buf.String(), strings.Replace(`
Header1  Header2  +
r1c1     r1c2  +
r2c1     r2c2  +

note1
note2

2 things
`, "+", "", -1))
		})

		t.Run("prints a table with header if Header is specified", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			table := Table{
				Content: "things",
				Header: []Header{
					NewHeader("Header1"),
					NewHeader("Header2"),
				},

				Rows: [][]Value{
					{ValueString{S: "r1c1"}, ValueString{S: "r1c2"}},
					{ValueString{S: "r2c1"}, ValueString{S: "r2c2"}},
				},

				Notes:         []string{"note1", "note2"},
				BackgroundStr: ".",
				BorderStr:     "|",
			}
			table.Print(buf)
			assert.Equal(t, "\n"+buf.String(), `
Header1|Header2|
r1c1...|r1c2|
r2c1...|r2c2|

note1
note2

2 things
`)
		})

		t.Run("prints a table without number of records if content is not specified", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			table := Table{
				Content: "",
				Header: []Header{
					NewHeader("Header1"),
					NewHeader("Header2"),
				},

				Rows: [][]Value{
					{ValueString{S: "r1c1"}, ValueString{S: "r1c2"}},
					{ValueString{S: "r2c1"}, ValueString{S: "r2c2"}},
				},

				Notes:         []string{"note1", "note2"},
				BackgroundStr: ".",
				BorderStr:     "|",
			}
			table.Print(buf)
			assert.Equal(t, "\n"+buf.String(), `
Header1|Header2|
r1c1...|r1c2|
r2c1...|r2c2|

note1
note2
`)
		})

		t.Run("prints a table sorted based on SortBy", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			table := Table{
				SortBy: []ColumnSort{{Column: 1}, {Column: 0, Asc: true}},

				Rows: [][]Value{
					{ValueString{S: "a"}, ValueInt{I: -1}},
					{ValueString{S: "b"}, ValueInt{I: 0}},
					{ValueString{S: "d"}, ValueInt{I: 20}},
					{ValueString{S: "c"}, ValueInt{I: 20}},
					{ValueString{S: "d"}, ValueInt{I: 100}},
				},

				BackgroundStr: ".",
				BorderStr:     "|",
			}
			table.Print(buf)
			assert.Equal(t, "\n"+buf.String(), `
d|100|
c|20|
d|20|
b|0|
a|-1|
`)
		})

		t.Run("prints a table without a header if Header is not specified", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			table := Table{
				Content: "things",

				Rows: [][]Value{
					{ValueString{S: "r1c1"}, ValueString{S: "r1c2"}},
					{ValueString{S: "r2c1"}, ValueString{S: "r2c2"}},
				},

				BackgroundStr: ".",
				BorderStr:     "|",
			}
			table.Print(buf)
			assert.Equal(t, "\n"+buf.String(), `
r1c1|r1c2|
r2c1|r2c2|
`)
		})

		t.Run("prints a table with a title and a header", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			table := Table{
				Title:   "Title",
				Content: "things",
				Header: []Header{
					NewHeader("Header1"),
					NewHeader("Header2"),
				},

				Rows: [][]Value{
					{ValueString{S: "r1c1"}, ValueString{S: "r1c2"}},
					{ValueString{S: "r2c1"}, ValueString{S: "r2c2"}},
				},

				Notes:         []string{"note1", "note2"},
				BackgroundStr: ".",
				BorderStr:     "|",
			}
			table.Print(buf)
			assert.Equal(t, "\n"+buf.String(), `
Title

Header1|Header2|
r1c1...|r1c2|
r2c1...|r2c2|

note1
note2

2 things
`)
		})

		t.Run("when sections are provided", func(t *testing.T) {
			t.Run("prints a table *without* sections for now", func(t *testing.T) {
				buf := bytes.NewBufferString("")
				table := Table{
					Content: "things",
					Sections: []Section{
						{
							Rows: [][]Value{
								{ValueString{S: "r1c1"}, ValueString{S: "r1c2"}},
							},
						},
						{
							Rows: [][]Value{
								{ValueString{S: "r2c1"}, ValueString{S: "r2c2"}},
							},
						},
					},
					BackgroundStr: ".",
					BorderStr:     "|",
				}
				table.Print(buf)
				assert.Equal(t, "\n"+buf.String(), `
r1c1|r1c2|
r2c1|r2c2|
`)
			})

			t.Run("prints a table with first column set", func(t *testing.T) {
				buf := bytes.NewBufferString("")
				table := Table{
					Content: "things",
					Sections: []Section{
						{
							FirstColumn: ValueString{S: "r1c1"},

							Rows: [][]Value{
								{ValueString{S: ""}, ValueString{S: "r1c2"}},
								{ValueString{S: ""}, ValueString{S: "r2c2"}},
							},
						},
						{
							Rows: [][]Value{
								{ValueString{S: "r3c1"}, ValueString{S: "r3c2"}},
							},
						},
					},
					BackgroundStr: ".",
					BorderStr:     "|",
				}
				table.Print(buf)
				assert.Equal(t, "\n"+buf.String(), `
r1c1|r1c2|
^...|r2c2|
r3c1|r3c2|
`)
			})

			t.Run("prints a table with first column filled for all rows when option is set", func(t *testing.T) {
				buf := bytes.NewBufferString("")
				table := Table{
					Content: "things",
					Sections: []Section{
						{
							FirstColumn: ValueString{S: "r1c1"},
							Rows: [][]Value{
								{ValueString{S: ""}, ValueString{S: "r1c2"}},
								{ValueString{S: ""}, ValueString{S: "r2c2"}},
							},
						},
						{
							Rows: [][]Value{
								{ValueString{S: "r3c1"}, ValueString{S: "r3c2"}},
							},
						},
						{
							FirstColumn: ValueString{S: "r4c1"},
							Rows: [][]Value{
								{ValueString{S: ""}, ValueString{S: "r4c2"}},
								{ValueString{S: ""}, ValueString{S: "r5c2"}},
								{ValueString{S: ""}, ValueString{S: "r6c2"}},
							},
						},
					},
					FillFirstColumn: true,
					BackgroundStr:   ".",
					BorderStr:       "|",
				}
				table.Print(buf)
				assert.Equal(t, "\n"+buf.String(), `
r1c1|r1c2|
r1c1|r2c2|
r3c1|r3c2|
r4c1|r4c2|
r4c1|r5c2|
r4c1|r6c2|
`)
			})

			t.Run("prints a footer including the counts for rows in sections", func(t *testing.T) {
				buf := bytes.NewBufferString("")
				table := Table{
					Content: "things",
					Header: []Header{
						NewHeader("Header1"),
						NewHeader("Header2"),
					},
					Sections: []Section{
						{
							FirstColumn: ValueString{S: "s1c1"},
							Rows: [][]Value{
								{ValueString{S: ""}, ValueString{S: "s1r1c2"}},
								{ValueString{S: ""}, ValueString{S: "s1r2c2"}},
							},
						},
						{
							Rows: [][]Value{
								{ValueString{S: "r3c1"}, ValueString{S: "r3c2"}},
							},
						},
					},
					Rows: [][]Value{
						{ValueString{S: "r4c1"}, ValueString{S: "r4c2"}},
					},
					FillFirstColumn: true,
					BackgroundStr:   ".",
					BorderStr:       "|",
				}
				table.Print(buf)
				assert.Equal(t, "\n"+buf.String(), `
Header1|Header2|
s1c1...|s1r1c2|
s1c1...|s1r2c2|
r3c1...|r3c2|
r4c1...|r4c2|

4 things
`)
			})
		})

		t.Run("prints values in table that span multiple lines", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			table := Table{
				Content: "things",

				Rows: [][]Value{
					{ValueString{S: "r1c1"}, ValueString{S: "r1c2.1\nr1c2.2"}},
					{ValueString{S: "r2c1"}, ValueString{S: "r2c2"}},
				},

				BackgroundStr: ".",
				BorderStr:     "|",
			}
			table.Print(buf)
			assert.Equal(t, "\n"+buf.String(), `
r1c1|r1c2.1|
....|r1c2.2|
r2c1|r2c2|
`)
		})

		t.Run("removes duplicate values in the first column", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			table := Table{
				Content: "things",

				Rows: [][]Value{
					{ValueString{S: "dup"}, ValueString{S: "dup"}},
					{ValueString{S: "dup"}, ValueString{S: "dup"}},
					{ValueString{S: "dup2"}, ValueString{S: "dup"}},
					{ValueString{S: "dup2"}, ValueString{S: "dup"}},
					{ValueString{S: "other"}, ValueString{S: "dup"}},
				},

				BackgroundStr: ".",
				BorderStr:     "|",
			}
			table.Print(buf)
			assert.Equal(t, "\n"+buf.String(), `
dup..|dup|
^....|dup|
dup2.|dup|
^....|dup|
other|dup|
`)
		})

		t.Run("does not removes duplicate values in the first column if FillFirstColumn is true", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			table := Table{
				Content: "things",

				Rows: [][]Value{
					{ValueString{S: "dup"}, ValueString{S: "dup"}},
					{ValueString{S: "dup"}, ValueString{S: "dup"}},
					{ValueString{S: "dup2"}, ValueString{S: "dup"}},
					{ValueString{S: "dup2"}, ValueString{S: "dup"}},
					{ValueString{S: "other"}, ValueString{S: "dup"}},
				},

				FillFirstColumn: true,
				BackgroundStr:   ".",
				BorderStr:       "|",
			}
			table.Print(buf)
			assert.Equal(t, "\n"+buf.String(), `
dup..|dup|
dup..|dup|
dup2.|dup|
dup2.|dup|
other|dup|
`)
		})

		t.Run("removes duplicate values in the first column even with sections", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			table := Table{
				Content: "things",

				Sections: []Section{
					{
						FirstColumn: ValueString{S: "dup"},
						Rows: [][]Value{
							{ValueNone{}, ValueString{S: "dup"}},
							{ValueNone{}, ValueString{S: "dup"}},
						},
					},
					{
						FirstColumn: ValueString{S: "dup2"},
						Rows: [][]Value{
							{ValueNone{}, ValueString{S: "dup"}},
							{ValueNone{}, ValueString{S: "dup"}},
						},
					},
				},

				BackgroundStr: ".",
				BorderStr:     "|",
			}
			table.Print(buf)
			assert.Equal(t, "\n"+buf.String(), `
dup.|dup|
^...|dup|
dup2|dup|
^...|dup|
`)
		})

		t.Run("removes duplicate values in the first column after sorting", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			table := Table{
				Content: "things",

				SortBy: []ColumnSort{{Column: 1, Asc: true}},

				Rows: [][]Value{
					{ValueString{S: "dup"}, ValueInt{I: 1}},
					{ValueString{S: "dup2"}, ValueInt{I: 3}},
					{ValueString{S: "dup"}, ValueInt{I: 2}},
					{ValueString{S: "dup2"}, ValueInt{I: 4}},
					{ValueString{S: "other"}, ValueInt{I: 5}},
				},

				BackgroundStr: ".",
				BorderStr:     "|",
			}
			table.Print(buf)
			assert.Equal(t, "\n"+buf.String(), `
dup..|1|
^....|2|
dup2.|3|
^....|4|
other|5|
`)
		})

		t.Run("prints empty values as dashes", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			table := Table{
				Rows: [][]Value{
					{ValueString{S: ""}, ValueNone{}},
					{ValueString{S: ""}, ValueNone{}},
				},

				BackgroundStr: ".",
				BorderStr:     "|",
			}
			table.Print(buf)
			assert.Equal(t, "\n"+buf.String(), `
-|-|
^|-|
`)
		})

		t.Run("prints empty tables without rows and section", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			table := Table{
				Content: "content",
				Header: []Header{
					NewHeader("Header1"),
					NewHeader("Header2"),
				},
				BackgroundStr: ".",
				BorderStr:     "|",
			}
			table.Print(buf)
			assert.Equal(t, "\n"+buf.String(), `
Header1|Header2|

0 content
`)
		})

		t.Run("table has Transpose:true", func(t *testing.T) {
			t.Run("prints as transposed table", func(t *testing.T) {
				buf := bytes.NewBufferString("")
				table := Table{
					Content: "errands",
					Header: []Header{
						NewHeader("Header1"),
						NewHeader("OtherHeader2"),
						NewHeader("Header3"),
					},
					Rows: [][]Value{
						{ValueString{S: "r1c1"}, ValueString{S: "longr1c2"}, ValueString{S: "r1c3"}},
						{ValueString{S: "r2c1"}, ValueString{S: "r2c2"}, ValueString{S: "r2c3"}},
					},
					BackgroundStr: ".",
					BorderStr:     "|",
					Transpose:     true,
				}
				table.Print(buf)
				assert.Equal(t, "\n"+buf.String(), `
Header1.....|r1c1|
OtherHeader2|longr1c2|
Header3.....|r1c3|

Header1.....|r2c1|
OtherHeader2|r2c2|
Header3.....|r2c3|

2 errands
`)
			})

			t.Run("prints a filtered transposed table", func(t *testing.T) {
				buf := bytes.NewBufferString("")
				nonVisibleHeader := NewHeader("Header3")
				nonVisibleHeader.Hidden = true

				table := Table{
					Content: "errands",

					Header: []Header{
						NewHeader("Header1"),
						NewHeader("Header2"),
						nonVisibleHeader,
					},
					Rows: [][]Value{
						{ValueString{S: "v1"}, ValueString{S: "v2"}, ValueString{S: "v3"}},
					},
					BorderStr: "|",
					Transpose: true,
				}
				table.Print(buf)
				assert.Equal(t, "\n"+buf.String(), `
Header1|v1|
Header2|v2|

1 errands
`)
			})

			t.Run("when table also has a SortBy value set", func(t *testing.T) {
				t.Run("prints as transposed table with sections sorted by the SortBy", func(t *testing.T) {
					buf := bytes.NewBufferString("")
					table := Table{
						Content: "errands",
						Header: []Header{
							NewHeader("Header1"),
							NewHeader("OtherHeader2"),
							NewHeader("Header3"),
						},
						Rows: [][]Value{
							{ValueString{S: "r1c1"}, ValueString{S: "longr1c2"}, ValueString{S: "r1c3"}},
							{ValueString{S: "r2c1"}, ValueString{S: "r2c2"}, ValueString{S: "r2c3"}},
						},
						SortBy: []ColumnSort{
							{Column: 0, Asc: true},
						},
						BackgroundStr: ".",
						BorderStr:     "|",
						Transpose:     true,
					}
					table.Print(buf)
					assert.Equal(t, "\n"+buf.String(), `
Header1.....|r1c1|
OtherHeader2|longr1c2|
Header3.....|r1c3|

Header1.....|r2c1|
OtherHeader2|r2c2|
Header3.....|r2c3|

2 errands
`)
				})
			})
		})

		t.Run("when column filtering is used", func(t *testing.T) {
			t.Run("prints all non-filtered out columns", func(t *testing.T) {
				buf := bytes.NewBufferString("")
				nonVisibleHeader := NewHeader("Header3")
				nonVisibleHeader.Hidden = true

				table := Table{
					Content: "content",

					Header: []Header{
						NewHeader("Header1"),
						NewHeader("Header2"),
						nonVisibleHeader,
					},
					Rows: [][]Value{
						{ValueString{S: "v1"}, ValueString{S: "v2"}, ValueString{S: "v3"}},
					},
					BorderStr: "|",
				}
				table.Print(buf)
				assert.Equal(t, "\n"+buf.String(), `
Header1|Header2|
v1     |v2|

1 content
`)
			})
		})
	})

	t.Run("AddColumn", func(t *testing.T) {
		t.Run("returns an updated table with the new column", func(t *testing.T) {
			table := Table{
				Content: "content",
				Header: []Header{
					NewHeader("Header1"),
					NewHeader("Header2"),
				},
				Rows: [][]Value{
					{ValueString{S: "r1c1"}, ValueString{S: "r1c2"}},
					{ValueString{S: "r2c1"}, ValueString{S: "r2c2"}},
				},
				BackgroundStr: ".",
				BorderStr:     "|",
			}

			newTable := table.AddColumn("Header3", []Value{ValueString{S: "r1c3"}, ValueString{S: "r2c3"}})
			assert.Equal(t, newTable, Table{
				Content: "content",
				Header: []Header{
					NewHeader("Header1"),
					NewHeader("Header2"),
					NewHeader("Header3"),
				},
				Rows: [][]Value{
					{ValueString{S: "r1c1"}, ValueString{S: "r1c2"}, ValueString{S: "r1c3"}},
					{ValueString{S: "r2c1"}, ValueString{S: "r2c2"}, ValueString{S: "r2c3"}},
				},
				BackgroundStr: ".",
				BorderStr:     "|",
			})
		})
	})
}
