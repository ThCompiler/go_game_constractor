package functions

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/objects"
	"regexp"
	"strings"
)

const (
	AllowCharRegex     = `[^ \d.,/+-]`
	AllowSignCharRegex = `[+-]`
	SignCharRegex      = `^[ ]*[+-]`
	DelimiterCharRegex = `[.,/]`
	ZerosInBeginRegex  = `^0+`
	ZerosInEndRegex    = `0+$`
)

func ValidateNumber(number string) error {
	if found := regexp.MustCompile(AllowCharRegex).FindAllString(number, -1); found != nil {
		return errorNumberContainsIncorrectChars(found)
	}

	if found := regexp.MustCompile(AllowSignCharRegex).FindAllString(number, -1); found != nil {
		if len(found) > 1 {
			return ErrorNumberHaveManySignChars
		}

		if found = regexp.MustCompile(SignCharRegex).FindAllString(number, -1); found == nil {
			return ErrorNumberHaveSignCharNotInBegin
		}
	}

	if found := regexp.MustCompile(DelimiterCharRegex).FindAllString(number, -1); len(found) > 1 {
		return errorNumberHaveManyDelimiterChars(found)
	}

	return nil
}

func ExtractNumber(Number string) (objects.Number, error) {
	number := objects.Number{
		Divider: constants.DECIMAL_NUMBER,
		Sign:    "+",
	}

	if err := ValidateNumber(Number); err != nil {
		return number, err
	}

	Number = RemoveFromString(Number, " ")
	if len(Number) < 1 {
		Number = "0"
	}

	if strings.IndexRune(Number, '-') == 0 {
		number.Sign = "-"
		Number = Number[1:]
	}

	if strings.IndexRune(Number, '+') == 0 {
		Number = Number[1:]
	}

	leftPart := ""
	rightPart := ""
	found := false //nolint:ifshort // it's not possibale to initialize found in if statement
	// Добавить разделитель числа в массив и разделить число
	if leftPart, rightPart, found = strings.Cut(Number, ","); found {
		number.Divider = constants.DECIMAL_NUMBER
	} else if leftPart, rightPart, found = strings.Cut(Number, "."); found {
		number.Divider = constants.DECIMAL_NUMBER
	} else if leftPart, rightPart, found = strings.Cut(Number, "/"); found {
		number.Divider = constants.FRACTIONAL_NUMBER
	}

	// Убрать лишние нули из целой части
	number.FirstPart = RemoveFromString(leftPart, ZerosInBeginRegex)

	number.SecondPart = rightPart
	// Убрать лишние нули из дробной части
	if number.Divider == constants.FRACTIONAL_NUMBER {
		number.SecondPart = RemoveFromString(rightPart, ZerosInBeginRegex)
	} else {
		number.SecondPart = RemoveFromString(rightPart, ZerosInEndRegex)
	}

	// Заменить пустые значения на ноль
	if number.SecondPart == "" {
		number.SecondPart = "0"
	}

	if number.FirstPart == "" {
		number.FirstPart = "0"
	}

	return number, nil
}

func RemoveFromString(str, regex string) string {
	return ReplaceInString(str, regex, ``)
}

func ReplaceInString(str, regex, repl string) string {
	if len(regex) == 0 {
		return str
	}
	return regexp.MustCompile(regex).ReplaceAllString(str, repl)
}
