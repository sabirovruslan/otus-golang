package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	prev := ""
	result := ""
	for _, s := range input {
		number, err := strconv.Atoi(string(s))

		if len(prev) <= 0 && err == nil {
			return "", ErrInvalidString
		}

		if number > 0 {
			result += strings.Repeat(prev, int(s-'0')-1)
			prev = ""
			continue
		}

		if number == 0 && err == nil {
			result = trimLastChar(result)
			prev = ""
			continue
		}

		prev = string(s)
		result += prev
	}

	return result, nil
}

func trimLastChar(s string) string {
	if len(s) > 0 {
		_, size := utf8.DecodeLastRuneInString(s)

		return s[:len(s)-size]
	}

	return s
}
