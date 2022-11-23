package words2num

import (
	words2 "github.com/ThCompiler/go_game_constractor/pkg/convertor/words"
	"math"
	"strings"
)

// Convert convertor number into the words representation.
func Convert(str string) (int64, error) {
	if str == "" {
		return -1, nil
	}

	return convert(str)
}

const (
	ten      = 10
	hundred  = 100
	thousand = 1000
)

func convert(str string) (int64, error) {
	words := strings.Fields(str)
	res := int64(0)

	for _, word := range words {
		if value, is := words2.WordConstants.W2n.Digit.Units[word]; is {
			res += int64(value)
		} else if value, is = words2.WordConstants.W2n.Digit.Tens[word]; is {
			res += ten + int64(value)
		} else if value, is = words2.WordConstants.W2n.Digit.Dozens[word]; is {
			res += int64(value) * ten
		} else if value, is = words2.WordConstants.W2n.Digit.Hundreds[word]; is {
			res += int64(value) * hundred
		} else if scale, is := words2.WordConstants.W2n.UnitScalesNamesToNumber.Words[word]; is {
			res *= int64(math.Pow(thousand, float64(scale)))
		}
	}

	return res, nil
}
