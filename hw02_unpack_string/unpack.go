package hw02_unpack_string //nolint:golint,stylecheck,gomnd

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var runeBuf rune
	var output string

	for _, char := range s {
		if !unicode.IsDigit(char) {
			output += string(char)
			if char == runeBuf {
				// repeated characters
				return "", ErrInvalidString
			}
			runeBuf = char
		} else {
			count, _ := strconv.Atoi(string(char))
			output += strings.Repeat(string(runeBuf), count-1)
			if runeBuf == 0 || unicode.IsDigit(runeBuf) {
				// sequential numbers or string starting with number
				return "", ErrInvalidString
			}
			runeBuf = char // for sequential numbers check
		}
	}
	return output, nil
}
