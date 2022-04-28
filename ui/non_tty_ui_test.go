package ui_test

import (
	"testing"

	. "github.com/cppforlife/go-cli-ui/ui"
	fakeui "github.com/cppforlife/go-cli-ui/ui/fakes"
	. "github.com/cppforlife/go-cli-ui/ui/table"
	"github.com/stretchr/testify/assert"
)

func TestNonTTYUI(t *testing.T) {
	t.Run("ErrorLinef", func(t *testing.T) {
		t.Run("includes in Lines", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonTTYUI(parentUI)

			ui.ErrorLinef("fake-line1")
			assert.Equal(t, len(parentUI.Said), 0)
			assert.Equal(t, parentUI.Errors, []string{"fake-line1"})
		})
	})

	t.Run("PrintLinef", func(t *testing.T) {
		t.Run("does not include in Lines", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonTTYUI(parentUI)

			ui.PrintLinef("fake-line1")
			assert.Equal(t, len(parentUI.Said), 0)
			assert.Equal(t, len(parentUI.Errors), 0)
		})
	})

	t.Run("BeginLinef", func(t *testing.T) {
		t.Run("does not include in Lines", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonTTYUI(parentUI)

			ui.BeginLinef("fake-line1")
			assert.Equal(t, len(parentUI.Said), 0)
			assert.Equal(t, len(parentUI.Errors), 0)
		})
	})

	t.Run("EndLinef", func(t *testing.T) {
		t.Run("does not include in Lines", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonTTYUI(parentUI)

			ui.EndLinef("fake-line1")
			assert.Equal(t, len(parentUI.Said), 0)
			assert.Equal(t, len(parentUI.Errors), 0)
		})
	})

	t.Run("PrintBlock", func(t *testing.T) {
		t.Run("delegates to the parent UI", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonTTYUI(parentUI)

			ui.PrintBlock([]byte("block"))
			assert.Equal(t, parentUI.Blocks, []string{"block"})
		})
	})

	t.Run("PrintErrorBlock", func(t *testing.T) {
		t.Run("delegates to the parent UI", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonTTYUI(parentUI)

			ui.PrintBlock([]byte("block"))
			assert.Equal(t, parentUI.Blocks, []string{"block"})
		})
	})

	t.Run("PrintTable", func(t *testing.T) {
		t.Run("delegates to the parent UI with re-configured table", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonTTYUI(parentUI)

			ui.PrintTable(Table{
				Title:  "title",
				Header: []Header{NewHeader("header1")},

				Notes:   []string{"note1"},
				Content: "things",

				SortBy: []ColumnSort{{Column: 1}},

				Sections: []Section{
					{
						FirstColumn: ValueString{S: "section1"},
						Rows:        [][]Value{{ValueString{S: "row1"}}},
					},
				},

				Rows: [][]Value{{ValueString{S: "row1"}}},

				FillFirstColumn: false,
				BackgroundStr:   "-",
				BorderStr:       "",
			})

			assert.Equal(t, parentUI.Table, Table{
				Title: "",
				Header: []Header{
					{Key: "header1", Title: "header1", Hidden: false},
				},
				HeaderFormatFunc: nil,

				Notes:   nil,
				Content: "",

				SortBy: []ColumnSort{{Column: 1}},

				Sections: []Section{
					{
						FirstColumn: ValueString{S: "section1"},
						Rows:        [][]Value{{ValueString{S: "row1"}}},
					},
				},

				Rows: [][]Value{{ValueString{S: "row1"}}},

				FillFirstColumn: true,
				DataOnly:        true,
				BackgroundStr:   "-",
				BorderStr:       "\t",
			})
		})
	})

	t.Run("IsInteractive", func(t *testing.T) {
		t.Run("delegates to the parent UI", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonTTYUI(parentUI)

			parentUI.Interactive = true
			assert.Equal(t, ui.IsInteractive(), true)

			parentUI.Interactive = false
			assert.Equal(t, ui.IsInteractive(), false)
		})
	})

	t.Run("Flush", func(t *testing.T) {
		t.Run("delegates to the parent UI", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonTTYUI(parentUI)

			ui.Flush()
			assert.Equal(t, parentUI.Flushed, true)
		})
	})
}
