package convertor

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/constants"
	"strconv"

	"github.com/ThCompiler/go_game_constractor/pkg/convertor/num2words/functions"
)

type ConvertableToString interface {
	ToString() string
}

type NumberToConvert interface {
	ConvertableToString
	string |
		int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// Convert convertor number into the words representation.
func Convert[N NumberToConvert](number N, options Options) (string, error) {
	numberString := convertNumberToString(number)
	if numberString == "" {
		return "", nil
	}

	// Обработать введенное число
	extractedNumber, err := functions.ExtractNumber(numberString)
	if err != nil {
		return "", err
	}

	if len(extractedNumber.FirstPart) > constants.MaxNumberPartLength {
		extractedNumber.FirstPart =
			extractedNumber.FirstPart[(len(extractedNumber.FirstPart) - constants.MaxNumberPartLength):]
	}

	if len(extractedNumber.SecondPart) > constants.MaxNumberPartLength {
		if extractedNumber.Divider == constants.DecimalNumber {
			extractedNumber.SecondPart = extractedNumber.SecondPart[0:constants.MaxNumberPartLength]
		} else {
			extractedNumber.SecondPart =
				extractedNumber.SecondPart[(len(extractedNumber.SecondPart) - constants.MaxNumberPartLength):]
		}
	}

	// Собрать конечный словесный результат
	convertedNumberString := combineResultData(extractedNumber, options)

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
	case ConvertableToString:
		return convertNumber.ToString()
	default:
		return ""
	}
}
