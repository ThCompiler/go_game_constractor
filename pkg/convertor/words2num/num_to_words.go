package words2num

// TODO
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

func convert(str string) (int64, error) {
	words := strings.Fields(str)
	res := int64(0)
	for _, word := range words {
		if value, is := words2.WordConstants.W2n.WordsDigit.Units[word]; is {
			res += int64(value)
		} else if value, is = words2.WordConstants.W2n.WordsDigit.Tens[word]; is {
			res += 10 + int64(value)
		} else if value, is = words2.WordConstants.W2n.WordsDigit.Dozens[word]; is {
			res += int64(value) * 10
		} else if value, is = words2.WordConstants.W2n.WordsDigit.Hundreds[word]; is {
			res += int64(value) * 100
		} else if scale, is := words2.WordConstants.W2n.UnitScalesNamesToNumber.Words[word]; is {
			res *= int64(math.Pow(1000, float64(scale)))
		}
	}
	return res, nil
}
