package hw02_unpack_string //nolint:golint,stylecheck

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	input    string
	expected string
	err      error
}

func TestUnpack(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			input:    "abccd",
			expected: "abccd",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "aaa0b",
			expected: "aab",
		},
		{
			input:    "a0a0a0",
			expected: "",
		},
		{
			input:    "AAa0b",
			expected: "AAb",
		},
		{
			input:    "AAa0B4",
			expected: "AABBBB",
		},
		{
			input:    "hh*3m0",
			expected: "hh***",
		},
		{
			input:    "|",
			expected: "|",
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}

func TestUnpackError(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    "3abc",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "45",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "aaa10b",
			expected: "",
			err:      ErrInvalidString,
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}

func TestTrimLastChar(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    "bba",
			expected: "bb",
		},
		{
			input:    "aaaa",
			expected: "aaa",
		},
		{
			input:    "Ab*",
			expected: "Ab",
		},
		{
			input:    "4324",
			expected: "432",
		},
		{
			input:    "1",
			expected: "",
		},
		{
			input:    "",
			expected: "",
		},
	} {
		var builder strings.Builder
		builder.WriteString(tst.input)

		trimLastChar(&builder)
		require.Equal(t, tst.expected, builder.String())
	}
}

func TestUnpackWithEscape(t *testing.T) {
	t.Skip() // NeedRemove if task with asterisk completed

	for _, tst := range [...]test{
		{
			input:    `qwe\4\5`,
			expected: `qwe45`,
		},
		{
			input:    `qwe\45`,
			expected: `qwe44444`,
		},
		{
			input:    `qwe\\5`,
			expected: `qwe\\\\\`,
		},
		{
			input:    `qwe\\\3`,
			expected: `qwe\3`,
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}
