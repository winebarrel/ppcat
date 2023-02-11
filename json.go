package ppcat

import (
	"bytes"
	"encoding/json"
	"regexp"
)

func init() {
	re := regexp.MustCompile(`^\s*[[{]`)
	Matcher[re] = ParseJSON
}

func ParseJSON(line []byte) ([]byte, error) {
	var parsedLine any
	err := json.Unmarshal(line, &parsedLine)

	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	err = enc.Encode(parsedLine)

	if err != nil {
		return nil, err
	}

	pp := buf.Bytes()

	if pp[len(pp)-1] == '\n' {
		pp = pp[:len(pp)-1]
	}

	return pp, nil
}
