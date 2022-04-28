package ui_test

import (
	"bytes"
	"testing"

	. "github.com/cppforlife/go-cli-ui/ui"
	fakeui "github.com/cppforlife/go-cli-ui/ui/fakes"
	. "github.com/cppforlife/go-cli-ui/ui/table"
	"github.com/stretchr/testify/assert"
)

func TestIndentingUI(t *testing.T) {
	t.Run("ErrorLinef", func(t *testing.T) {
		t.Run("delegates to the parent UI with an indent", func(t *testing.T) {
			uiOut := bytes.NewBufferString("")
			uiErr := bytes.NewBufferString("")
			parentUI := NewWriterUI(uiOut, uiErr, NewRecordingLogger())

			ui := NewIndentingUI(parentUI)
			ui.ErrorLinef("fake-error-line")
			assert.Contains(t, uiErr.String(), "  fake-error-line\n")
			assert.Equal(t, uiOut.String(), "")
		})
	})

	t.Run("PrintLinef", func(t *testing.T) {
		t.Run("delegates to the parent UI with an indent", func(t *testing.T) {
			uiOut := bytes.NewBufferString("")
			uiErr := bytes.NewBufferString("")
			parentUI := NewWriterUI(uiOut, uiErr, NewRecordingLogger())

			ui := NewIndentingUI(parentUI)
			ui.PrintLinef("fake-line")
			assert.Contains(t, uiOut.String(), "  fake-line\n")
			assert.Equal(t, uiErr.String(), "")
		})
	})

	t.Run("BeginLinef", func(t *testing.T) {
		t.Run("delegates to the parent UI with an indent", func(t *testing.T) {
			uiOut := bytes.NewBufferString("")
			uiErr := bytes.NewBufferString("")
			parentUI := NewWriterUI(uiOut, uiErr, NewRecordingLogger())

			ui := NewIndentingUI(parentUI)
			ui.BeginLinef("fake-start")
			assert.Contains(t, uiOut.String(), "  fake-start")
			assert.Equal(t, uiErr.String(), "")
		})
	})

	t.Run("EndLinef", func(t *testing.T) {
		t.Run("delegates to the parent UI", func(t *testing.T) {
			uiOut := bytes.NewBufferString("")
			uiErr := bytes.NewBufferString("")
			parentUI := NewWriterUI(uiOut, uiErr, NewRecordingLogger())

			ui := NewIndentingUI(parentUI)
			ui.EndLinef("fake-end")
			assert.Contains(t, uiOut.String(), "fake-end\n")
			assert.Equal(t, uiErr.String(), "")
		})
	})

	t.Run("PrintBlock", func(t *testing.T) {
		t.Run("delegates to the parent UI", func(t *testing.T) {
			parentFakeUI := &fakeui.FakeUI{}
			ui := NewIndentingUI(parentFakeUI)
			ui.PrintBlock([]byte("block"))
			assert.Equal(t, parentFakeUI.Blocks, []string{"block"})
		})
	})

	t.Run("PrintErrorBlock", func(t *testing.T) {
		t.Run("delegates to the parent UI", func(t *testing.T) {
			parentFakeUI := &fakeui.FakeUI{}
			ui := NewIndentingUI(parentFakeUI)
			ui.PrintBlock([]byte("block"))
			assert.Equal(t, parentFakeUI.Blocks, []string{"block"})
		})
	})

	t.Run("PrintTable", func(t *testing.T) {
		t.Run("delegates to the parent UI", func(t *testing.T) {
			parentFakeUI := &fakeui.FakeUI{}
			ui := NewIndentingUI(parentFakeUI)
			table := Table{
				Content: "things",
				Header:  []Header{NewHeader("header1")},
			}

			ui.PrintTable(table)
			assert.Equal(t, parentFakeUI.Table, table)
		})
	})

	t.Run("IsInteractive", func(t *testing.T) {
		t.Run("delegates to the parent UI", func(t *testing.T) {
			parentFakeUI := &fakeui.FakeUI{}
			ui := NewIndentingUI(parentFakeUI)

			parentFakeUI.Interactive = true
			assert.Equal(t, ui.IsInteractive(), true)

			parentFakeUI.Interactive = false
			assert.Equal(t, ui.IsInteractive(), false)
		})
	})

	t.Run("Flush", func(t *testing.T) {
		t.Run("delegates to the parent UI", func(t *testing.T) {
			parentFakeUI := &fakeui.FakeUI{}
			ui := NewIndentingUI(parentFakeUI)
			ui.Flush()
			assert.Equal(t, parentFakeUI.Flushed, true)
		})
	})
}
