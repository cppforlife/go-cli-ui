package table_test

import (
	"bytes"
	"fmt"
	"testing"

	. "github.com/cppforlife/go-cli-ui/ui/table"
	"github.com/stretchr/testify/assert"
)

func TestWriter(t *testing.T) {
	t.Run("Write/Flush", func(t *testing.T) {
		t.Run("writes single row", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			writer := NewWriter(buf, "empty", ".", "||")
			visibleHeaders := []Header{{Hidden: false}, {Hidden: false}}

			writer.Write(visibleHeaders, []Value{ValueString{S: "c0r0"}, ValueString{S: "c1r0"}})
			writer.Flush()
			assert.Equal(t, buf.String(), "c0r0||c1r0||\n")
		})

		t.Run("writes multiple rows", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			writer := NewWriter(buf, "empty", ".", "||")
			visibleHeaders := []Header{{Hidden: false}, {Hidden: false}}

			writer.Write(visibleHeaders, []Value{ValueString{S: "c0r0"}, ValueString{S: "c1r0"}})
			writer.Write(visibleHeaders, []Value{ValueString{S: "c0r1"}, ValueString{S: "c1r1"}})
			writer.Flush()
			assert.Equal(t, "\n"+buf.String(), `
c0r0||c1r0||
c0r1||c1r1||
`)
		})

		t.Run("writes multiple rows that are not filtered out", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			writer := NewWriter(buf, "empty", ".", "||")
			lastHeaderNotVisible := []Header{{Hidden: false}, {Hidden: false}, {Hidden: true}}

			writer.Write(lastHeaderNotVisible, []Value{ValueString{S: "c0r0"}, ValueString{S: "c1r0"}, ValueString{S: "c2r0"}})
			writer.Write(lastHeaderNotVisible, []Value{ValueString{S: "c0r1"}, ValueString{S: "c1r1"}, ValueString{S: "c2r1"}})
			writer.Flush()
			assert.Equal(t, "\n"+buf.String(), `
c0r0||c1r0||
c0r1||c1r1||
`)
		})

		t.Run("writes every row if not given any headers", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			writer := NewWriter(buf, "empty", ".", "||")

			writer.Write(nil, []Value{ValueString{S: "c0r0"}, ValueString{S: "c1r0"}, ValueString{S: "c1r0"}})
			writer.Write(nil, []Value{ValueString{S: "c0r1"}, ValueString{S: "c1r1"}, ValueString{S: "c2r1"}})
			writer.Flush()
			assert.Equal(t, "\n"+buf.String(), `
c0r0||c1r0||c1r0||
c0r1||c1r1||c2r1||
`)
		})

		t.Run("properly deals with multi-width columns", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			writer := NewWriter(buf, "empty", ".", "||")
			visibleHeaders := []Header{{Hidden: false}, {Hidden: false}}

			writer.Write(visibleHeaders, []Value{ValueString{S: "c0r0-extra"}, ValueString{S: "c1r0"}})
			writer.Write(visibleHeaders, []Value{ValueString{S: "c0r1"}, ValueString{S: "c1r1-extra"}})
			writer.Write(visibleHeaders, []Value{ValueString{S: "c0r2-extra-extra"}, ValueString{S: "c1r2"}})
			writer.Flush()
			assert.Equal(t, "\n"+buf.String(), `
c0r0-extra......||c1r0||
c0r1............||c1r1-extra||
c0r2-extra-extra||c1r2||
`)
		})

		t.Run("properly deals with multi-width columns and multi-line values", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			writer := NewWriter(buf, "empty", ".", "||")
			visibleHeaders := []Header{{Hidden: false}, {Hidden: false}}

			writer.Write(visibleHeaders, []Value{ValueString{S: "c0r0-extra"}, ValueString{S: "c1r0"}})
			writer.Write(visibleHeaders, []Value{ValueString{S: "c0r1\nnext-line"}, ValueString{S: "c1r1-extra"}})
			writer.Write(visibleHeaders, []Value{ValueString{S: "c0r2-extra-extra"}, ValueString{S: "c1r2\n\nother\nanother"}})
			writer.Flush()
			assert.Equal(t, "\n"+buf.String(), `
c0r0-extra......||c1r0||
c0r1............||c1r1-extra||
next-line.......||||
c0r2-extra-extra||c1r2||
................||||
................||other||
................||another||
`)
		})

		t.Run("writes empty special value if values are empty", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			writer := NewWriter(buf, "empty", ".", "||")
			visibleHeaders := []Header{{Hidden: false}, {Hidden: false}}

			writer.Write(visibleHeaders, []Value{ValueString{S: ""}, ValueNone{}})
			writer.Write(visibleHeaders, []Value{ValueString{S: "c0r1"}, ValueString{S: "c1r1"}})
			writer.Flush()
			assert.Equal(t, "\n"+buf.String(), `
empty||empty||
c0r1.||c1r1||
`)
		})

		t.Run("uses custom Fprintf for values that support it including multi-line values", func(t *testing.T) {
			buf := bytes.NewBufferString("")
			writer := NewWriter(buf, "empty", ".", "||")
			visibleHeaders := []Header{{Hidden: false}, {Hidden: false}}

			formattedRegVal := ValueFmt{
				V: ValueString{S: "c0r0"},
				Func: func(pattern string, vals ...interface{}) string {
					return fmt.Sprintf(">%s<", fmt.Sprintf(pattern, vals...))
				},
			}

			formattedMutliVal := ValueFmt{
				V: ValueString{S: "c1r1\n\nother\nanother"},
				Func: func(pattern string, vals ...interface{}) string {
					return fmt.Sprintf(">%s<", fmt.Sprintf(pattern, vals...))
				},
			}

			writer.Write(visibleHeaders, []Value{formattedRegVal, ValueString{S: "c1r0"}})
			writer.Write(visibleHeaders, []Value{ValueString{S: "c0r1"}, formattedMutliVal})
			writer.Flush()

			// Maintains original width for values -- useful for colors since they are not visible
			assert.Equal(t, "\n"+buf.String(), `
>c0r0<||c1r0||
c0r1||>c1r1<||
....||><||
....||>other<||
....||>another<||
`)
		})
	})
}
