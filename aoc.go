package aoc

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// ReadCommaSeparatedInts returns a slice with ints
func ReadCommaSeparatedInts(reader io.Reader) ([]int, error) {
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}

	values := make([]int, 0)
	scanner := bufio.NewScanner(reader)
	scanner.Split(onComma)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}
		value, err := strconv.Atoi(text)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}
