package ppcat_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/ppcat"
)

func TestParseJSON_JSON(t *testing.T) {
	assert := assert.New(t)
	line := `{"foo":"bar"}`
	out, err := ppcat.ParseJSON([]byte(line))
	assert.NoError(err)
	assert.Equal(`{
  "foo": "bar"
}`, string(out))
}

func TestParseJSON_BrokenJSON(t *testing.T) {
	assert := assert.New(t)
	line := `{"foo":"`
	_, err := ppcat.ParseJSON([]byte(line))
	assert.Error(err)
}
