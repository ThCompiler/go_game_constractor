package converter

import (
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/convert"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/functions"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/option"

	"strconv"
	"strings"
	"unicode"
)

type NumberToConvert interface {
	string |
		int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// Convert converter number into the words representation.
func Convert[N NumberToConvert](number N, options option.Options) (string, error) {
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
		if extractedNumber.Divider == constants.DECIMAL_NUMBER {
			extractedNumber.SecondPart = extractedNumber.SecondPart[0:constants.MaxNumberPartLength]
		} else {
			extractedNumber.SecondPart =
				extractedNumber.SecondPart[(len(extractedNumber.SecondPart) - constants.MaxNumberPartLength):]
		}
	}

	// Собрать конечный словесный результат
	convertedNumberString := ConvertByNumber(extractedNumber, options)

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

func ConvertByNumber(number objects.Number, appliedOptions option.Options) string {
	conv := convert.NewConverter(appliedOptions)

	convertedNumber := objects.ResultNumberT{}
	modifiedNumber := number

	// Если есть знак минус
	if number.Sign == "-" {
		// Если отображать знак минус словом
		if appliedOptions.ConvertMinusSignToWord {
			convertedNumber.Sign = appliedOptions.Language.GetMinusString()
		} else {
			convertedNumber.Sign = "-"
		}
	}

	// Если указана валюта
	if appliedOptions.Language.IsCurrency() {
		// Округлить число до 2 знаков после запятой
		modifiedNumber = functions.RoundNumber(modifiedNumber, constants.TwoSignAfterRoundForCurrency)
	} else {
		// Округлить число до заданной точности
		modifiedNumber = functions.RoundNumber(number, appliedOptions.RoundNumber)
	}

	// Если нужно отображать целую часть числа
	if appliedOptions.ShowNumberParts.Integer {
		convertedNumber = conv.ConvertIntegerPart(convertedNumber, modifiedNumber.FirstPart, modifiedNumber.Divider)
	}

	// Если нужно отображать дробную часть числа
	if appliedOptions.ShowNumberParts.Fractional {
		convertedNumber = conv.ConvertFractionalPart(convertedNumber, modifiedNumber.FirstPart,
			modifiedNumber.SecondPart, modifiedNumber.Divider)
	}

	// Объединить полученный результат
	convertedNumberResult := convertedNumber.Sign + " " +
		convertedNumber.FirstPart + " " + convertedNumber.FirstPartName + " " +
		convertedNumber.SecondPart + " " + convertedNumber.SecondPartName

	convertedNumberResult = strings.TrimSpace(functions.ReplaceInString(convertedNumberResult, `\s+`, ` `))

	if appliedOptions.AddUppercaseToFirstSymbol {
		// Сделать первую букву заглавной
		r := []rune(convertedNumberResult)
		return string(append([]rune{unicode.ToUpper(r[0])}, r[1:]...))
	}
	return convertedNumberResult
}
