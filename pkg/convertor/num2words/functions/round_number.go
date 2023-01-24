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

func roundDigitInLoop(increaseDigit bool, currentDigit int8, numberPartToRound []rune, i int) (bool, []rune, bool) {
	// Если нужно было увеличивать цифру
	if increaseDigit {
		if currentDigit != nine { // Если любая другая цифра
			numberPartToRound[i]++

			return increaseDigit, numberPartToRound, false
		}
		// Если текущая цифра 9, то увеличить следующую
		numberPartToRound[i] = '0'
		// Если это уже самая первая цифра слева, то добавить "1" в начало
		if i == 0 {
			numberPartToRound = append([]rune{'1'}, numberPartToRound...)
		}
	}

	// Если не нужно было увеличивать цифру
	// Если текущая цифра <= 4, то завершить цикл
	if currentDigit <= four {
		return increaseDigit, numberPartToRound, false
	}
	/* Если текущая цифра >= 5,
	   то увеличить следующую цифру (соседнюю слева) */
	increaseDigit = true

	return increaseDigit, numberPartToRound, true
}

func RoundNumber(number objects.Number, precision int64) objects.Number {
	if number.Divider == constants.FractionalNumber ||
		int64(len(number.SecondPart)) <= precision || precision < 0 {
		return number
	}

	integerPart := number.FirstPart
	// Обрезать дробную часть
	fractionalPart := number.SecondPart[0 : precision+1]
	numberPartToRound := []rune(integerPart + `.` + fractionalPart)
	increaseDigit := false
	continueLoop := true

	// Цикл от последней цифры к первой (справа налево)
	for i := len(numberPartToRound) - 1; i >= 0 && continueLoop; i-- {
		// Если текущий символ - это цифра (не знак разделителя)
		if unicode.IsDigit(numberPartToRound[i]) {
			currentDigit := stringutilits.ToDigit(numberPartToRound[i])
			increaseDigit, numberPartToRound, continueLoop = roundDigitInLoop(increaseDigit, currentDigit,
				numberPartToRound, i)
		}
	}

	result := string(numberPartToRound[0 : len(numberPartToRound)-1])

	number.FirstPart, number.SecondPart, _ = strings.Cut(result, ".")
	// Убрать лишние нули из дробной части справа
	number.SecondPart = ClearFromString(number.SecondPart, `^0+/`)
	if number.SecondPart == "" && precision != 0 {
		number.SecondPart = "0"
	}

	return number
}
