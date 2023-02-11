package ppcat_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/ppcat"
)

func TestCat_ParseLineWithJSON(t *testing.T) {
	assert := assert.New(t)

	lines := `London Bridge is broken down,
{"foo":"bar","zoo":[100,200]}
Broken down, broken down.
[{"foo":"bar","zoo":[100,200]},{"zoo":{"baz":[300,400]}}]
London Bridge is broken down,
{"foo": // broken
My fair lady.
          {"foo":"bar"}
`

	reader := strings.NewReader(lines)
	var buf bytes.Buffer
	err := ppcat.Cat(reader, &buf)
	assert.NoError(err)
	assert.Equal(`London Bridge is broken down,
{
  "foo": "bar",
  "zoo": [
    100,
    200
  ]
}
Broken down, broken down.
[
  {
    "foo": "bar",
    "zoo": [
      100,
      200
    ]
  },
  {
    "zoo": {
      "baz": [
        300,
        400
      ]
    }
  }
]
London Bridge is broken down,
{"foo": // broken
My fair lady.
{
  "foo": "bar"
}
`, buf.String())
}

func TestCat_ParseLineWithoutJSON(t *testing.T) {
	assert := assert.New(t)
	assert.True(true)

	lines := `London Bridge is broken down,
	Broken down, broken down.
	London Bridge is broken down,
	My fair lady.
`

	reader := strings.NewReader(lines)
	var buf bytes.Buffer
	err := ppcat.Cat(reader, &buf)
	assert.NoError(err)
	assert.Equal(lines, buf.String())
}
