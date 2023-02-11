package ppcat

import (
	"bufio"
	"fmt"
	"io"
)

var (
	ReadLineBufSize = 4096
)

func Cat(in io.Reader, out io.Writer) error {
	reader := bufio.NewReader(in)

	for {
		line, err := readLine(reader)

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		line = ParseLine(line)
		fmt.Fprintln(out, string(line))
	}

	return nil
}

func readLine(reader *bufio.Reader) ([]byte, error) {
	buf := make([]byte, 0, ReadLineBufSize)
	var err error

	for {
		line, isPrefix, e := reader.ReadLine()
		err = e

		if len(line) > 0 {
			buf = append(buf, line...)
		}

		if !isPrefix || err != nil {
			break
		}
	}

	return buf, err
}
