package hw02_unpack_string //nolint:golint,stylecheck

import (
	"bytes"
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var a = []rune(s)
	var lenArr = len(a) - 1
	var buff bytes.Buffer

	for i, b := range a {
		switch {
		case (unicode.IsDigit(b) && i == 0) ||
			(unicode.IsDigit(b) && (i < lenArr && unicode.IsDigit(a[i+1]))):
			return "", ErrInvalidString
		case b == '0':
			buff.Truncate(buff.Len() - 1)
		case unicode.IsDigit(b):
			buff.WriteString(strings.Repeat(
				string(a[i-1]),
				int(b-'0')-1),
			)
		default:
			buff.WriteRune(b)
		}
	}
	return buff.String(), nil
}
