package ui_test

import (
	"fmt"
	"testing"

	. "github.com/cppforlife/go-cli-ui/ui"
	fakeui "github.com/cppforlife/go-cli-ui/ui/fakes"
	. "github.com/cppforlife/go-cli-ui/ui/table"
	"github.com/stretchr/testify/assert"
)

func TestNonInteractiveUI(t *testing.T) {
	t.Run("ErrorLinef", func(t *testing.T) {
		t.Run("delegates to the parent UI", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonInteractiveUI(parentUI)

			ui.ErrorLinef("fake-error-line")
			assert.Equal(t, parentUI.Errors, []string{"fake-error-line"})
		})
	})

	t.Run("PrintLinef", func(t *testing.T) {
		t.Run("delegates to the parent UI", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonInteractiveUI(parentUI)

			ui.PrintLinef("fake-line")
			assert.Equal(t, parentUI.Said, []string{"fake-line"})
		})
	})

	t.Run("BeginLinef", func(t *testing.T) {
		t.Run("delegates to the parent UI", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonInteractiveUI(parentUI)

			ui.BeginLinef("fake-start")
			assert.Equal(t, parentUI.Said, []string{"fake-start"})
		})
	})

	t.Run("EndLinef", func(t *testing.T) {
		t.Run("delegates to the parent UI", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonInteractiveUI(parentUI)

			ui.EndLinef("fake-end")
			assert.Equal(t, parentUI.Said, []string{"fake-end"})
		})
	})

	t.Run("PrintBlock", func(t *testing.T) {
		t.Run("delegates to the parent UI", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonInteractiveUI(parentUI)

			ui.PrintBlock([]byte("block"))
			assert.Equal(t, parentUI.Blocks, []string{"block"})
		})
	})

	t.Run("PrintErrorBlock", func(t *testing.T) {
		t.Run("delegates to the parent UI", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonInteractiveUI(parentUI)

			ui.PrintErrorBlock("block")
			assert.Equal(t, parentUI.Blocks, []string{"block"})
		})
	})

	t.Run("PrintTable", func(t *testing.T) {
		t.Run("delegates to the parent UI", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonInteractiveUI(parentUI)

			table := Table{
				Content: "things",
				Header:  []Header{NewHeader("header1")},
			}

			ui.PrintTable(table)

			assert.Equal(t, parentUI.Table, table)
		})
	})

	t.Run("AskForText", func(t *testing.T) {
		t.Run("default non empty", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonInteractiveUI(parentUI)

			text, err := ui.AskForText(TextOpts{
				Label:   "",
				Default: "foo",
				ValidateFunc: func(s string) (bool, string, error) {
					if s == "" {
						return false, "", fmt.Errorf("should not be empty")
					}
					return true, "", nil
				},
			})

			assert.Equal(t, text, "foo")
			assert.Nil(t, err)
		})
		t.Run("default empty", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonInteractiveUI(parentUI)

			text, err := ui.AskForText(TextOpts{
				Label:   "",
				Default: "",
				ValidateFunc: func(s string) (bool, string, error) {
					if s == "" {
						return false, "", fmt.Errorf("should not be empty")
					}
					return true, "", nil
				},
			})

			assert.Equal(t, text, "")
			assert.NotNil(t, err)
		})
	})

	t.Run("AskForPassword", func(t *testing.T) {
		t.Run("panics", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonInteractiveUI(parentUI)

			assert.Panics(t, func() { ui.AskForPassword("") })
		})
	})

	t.Run("AskForChoice", func(t *testing.T) {
		t.Run("non negative default value", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonInteractiveUI(parentUI)

			choice, err := ui.AskForChoice(ChoiceOpts{
				Label:   "",
				Default: 1,
				Choices: []string{"a", "b", "c"},
			})

			assert.Equal(t, choice, 1)
			assert.Nil(t, err)
		})
	})

	t.Run("AskForConfirmation", func(t *testing.T) {
		t.Run("responds affirmatively with no error", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonInteractiveUI(parentUI)

			assert.Equal(t, ui.AskForConfirmation(), nil)
		})
	})

	t.Run("IsInteractive", func(t *testing.T) {
		t.Run("returns false", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonInteractiveUI(parentUI)

			assert.Equal(t, ui.IsInteractive(), false)
		})
	})

	t.Run("Flush", func(t *testing.T) {
		t.Run("delegates to the parent UI", func(t *testing.T) {
			parentUI := &fakeui.FakeUI{}
			ui := NewNonInteractiveUI(parentUI)

			ui.Flush()
			assert.Equal(t, parentUI.Flushed, true)
		})
	})
}
