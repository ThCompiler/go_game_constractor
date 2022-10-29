package convertor

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/num2words/functions"
	"strconv"
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
	switch number.(type) {
	case string:
		return number.(string)
	case int:
		return strconv.FormatInt(int64(number.(int)), 10)
	case int8:
		return strconv.FormatInt(int64(number.(int8)), 10)
	case int16:
		return strconv.FormatInt(int64(number.(int16)), 10)
	case int32:
		return strconv.FormatInt(int64(number.(int32)), 10)
	case int64:
		return strconv.FormatInt(number.(int64), 10)
	case uint:
		return strconv.FormatUint(uint64(number.(uint)), 10)
	case uint8:
		return strconv.FormatUint(uint64(number.(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(number.(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(number.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(number.(uint64), 10)
	case float32:
		return strconv.FormatFloat(float64(number.(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(number.(float64), 'f', -1, 64)
	default:
		return ""
	}
}
