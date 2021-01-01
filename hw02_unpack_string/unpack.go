package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
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
			builder.WriteString(strings.Repeat(prev, int(s-'0')-1))
			prev = ""
			continue
		}

		if number == 0 && err == nil {
			trimLastChar(&builder)
			prev = ""
			continue
		}

		prev = string(s)
		builder.WriteRune(s)
	}

	return builder.String(), nil
}

func trimLastChar(b *strings.Builder) {
	s := b.String()
	if len(s) > 0 {
		_, size := utf8.DecodeLastRuneInString(s)
		b.Reset()
		b.WriteString(s[:len(s)-size])
	}
}
