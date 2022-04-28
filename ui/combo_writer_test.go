package ui_test

import (
	"bytes"
	"fmt"
	"testing"

	. "github.com/cppforlife/go-cli-ui/ui"
	"github.com/stretchr/testify/assert"
)

func TestComboWriter(t *testing.T) {
	type Example struct {
		Ins []string
		Out string
	}

	examples := []Example{
		{Ins: []string{""}, Out: ""},
		{Ins: []string{"", ""}, Out: ""},
		{Ins: []string{"\n"}, Out: "prefix: \n"},
		{Ins: []string{"", "\n"}, Out: "prefix: \n"},
		{Ins: []string{"\n\n", "\n"}, Out: "prefix: \nprefix: \nprefix: \n"},
		{Ins: []string{"piece1"}, Out: "prefix: piece1"},
		{Ins: []string{"piece1", "piece2"}, Out: "prefix: piece1piece2"},
		{Ins: []string{"piece1", "piece2\n"}, Out: "prefix: piece1piece2\n"},
		{Ins: []string{"\npiece1", "piece2"}, Out: "prefix: \nprefix: piece1piece2"},
		{Ins: []string{"piece1", "\npiece2"}, Out: "prefix: piece1\nprefix: piece2"},
		{Ins: []string{"piece1\n", "piece2"}, Out: "prefix: piece1\nprefix: piece2"},
	}

	for i, ex := range examples {
		ex := ex

		t.Run(fmt.Sprintf("prints correctly '%d'", i), func(t *testing.T) {
			outBuffer := bytes.NewBufferString("")
			errBuffer := bytes.NewBufferString("")
			logger := NewRecordingLogger()
			ui := NewWriterUI(outBuffer, errBuffer, logger)
			w := NewComboWriter(ui).Writer("prefix: ")

			for _, in := range ex.Ins {
				w.Write([]byte(in))
			}
			assert.Equal(t, outBuffer.String(), ex.Out)
		})
	}
}
