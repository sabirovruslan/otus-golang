package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var builder strings.Builder
	prev := ""
	for _, s := range input {
		number, err := strconv.Atoi(string(s))

		if len(prev) == 0 && err == nil {
			return "", ErrInvalidString
		}

		if number > 0 {
			builder.WriteString(strings.Repeat(prev, number))
			prev = ""
			continue
		}

		if number == 0 && err == nil {
			prev = ""
			continue
		}

		builder.WriteString(prev)
		prev = string(s)
	}

	builder.WriteString(prev)

	return builder.String(), nil
}
