package functions

import (
	"regexp"
	"strings"

	"github.com/ThCompiler/go_game_constractor/pkg/convertor/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/num2words/objects"
)

func SplitNumberToArray(startNumber string) objects.Number {
	number := objects.Number{
		Divider: constants.DecimalNumber,
		Sign:    "+",
	}

	// Убрать из строки всё лишнее
	cleanNumber := ClearFromString(startNumber, `[^\d\.\,\/\-]`)
	if len(cleanNumber) < 1 {
		cleanNumber = "0"
	}

	if strings.ContainsRune(cleanNumber, '-') {
		number.Sign = "-"
	}
	// Удалить все знаки минуса
	cleanNumber = ClearFromString(cleanNumber, `[\-]`)

	leftPart := ""
	rightPart := ""
	found := false //nolint:ifshort // it's not possibale to initialize found in if statement
	// Добавить разделитель числа в массив и разделить число
	if leftPart, rightPart, found = strings.Cut(cleanNumber, ","); found {
		number.Divider = constants.DecimalNumber
	} else if leftPart, rightPart, found = strings.Cut(cleanNumber, "."); found {
		number.Divider = constants.DecimalNumber
	} else if leftPart, rightPart, found = strings.Cut(cleanNumber, "/"); found {
		number.Divider = constants.FractionalNumber
	}

	// Удалить все разделители числа
	leftPart = ClearFromString(leftPart, `[\, \.\/]`)
	rightPart = ClearFromString(rightPart, `[\, \.\/]`)

	// Убрать лишние нули из целой части
	number.FirstPart = ClearFromString(leftPart, `^0+/`)

	number.SecondPart = rightPart
	// Убрать лишние нули из дробной части
	if number.Divider == constants.DecimalNumber {
		number.SecondPart = ClearFromString(rightPart, `^0+/`)
	}

	// Заменить пустые значения на ноль
	if number.SecondPart == "" {
		number.SecondPart = "0"
	}

	if number.FirstPart == "" {
		number.FirstPart = "0"
	}

	if len(number.FirstPart) > constants.MaxIntegerPartLength {
		// Убрать лишнюю целую часть числа
		number.FirstPart = number.FirstPart[0:constants.MaxIntegerPartLength]
	}

	if len(number.SecondPart) > constants.MaxIntegerPartLength {
		// Убрать лишнюю десятичную часть числа
		number.SecondPart = number.SecondPart[0:constants.MaxIntegerPartLength]
	}

	return number
}

func ClearFromString(str, regex string) string {
	return ReplaceInString(str, regex, ``)
}

func ReplaceInString(str, regex, repl string) string {
	return regexp.MustCompile(regex).ReplaceAllString(str, repl)
}
