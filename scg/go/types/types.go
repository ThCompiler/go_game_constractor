package types

import "strings"

const types = "bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string, bytes"

func IsValidType(s string) bool {
	if strings.Contains(types, s) {
		return true
	} else {
		return false
	}
}

func GetSupportTypes() string {
	return types
}
