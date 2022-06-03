package table_test

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
	"time"

	. "github.com/cppforlife/go-cli-ui/ui/table"
	"github.com/stretchr/testify/assert"
)

func TestValueString(t *testing.T) {
	t.Run("returns string", func(t *testing.T) {
		assert.Equal(t, ValueString{S: "val"}.String(), "val")
	})

	t.Run("returns itself", func(t *testing.T) {
		assert.Equal(t, ValueString{S: "val"}.Value(), ValueString{S: "val"})
	})

	t.Run("returns int based on string compare", func(t *testing.T) {
		assert.Equal(t, ValueString{S: "a"}.Compare(ValueString{S: "a"}), 0)
		assert.Equal(t, ValueString{S: "a"}.Compare(ValueString{S: "b"}), -1)
		assert.Equal(t, ValueString{S: "b"}.Compare(ValueString{S: "a"}), 1)
	})
}

func TestValueStrings(t *testing.T) {
	t.Run("returns new line joined strings", func(t *testing.T) {
		assert.Equal(t, ValueStrings{S: []string{"val1", "val2"}}.String(), "val1\nval2")
	})

	t.Run("returns itself", func(t *testing.T) {
		assert.Equal(t, ValueStrings{S: []string{"val1"}}.Value(), ValueStrings{S: []string{"val1"}})
	})

	t.Run("returns int based on string compare", func(t *testing.T) {
		assert.Equal(t, ValueStrings{S: []string{"val1"}}.Compare(ValueStrings{S: []string{"val1"}}), 0)
		assert.Equal(t, ValueStrings{S: []string{"val1"}}.Compare(ValueStrings{S: []string{"val1", "val2"}}), -1)
		assert.Equal(t, ValueStrings{S: []string{"val1", "val2"}}.Compare(ValueStrings{S: []string{"val1"}}), 1)
	})
}

func TestValueInt(t *testing.T) {
	t.Run("returns string", func(t *testing.T) {
		assert.Equal(t, ValueInt{I: 1}.String(), "1")
	})

	t.Run("returns itself", func(t *testing.T) {
		assert.Equal(t, ValueInt{I: 1}.Value(), ValueInt{I: 1})
	})

	t.Run("returns int based on int compare", func(t *testing.T) {
		assert.Equal(t, ValueInt{I: 1}.Compare(ValueInt{I: 1}), 0)
		assert.Equal(t, ValueInt{I: 1}.Compare(ValueInt{I: 2}), -1)
		assert.Equal(t, ValueInt{I: 2}.Compare(ValueInt{I: 1}), 1)
	})
}

func TestValueTime(t *testing.T) {
	t1 := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	t2 := time.Date(2009, time.November, 10, 23, 0, 0, 1, time.UTC)
	empty := time.Time{}

	t.Run("returns formatted full time", func(t *testing.T) {
		assert.Equal(t, ValueTime{T: t1}.String(), "2009-11-10T23:00:00Z")
	})

	t.Run("returns empty", func(t *testing.T) {
		assert.Equal(t, ValueTime{T: empty}.String(), "")
	})

	t.Run("returns itself", func(t *testing.T) {
		assert.Equal(t, ValueTime{T: t1}.Value(), ValueTime{T: t1})
	})

	t.Run("returns int based on time compare", func(t *testing.T) {
		assert.Equal(t, ValueTime{T: t1}.Compare(ValueTime{T: t1}), 0)
		assert.Equal(t, ValueTime{T: t1}.Compare(ValueTime{T: t2}), -1)
		assert.Equal(t, ValueTime{T: t2}.Compare(ValueTime{T: t1}), 1)
	})
}

func TestValueBool(t *testing.T) {
	t.Run("returns true/false as string", func(t *testing.T) {
		assert.Equal(t, ValueBool{B: true}.String(), "true")
		assert.Equal(t, ValueBool{B: false}.String(), "false")
	})

	t.Run("returns itself", func(t *testing.T) {
		assert.Equal(t, ValueBool{B: true}.Value(), ValueBool{B: true})
	})

	t.Run("returns int based on bool compare", func(t *testing.T) {
		assert.Equal(t, ValueBool{B: true}.Compare(ValueBool{B: true}), 0)
		assert.Equal(t, ValueBool{B: false}.Compare(ValueBool{B: true}), -1)
		assert.Equal(t, ValueBool{B: true}.Compare(ValueBool{B: false}), 1)
	})
}

func TestValueError(t *testing.T) {
	t.Run("returns empty string or error description", func(t *testing.T) {
		assert.Equal(t, ValueError{}.String(), "")
		assert.Equal(t, ValueError{E: errors.New("err")}.String(), "err")
	})

	t.Run("returns itself", func(t *testing.T) {
		assert.Equal(t, ValueError{E: errors.New("err")}.Value(), ValueError{E: errors.New("err")})
	})

	t.Run("does not allow comparison", func(t *testing.T) {
		f := func() { ValueError{}.Compare(ValueError{}) }
		assert.Panics(t, f)
	})
}

func TestValueNone(t *testing.T) {
	t.Run("returns empty string", func(t *testing.T) {
		assert.Equal(t, ValueNone{}.String(), "")
	})

	t.Run("returns itself", func(t *testing.T) {
		assert.Equal(t, ValueNone{}.Value(), ValueNone{})
	})

	t.Run("does not allow comparison", func(t *testing.T) {
		f := func() { ValueNone{}.Compare(ValueNone{}) }
		assert.Panics(t, f)
	})
}

func TestValueFmt(t *testing.T) {
	fmtFunc := func(pattern string, vals ...interface{}) string {
		return fmt.Sprintf(">%s<", fmt.Sprintf(pattern, vals...))
	}

	t.Run("returns plain string (not formatted with fmt func)", func(t *testing.T) {
		assert.Equal(t, ValueFmt{V: ValueInt{I: 1}, Func: fmtFunc}.String(), "1")
	})

	t.Run("returns wrapped value", func(t *testing.T) {
		assert.Equal(t, ValueFmt{V: ValueInt{I: 1}, Func: fmtFunc}.Value(), ValueInt{I: 1})
	})

	t.Run("does not allow comparison", func(t *testing.T) {
		f := func() { ValueFmt{V: ValueInt{I: 1}, Func: fmtFunc}.Compare(ValueFmt{}) }
		assert.Panics(t, f)
	})

	t.Run("writes out value using custom Fprintf", func(t *testing.T) {
		buf := bytes.NewBufferString("")
		ValueFmt{V: ValueInt{I: 1}, Func: fmtFunc}.Fprintf(buf, "%s,%s", "val1", "val2")
		assert.Equal(t, buf.String(), ">val1,val2<")
	})

	t.Run("uses fmt.Fprintf if fmt func is not set", func(t *testing.T) {
		buf := bytes.NewBufferString("")
		ValueFmt{V: ValueInt{I: 1}}.Fprintf(buf, "%s,%s", "val1", "val2")
		assert.Equal(t, buf.String(), "val1,val2")
	})
}

type failsToYAMLMarshal struct{}

func (s failsToYAMLMarshal) MarshalYAML() (interface{}, error) {
	return nil, errors.New("marshal-err")
}

func TestValueInterface(t *testing.T) {
	t.Run("returns map as a string", func(t *testing.T) {
		i := map[string]interface{}{"key": "value", "num": 123}
		assert.Equal(t, ValueInterface{I: i}.String(), "key: value\nnum: 123")
	})

	t.Run("returns nested items as a string", func(t *testing.T) {
		i := map[string]interface{}{"key": map[string]interface{}{"nested_key": "nested_value"}}
		assert.Equal(t, ValueInterface{I: i}.String(), "key:\n  nested_key: nested_value")
	})

	// Tests contract specified here: https://github.com/go-yaml/yaml/blob/v3/yaml.go#L44-L52
	t.Run("respects Marshaller interface", func(t *testing.T) {
		i := failsToYAMLMarshal{}
		assert.Equal(t, ValueInterface{I: i}.String(), `<serialization error> : table_test.failsToYAMLMarshal{}`)
	})

	t.Run("returns error on failure", func(t *testing.T) {
		i := map[string]interface{}{"foo": make(chan int)}
		assert.Contains(t, ValueInterface{I: i}.String(), `<serialization error> :`)
	})

	t.Run("returns nil items as blank string", func(t *testing.T) {
		assert.Equal(t, ValueInterface{I: nil}.String(), "")
	})

	t.Run("returns an empty map as blank string", func(t *testing.T) {
		i := map[string]interface{}{}
		assert.Equal(t, ValueInterface{I: i}.String(), "")
	})

	t.Run("returns an empty slice as blank string", func(t *testing.T) {
		i := []string{}
		assert.Equal(t, ValueInterface{I: i}.String(), "")
	})
}

func TestValueSuffix(t *testing.T) {
	t.Run("returns formatted string with suffix", func(t *testing.T) {
		assert.Equal(t, ValueSuffix{V: ValueInt{I: 1}, Suffix: "*"}.String(), "1*")
		assert.Equal(t, ValueSuffix{V: ValueString{S: "val"}, Suffix: "*"}.String(), "val*")
	})

	t.Run("returns wrapped value", func(t *testing.T) {
		assert.Equal(t, ValueSuffix{V: ValueInt{I: 1}, Suffix: "*"}.Value(), ValueInt{I: 1})
	})

	t.Run("does not allow comparison", func(t *testing.T) {
		f := func() { ValueSuffix{V: ValueInt{I: 1}, Suffix: ""}.Compare(ValueSuffix{}) }
		assert.Panics(t, f)
	})
}
