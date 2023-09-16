package functions

import (
	"strings"
	"unicode"

	"github.com/ThCompiler/go_game_constractor/pkg/convertor/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/num2words/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/stringutilits"
)

const (
	nine = 9
	four = 4
)

func roundDigit(increaseDigit bool, numberToRound []rune, i int) (bool, []rune) {
	currentDigit := stringutilits.ToDigit(numberToRound[i])
	// Если нужно было увеличивать цифру
	if increaseDigit {
		if currentDigit != nine { // Если любая другая цифра
			numberToRound[i]++

			return false, numberToRound
		}
		// Если текущая цифра 9, то увеличить следующую
		numberToRound[i] = '0'
		// Если это уже самая первая цифра слева, то добавить "1" в начало
		if i == 0 {
			numberToRound = append([]rune{'1'}, numberToRound...)
		}

		return true, numberToRound
	}

	// Если не нужно было увеличивать цифру
	// Если текущая цифра <= 4, то завершить цикл
	if currentDigit <= four {
		return false, numberToRound
	}

	/* Если текущая цифра >= 5,
	то увеличить следующую цифру (соседнюю слева) */
	return true, numberToRound
}

func RoundNumber(number objects.Number, precision int64) objects.Number {
	if number.Divider == constants.FRACTIONAL_NUMBER ||
		int64(len(number.SecondPart)) <= precision || precision < 0 {
		return number
	}

	integerPart := number.FirstPart
	// Обрезать дробную часть
	decimalPart := number.SecondPart[0 : precision+1]
	numberToRound := []rune(integerPart + string(constants.DECIMAL_NUMBER) + decimalPart)
	increaseDigit := false

	// Цикл от последней цифры к первой (справа налево)
	for i := len(numberToRound) - 1; i >= 0; i-- {
		// Если текущий символ - это цифра (не знак разделителя)
		if !unicode.IsDigit(numberToRound[i]) {
			continue
		}

		if increaseDigit, numberToRound = roundDigit(increaseDigit, numberToRound, i); !increaseDigit {
			break
		}
	}

	result := string(numberToRound[0 : len(numberToRound)-1])

	if increaseDigit {
		result = "1" + result
	}

	number.FirstPart, number.SecondPart, _ = strings.Cut(result, string(constants.DECIMAL_NUMBER))
	// Убрать лишние нули из дробной части справа
	number.SecondPart = RemoveFromString(number.SecondPart, ZerosInEndRegex)
	if number.SecondPart == "" {
		number.SecondPart = "0"
	}

	return number
}
