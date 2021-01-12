package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "test1", input: "a4bc2d5e", expected: "aaaabccddddde"},
		{name: "test2", input: "abccd", expected: "abccd"},
		{name: "test3", input: "", expected: ""},
		{name: "test4", input: "aaa0b", expected: "aab"},
		{name: "test-", input: "aa-5b", expected: "aa-----b"},
		{name: "test-", input: "ab5", expected: "abbbbb"},
		{name: "test-emoji", input: "a-😡5", expected: "a-😡😡😡😡😡"},
		{name: "test--", input: "-", expected: "-"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", "0"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
