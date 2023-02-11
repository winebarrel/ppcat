package ppcat

import (
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

	return json.MarshalIndent(parsedLine, "", "  ")
}
