package ui_test

import (
	"encoding/json"
	"testing"

	. "github.com/cppforlife/go-cli-ui/ui"
	fakeui "github.com/cppforlife/go-cli-ui/ui/fakes"
	. "github.com/cppforlife/go-cli-ui/ui/table"
	"github.com/stretchr/testify/assert"
)

func TestJSONUI(t *testing.T) {
	type tableResp struct {
		Content string
		Header  map[string]string
		Rows    []map[string]string
		Notes   []string
	}

	type uiResp struct {
		Tables []tableResp
		Blocks []string
		Lines  []string
	}

	finalOutput := func(ui UI, parentUI *fakeui.FakeUI) uiResp {
		ui.Flush()

		var val uiResp

		err := json.Unmarshal([]byte(parentUI.Blocks[0]), &val)
		if err != nil {
			panic("Unmarshaling")
		}

		return val
	}

	t.Run("ErrorLinef", func(t *testing.T) {
		t.Run("includes in Lines", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewJSONUI(parentUI, NewRecordingLogger())

			ui.ErrorLinef("fake-line1")
			ui.ErrorLinef("fake-line2")
			assert.Equal(t, finalOutput(ui, parentUI), uiResp{
				Lines: []string{"fake-line1", "fake-line2"},
			})
		})
	})

	t.Run("PrintLinef", func(t *testing.T) {
		t.Run("includes in Lines", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewJSONUI(parentUI, NewRecordingLogger())

			ui.PrintLinef("fake-line1")
			ui.PrintLinef("fake-line2")
			assert.Equal(t, finalOutput(ui, parentUI), uiResp{
				Lines: []string{"fake-line1", "fake-line2"},
			})
		})
	})

	t.Run("BeginLinef", func(t *testing.T) {
		t.Run("includes in Lines", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewJSONUI(parentUI, NewRecordingLogger())

			ui.BeginLinef("fake-line1")
			ui.BeginLinef("fake-line2")
			assert.Equal(t, finalOutput(ui, parentUI), uiResp{
				Lines: []string{"fake-line1", "fake-line2"},
			})
		})
	})

	t.Run("EndLinef", func(t *testing.T) {
		t.Run("includes in Lines", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewJSONUI(parentUI, NewRecordingLogger())

			ui.EndLinef("fake-line1")
			ui.EndLinef("fake-line2")
			assert.Equal(t, finalOutput(ui, parentUI), uiResp{
				Lines: []string{"fake-line1", "fake-line2"},
			})
		})
	})

	t.Run("PrintBlock", func(t *testing.T) {
		t.Run("includes in Blocks", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewJSONUI(parentUI, NewRecordingLogger())

			ui.PrintBlock([]byte("fake-block1"))
			ui.PrintBlock([]byte("fake-block2"))
			assert.Equal(t, finalOutput(ui, parentUI), uiResp{
				Blocks: []string{"fake-block1", "fake-block2"},
			})
		})
	})

	t.Run("PrintErrorBlock", func(t *testing.T) {
		t.Run("includes in Blocks", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewJSONUI(parentUI, NewRecordingLogger())

			ui.PrintErrorBlock("fake-block1")
			ui.PrintErrorBlock("fake-block2")
			assert.Equal(t, finalOutput(ui, parentUI), uiResp{
				Blocks: []string{"fake-block1", "fake-block2"},
			})
		})
	})

	t.Run("PrintTable", func(t *testing.T) {
		t.Run("includes table response in Tables", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewJSONUI(parentUI, NewRecordingLogger())

			table := Table{
				Content: "things",
				Header:  []Header{NewHeader("Header & ( foo )  1 "), NewHeader("Header-2 header 3")},

				Rows: [][]Value{
					{ValueString{S: "r1c1"}, ValueString{S: "r1c2"}},
					{ValueString{S: "r2c1"}, ValueString{S: "r2c2"}},
				},

				Notes: []string{"note1", "note2"},
			}

			table2 := Table{
				Content: "things2",
			}

			ui.PrintTable(table)
			ui.PrintTable(table2)

			assert.Equal(t, finalOutput(ui, parentUI), uiResp{
				Tables: []tableResp{
					{
						Content: "things",
						Header:  map[string]string{"header_foo_1": "Header & ( foo )  1 ", "header_2_header_3": "Header-2 header 3"},
						Rows: []map[string]string{
							{"header_foo_1": "r1c1", "header_2_header_3": "r1c2"},
							{"header_foo_1": "r2c1", "header_2_header_3": "r2c2"},
						},
						Notes: []string{"note1", "note2"},
					},
					{
						Content: "things2",
						Header:  map[string]string{},
						Rows:    []map[string]string{},
					},
				},
			})
		})

		t.Run("generates header keys for tables with row content and no header content", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewJSONUI(parentUI, NewRecordingLogger())

			table := Table{
				Content: "things",
				Header:  []Header{},

				Rows: [][]Value{
					{ValueString{S: "r1c1"}, ValueString{S: "r1c2"}},
					{ValueString{S: "r2c1"}, ValueString{S: "r2c2"}},
				},

				Notes: []string{"note1", "note2"},
			}

			ui.PrintTable(table)

			assert.Equal(t, finalOutput(ui, parentUI), uiResp{
				Tables: []tableResp{
					{
						Content: "things",
						Header:  map[string]string{"col_0": "", "col_1": ""},
						Rows: []map[string]string{
							{"col_0": "r1c1", "col_1": "r1c2"},
							{"col_0": "r2c1", "col_1": "r2c2"},
						},
						Notes: []string{"note1", "note2"},
					},
				},
			})
		})

		t.Run("includes Headers in Tables", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewJSONUI(parentUI, NewRecordingLogger())

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

			table2 := Table{
				Content: "things2",
			}

			ui.PrintTable(table)
			ui.PrintTable(table2)

			assert.Equal(t, finalOutput(ui, parentUI), uiResp{
				Tables: []tableResp{
					{
						Content: "things",
						Header:  map[string]string{"header1": "Header1", "header2": "Header2"},
						Rows: []map[string]string{{"header1": "r1c1", "header2": "r1c2"},
							{"header1": "r2c1", "header2": "r2c2"}},
						Notes: []string{"note1", "note2"},
					},
					{
						Content: "things2",
						Header:  map[string]string{},
						Rows:    []map[string]string{},
					},
				},
			})
		})

		t.Run("convert non-alphanumeric to _", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewJSONUI(parentUI, NewRecordingLogger())

			table := Table{
				Content: "things",
				Header: []Header{
					NewHeader("#"),
					NewHeader("foo"),
					NewHeader("$"),
				},

				Rows: [][]Value{
					{ValueString{S: "r1c1"}, ValueString{S: "r1c2"}, ValueString{S: "r1c3"}},
					{ValueString{S: "r2c1"}, ValueString{S: "r2c2"}, ValueString{S: "r2c3"}},
				},

				Notes: []string{},
			}

			ui.PrintTable(table)

			tableOutput := finalOutput(ui, parentUI)
			assert.Equal(t, len(tableOutput.Tables), 1)
			assert.Equal(t, tableOutput.Tables[0].Content, "things")
			assert.Equal(t, tableOutput.Tables[0].Header, map[string]string{"0": "#", "2": "$", "foo": "foo"})

			assert.Equal(t, len(tableOutput.Tables[0].Rows), 2)
			assert.Equal(t, tableOutput.Tables[0].Rows[0], map[string]string{"0": "r1c1", "foo": "r1c2", "2": "r1c3"})
			assert.Equal(t, tableOutput.Tables[0].Rows[1], map[string]string{"0": "r2c1", "foo": "r2c2", "2": "r2c3"})
		})

		t.Run("includes in Tables when table has sections and fills in first column", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewJSONUI(parentUI, NewRecordingLogger())

			table := Table{
				Content: "things",
				Header:  []Header{NewHeader("Header1"), NewHeader("Header2")},

				Sections: []Section{
					{
						FirstColumn: ValueString{S: "first-col"},
						Rows: [][]Value{
							{ValueString{S: ""}, ValueString{S: "r1c2"}},
							{ValueString{S: ""}, ValueString{S: "r2c2"}},
						},
					},
				},

				Notes: []string{"note1", "note2"},
			}

			ui.PrintTable(table)

			assert.Equal(t, finalOutput(ui, parentUI), uiResp{
				Tables: []tableResp{
					{
						Content: "things",
						Header:  map[string]string{"header1": "Header1", "header2": "Header2"},
						Rows: []map[string]string{{"header1": "first-col", "header2": "r1c2"},
							{"header1": "first-col", "header2": "r2c2"}},
						Notes: []string{"note1", "note2"},
					},
				},
			})
		})
	})

	t.Run("AskForText", func(t *testing.T) {
		t.Run("panics", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewJSONUI(parentUI, NewRecordingLogger())

			assert.Panics(t, func() { ui.AskForText(TextOpts{}) })
		})
	})

	t.Run("AskForPassword", func(t *testing.T) {
		t.Run("panics", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewJSONUI(parentUI, NewRecordingLogger())

			assert.Panics(t, func() { ui.AskForPassword("") })
		})
	})

	t.Run("AskForChoice", func(t *testing.T) {
		t.Run("panics", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewJSONUI(parentUI, NewRecordingLogger())

			assert.Panics(t, func() { ui.AskForChoice(ChoiceOpts{}) })
		})
	})

	t.Run("AskForConfirmation", func(t *testing.T) {
		t.Run("panics", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewJSONUI(parentUI, NewRecordingLogger())

			assert.Panics(t, func() { ui.AskForConfirmation() })
		})
	})

	t.Run("IsInteractive", func(t *testing.T) {
		t.Run("delegates to the parent UI", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewJSONUI(parentUI, NewRecordingLogger())

			parentUI.Interactive = true
			assert.Equal(t, ui.IsInteractive(), true)

			parentUI.Interactive = false
			assert.Equal(t, ui.IsInteractive(), false)
		})
	})

	t.Run("Flush", func(t *testing.T) {
		t.Run("does not output anything when nothing was recorded", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewJSONUI(parentUI, NewRecordingLogger())

			ui.Flush()
			assert.Equal(t, len(parentUI.Said), 0)
		})

		t.Run("outputs everything when something was recorded", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewJSONUI(parentUI, NewRecordingLogger())

			ui.PrintLinef("fake-line1")
			ui.Flush()
			assert.Equal(t, parentUI.Blocks[0], `{
    "Tables": null,
    "Blocks": null,
    "Lines": [
        "fake-line1"
    ]
}`)
		})
	})
}
