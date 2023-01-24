package stringutilits

import "unicode"

func ToDigit(r rune) int8 {
	if unicode.IsDigit(r) {
		return int8(r - '0')
	}

	return -1
}

func ToRune(n int8) rune {
	if n < 0 || n > 9 {
		return rune(0)
	}

	return rune('0' + n)
}
