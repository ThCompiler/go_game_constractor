package convertor

import (
	"strconv"

	"github.com/ThCompiler/go_game_constractor/pkg/convertor/num2words/functions"
)

// Convert convertor number into the words representation.
func Convert(number interface{}, options Options) (string, error) {
	numberString := convertNumberToString(number)
	if numberString == "" {
		return "", nil
	}

	// Обработать введенное число
	numberArray := functions.SplitNumberToArray(numberString)

	// Собрать конечный словесный результат
	convertedNumberString := combineResultData(numberArray, options)

	return convertedNumberString, nil
}

func convertNumberToString(number interface{}) string {
	switch convertNumber := number.(type) {
	case string:
		return convertNumber
	case int:
		return strconv.FormatInt(int64(convertNumber), 10)
	case int8:
		return strconv.FormatInt(int64(convertNumber), 10)
	case int16:
		return strconv.FormatInt(int64(convertNumber), 10)
	case int32:
		return strconv.FormatInt(int64(convertNumber), 10)
	case int64:
		return strconv.FormatInt(convertNumber, 10)
	case uint:
		return strconv.FormatUint(uint64(convertNumber), 10)
	case uint8:
		return strconv.FormatUint(uint64(convertNumber), 10)
	case uint16:
		return strconv.FormatUint(uint64(convertNumber), 10)
	case uint32:
		return strconv.FormatUint(uint64(convertNumber), 10)
	case uint64:
		return strconv.FormatUint(convertNumber, 10)
	case float32:
		return strconv.FormatFloat(float64(convertNumber), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(convertNumber, 'f', -1, 64)
	default:
		return ""
	}
}
