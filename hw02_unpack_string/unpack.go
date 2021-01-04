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

		if (unicode.IsDigit(b) && i == 0) ||
			(unicode.IsDigit(b) && (i < lenArr && unicode.IsDigit(a[i+1]))) {
			return "", ErrInvalidString
		} else if b == '0' {
			// если 0 от убираем  последний добавленный символ
			buff.Truncate(buff.Len() - 1)
		} else if unicode.IsDigit(b) {
			// гарда от 0 выше
			buff.WriteString(strings.Repeat(
				string(a[i-1]),
				int(b-'0')-1),
			)
		} else {
			buff.WriteRune(b)
		}
	}
	return buff.String(), nil
}
