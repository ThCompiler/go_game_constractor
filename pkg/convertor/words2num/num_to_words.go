package words2num

/*
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
		res = AddConvertedWordToNumber(word, res)
	}

	return res, nil
}

func AddConvertedWordToNumber(word string, number int64) int64 {
	if value, is := ru.WordConstants.W2n.Digit.Units[word]; is {
		return number + int64(value)
	}

	if value, is := words2.WordConstants.W2n.Digit.Tens[word]; is {
		return number + ten + int64(value)
	}

	if value, is := words2.WordConstants.W2n.Digit.Dozens[word]; is {
		return number + int64(value)*ten
	}

	if value, is := words2.WordConstants.W2n.Digit.Hundreds[word]; is {
		return number + int64(value)*hundred
	}

	if scale, is := words2.WordConstants.W2n.UnitScalesNamesToNumber.Words[word]; is {
		return number * int64(math.Pow(thousand, float64(scale)))
	}

	return number
}
*/
