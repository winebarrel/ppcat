package ppcat

import (
	"regexp"
)

var Matcher = map[*regexp.Regexp]func([]byte) ([]byte, error){}

func ParseLine(line []byte) []byte {
	for m, parse := range Matcher {
		if m.Match(line) {
			newLine, err := parse(line)

			if err != nil {
				continue
			}

			return newLine
		}
	}

	return line
}
