package ui_test

import (
	"bytes"
	"io"
	"testing"

	. "github.com/cppforlife/go-cli-ui/ui"
	. "github.com/cppforlife/go-cli-ui/ui/table"
	"github.com/stretchr/testify/assert"
)

func TestUI(t *testing.T) {
	t.Run("ErrorLinef", func(t *testing.T) {
		t.Run("prints to errWriter with a trailing newline", func(t *testing.T) {
			uiOutBuffer := bytes.NewBufferString("")
			uiErrBuffer := bytes.NewBufferString("")
			ui := NewWriterUI(uiOutBuffer, uiErrBuffer, NewRecordingLogger())

			ui.ErrorLinef("fake-error-line")
			assert.Equal(t, uiOutBuffer.String(), "")
			assert.Contains(t, uiErrBuffer.String(), "fake-error-line\n")
		})

		t.Run("when writing fails", func(t *testing.T) {
			t.Run("logs an error", func(t *testing.T) {
				reader, writer := io.Pipe()
				reader.Close()

				uiOutBuffer := bytes.NewBufferString("")
				logger := NewRecordingLogger()
				ui := NewWriterUI(uiOutBuffer, writer, logger)

				ui.ErrorLinef("fake-error-line")

				assert.Equal(t, uiOutBuffer.String(), "")
				assert.Contains(t, logger.ErrOut.String(), "UI.ErrorLinef failed (message='fake-error-line')")
			})
		})
	})

	t.Run("PrintLinef", func(t *testing.T) {
		t.Run("prints to outWriter with a trailing newline", func(t *testing.T) {
			uiOutBuffer := bytes.NewBufferString("")
			uiErrBuffer := bytes.NewBufferString("")
			ui := NewWriterUI(uiOutBuffer, uiErrBuffer, NewRecordingLogger())

			ui.PrintLinef("fake-line")
			assert.Contains(t, uiOutBuffer.String(), "fake-line\n")
			assert.Equal(t, uiErrBuffer.String(), "")
		})

		t.Run("when writing fails", func(t *testing.T) {
			t.Run("logs an error", func(t *testing.T) {
				reader, writer := io.Pipe()
				reader.Close()

				uiErrBuffer := bytes.NewBufferString("")
				logger := NewRecordingLogger()
				ui := NewWriterUI(writer, uiErrBuffer, logger)

				ui.PrintLinef("fake-start")

				assert.Equal(t, uiErrBuffer.String(), "")
				assert.Contains(t, logger.ErrOut.String(), "UI.PrintLinef failed (message='fake-start')")
			})
		})
	})

	t.Run("BeginLinef", func(t *testing.T) {
		t.Run("prints to outWriter", func(t *testing.T) {
			uiOutBuffer := bytes.NewBufferString("")
			uiErrBuffer := bytes.NewBufferString("")
			ui := NewWriterUI(uiOutBuffer, uiErrBuffer, NewRecordingLogger())

			ui.BeginLinef("fake-start")
			assert.Contains(t, uiOutBuffer.String(), "fake-start")
			assert.Equal(t, uiErrBuffer.String(), "")
		})

		t.Run("when writing fails", func(t *testing.T) {
			t.Run("logs an error", func(t *testing.T) {
				reader, writer := io.Pipe()
				reader.Close()

				uiErrBuffer := bytes.NewBufferString("")
				logger := NewRecordingLogger()
				ui := NewWriterUI(writer, uiErrBuffer, logger)

				ui.BeginLinef("fake-start")

				assert.Equal(t, uiErrBuffer.String(), "")
				assert.Contains(t, logger.ErrOut.String(), "UI.BeginLinef failed (message='fake-start')")
			})
		})
	})

	t.Run("EndLinef", func(t *testing.T) {
		t.Run("prints to outWriter with a trailing newline", func(t *testing.T) {
			uiOutBuffer := bytes.NewBufferString("")
			uiErrBuffer := bytes.NewBufferString("")
			ui := NewWriterUI(uiOutBuffer, uiErrBuffer, NewRecordingLogger())

			ui.EndLinef("fake-end")
			assert.Contains(t, uiOutBuffer.String(), "fake-end\n")
			assert.Equal(t, uiErrBuffer.String(), "")
		})

		t.Run("when writing fails", func(t *testing.T) {
			t.Run("logs an error", func(t *testing.T) {
				reader, writer := io.Pipe()
				reader.Close()

				uiErrBuffer := bytes.NewBufferString("")
				logger := NewRecordingLogger()
				ui := NewWriterUI(writer, uiErrBuffer, logger)

				ui.EndLinef("fake-start")

				assert.Equal(t, uiErrBuffer.String(), "")
				assert.Contains(t, logger.ErrOut.String(), "UI.EndLinef failed (message='fake-start')")
			})
		})
	})

	t.Run("PrintBlock", func(t *testing.T) {
		t.Run("prints to outWriter as is", func(t *testing.T) {
			uiOutBuffer := bytes.NewBufferString("")
			uiErrBuffer := bytes.NewBufferString("")
			ui := NewWriterUI(uiOutBuffer, uiErrBuffer, NewRecordingLogger())

			ui.PrintBlock([]byte("block"))
			assert.Equal(t, uiOutBuffer.String(), "block")
			assert.Equal(t, uiErrBuffer.String(), "")
		})

		t.Run("when writing fails", func(t *testing.T) {
			t.Run("logs an error", func(t *testing.T) {
				reader, writer := io.Pipe()
				reader.Close()

				uiErrBuffer := bytes.NewBufferString("")
				logger := NewRecordingLogger()
				ui := NewWriterUI(writer, uiErrBuffer, logger)

				ui.PrintBlock([]byte("block"))
				assert.Equal(t, uiErrBuffer.String(), "")
				assert.Contains(t, logger.ErrOut.String(), "UI.PrintBlock failed (message='block')")
			})
		})
	})

	t.Run("PrintErrorBlock", func(t *testing.T) {
		t.Run("prints to outWriter as is", func(t *testing.T) {
			uiOutBuffer := bytes.NewBufferString("")
			uiErrBuffer := bytes.NewBufferString("")
			ui := NewWriterUI(uiOutBuffer, uiErrBuffer, NewRecordingLogger())

			ui.PrintErrorBlock("block")
			assert.Equal(t, uiOutBuffer.String(), "block")
			assert.Equal(t, uiErrBuffer.String(), "")
		})

		t.Run("when writing fails", func(t *testing.T) {
			t.Run("logs an error", func(t *testing.T) {
				reader, writer := io.Pipe()
				reader.Close()

				uiErrBuffer := bytes.NewBufferString("")
				logger := NewRecordingLogger()
				ui := NewWriterUI(writer, uiErrBuffer, logger)

				ui.PrintErrorBlock("block")
				assert.Equal(t, uiErrBuffer.String(), "")
				assert.Contains(t, logger.ErrOut.String(), "UI.PrintErrorBlock failed (message='block')")
			})
		})
	})

	t.Run("PrintTable", func(t *testing.T) {
		t.Run("prints table", func(t *testing.T) {
			uiOutBuffer := bytes.NewBufferString("")
			uiErrBuffer := bytes.NewBufferString("")
			ui := NewWriterUI(uiOutBuffer, uiErrBuffer, NewRecordingLogger())

			table := Table{
				Title:   "Title",
				Content: "things",
				Header:  []Header{NewHeader("Header1"), NewHeader("Header2")},

				Rows: [][]Value{
					{ValueString{S: "r1c1"}, ValueString{S: "r1c2"}},
					{ValueString{S: "r2c1"}, ValueString{S: "r2c2"}},
				},

				Notes:         []string{"note1", "note2"},
				BackgroundStr: ".",
				BorderStr:     "|",
			}
			ui.PrintTable(table)
			assert.Equal(t, "\n"+uiOutBuffer.String(), `
Title

Header1|Header2|
r1c1...|r1c2|
r2c1...|r2c2|

note1
note2

2 things
`)
		})
	})

	t.Run("IsInteractive", func(t *testing.T) {
		t.Run("returns true", func(t *testing.T) {
			uiOutBuffer := bytes.NewBufferString("")
			uiErrBuffer := bytes.NewBufferString("")
			ui := NewWriterUI(uiOutBuffer, uiErrBuffer, NewRecordingLogger())

			assert.Equal(t, ui.IsInteractive(), true)
		})
	})

	t.Run("Flush", func(t *testing.T) {
		t.Run("does nothing", func(t *testing.T) {
			uiOutBuffer := bytes.NewBufferString("")
			uiErrBuffer := bytes.NewBufferString("")
			ui := NewWriterUI(uiOutBuffer, uiErrBuffer, NewRecordingLogger())

			assert.NotPanics(t, func() { ui.Flush() })
		})
	})
}
