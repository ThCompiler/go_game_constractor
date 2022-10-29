package stringutilits

import "unicode"

func ToDigit(r rune) int8 {
	if unicode.IsDigit(r) {
		return int8(r - '0')
	} else {
		return -1
	}
}

func ToRune(n int8) rune {
	return rune('0' + n)
}
