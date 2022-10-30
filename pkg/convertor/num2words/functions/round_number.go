package functions

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/num2words/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/stringutilits"
	"strings"
	"unicode"
)

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
	// Цикл от последней цифры к первой (справа налево)
	for i := len(numberPartToRound) - 1; i >= 0; i-- {
		// Если текущий символ - это цифра (не знак разделителя)
		if unicode.IsDigit(numberPartToRound[i]) {
			currentDigit := stringutilits.ToDigit(numberPartToRound[i])
			// Если нужно было увеличивать цифру
			if increaseDigit {
				// Если текущая цифра 9, то увеличить следующую
				if currentDigit == 9 {
					numberPartToRound[i] = '0'
					increaseDigit = true
					// Если это уже самая первая цифра слева, то добавить "1" в начало
					if i == 0 {
						numberPartToRound = append([]rune{'1'}, numberPartToRound...)
					}
					// Если любая другая цифра
				} else {
					numberPartToRound[i]++
					increaseDigit = false
					break
				}
				// Если не нужно было увеличивать цифру
			} else {
				// Если текущая цифра <= 4, то завершить цикл
				if currentDigit <= 4 {
					break
				} else {
					/* Если текущая цифра >= 5,
					   то увеличить следующую цифру (соседнюю слева) */
					increaseDigit = true
				}
			}
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
