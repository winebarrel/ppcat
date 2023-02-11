package ppcat_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/ppcat"
)

func TestParseLine_JSON(t *testing.T) {
	assert := assert.New(t)
	line := `{"foo":"bar"}`
	out := ppcat.ParseLine([]byte(line))
	assert.Equal(`{
  "foo": "bar"
}`, string(out))
}

func TestParseLine_BrokenJSON(t *testing.T) {
	assert := assert.New(t)
	line := `{"foo":"`
	out := ppcat.ParseLine([]byte(line))
	assert.Equal(`{"foo":"`, string(out))
}

func TestParseLine_PlainText(t *testing.T) {
	assert := assert.New(t)
	line := `London Bridge is broken down`
	out := ppcat.ParseLine([]byte(line))
	assert.Equal(`London Bridge is broken down`, string(out))
}
