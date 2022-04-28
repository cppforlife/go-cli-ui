package table_test

import (
	"testing"

	"github.com/cppforlife/go-cli-ui/ui/table"
	"github.com/stretchr/testify/assert"
)

func TestKeyifyHeader(t *testing.T) {
	t.Run("should convert alphanumeric to lowercase ", func(t *testing.T) {
		keyifyHeader := table.KeyifyHeader("Header1")
		assert.Equal(t, keyifyHeader, "header1")
	})

	t.Run("given a header that only contains non-alphanumeric and alphanumeric should non-alphanumeric to underscore", func(t *testing.T) {
		keyifyHeader := table.KeyifyHeader("FOO!@AND#$BAR")
		assert.Equal(t, keyifyHeader, "foo_and_bar")
	})

	t.Run("given a header that only contains non-alphanumeric", func(t *testing.T) {
		t.Run("should convert to underscore", func(t *testing.T) {
			keyifyHeader := table.KeyifyHeader("!@#$")
			assert.Equal(t, keyifyHeader, "_")
		})

		t.Run("should convert empty header to underscore", func(t *testing.T) {
			keyifyHeader := table.KeyifyHeader("")
			assert.Equal(t, keyifyHeader, "_")
		})
	})
}

func TestSetColumnVisibility(t *testing.T) {
	t.Run("when given a header that does not exist should return an error", func(t *testing.T) {
		tbl := table.Table{
			Header: []table.Header{table.NewHeader("header1")},
		}
		err := tbl.SetColumnVisibility([]table.Header{table.NewHeader("non-matching-header")})
		assert.Error(t, err)
	})
}
