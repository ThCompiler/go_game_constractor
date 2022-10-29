package convertor

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/currency"
	functions2 "github.com/ThCompiler/go_game_constractor/pkg/convertor/num2words/functions"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/num2words/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words/declension"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words/genders"
	"github.com/ThCompiler/go_game_constractor/pkg/stringutilits"
	"strings"
	"unicode"
)

func combineResultData(number objects.Number, appliedOptions Options) string {
	/* Пример convertedNumberArr:
	['минус', 'двадцать два', 'рубля', 'сорок одна', 'копейка'] */
	convertedNumber := objects.ResultNumberT{}
	modifiedNumber := number

	// Определить отображение валюты
	currencyObject := appliedOptions.getCurrencyObject()

	// Если есть знак минус
	if number.Sign == "-" {
		// Если отображать знак минус словом
		if appliedOptions.convertMinusSignToWord {
			convertedNumber.Sign = words.WordConstants.N2w.Sign.Minus
		} else {
			convertedNumber.Sign = "-"
		}
	}

	// Округлить число до заданной точности
	modifiedNumber = functions2.RoundNumber(number, appliedOptions.roundNumber)
	// Если указана валюта
	if appliedOptions.currency != "" && appliedOptions.currency != currency.NUMBER {
		// Округлить число до 2 знаков после запятой
		modifiedNumber = functions2.RoundNumber(modifiedNumber, 2)
	}

	integerScalesArray := modifiedNumber.FirstPart
	fractionalScalesArray := modifiedNumber.SecondPart
	delimiter := modifiedNumber.Divider

	// Если нужно отображать целую часть числа
	if appliedOptions.showNumberParts.Integer {
		// По умолчанию число не конвертировано в слова
		convertedNumber.FirstPart = integerScalesArray
		// Получить результат конвертирования числа
		convertedIntegerObject := functions2.ConvertEachScaleToWords(
			functions2.NumberToScales(integerScalesArray),
			currencyObject.CurrencyNounGender.Integer,
			appliedOptions.declension,
		)

		// Если нужно конвертировать число в слова
		if appliedOptions.convertNumberToWords.Integer {
			// Если разделитель - не дробная черта
			if delimiter != constants.FractionalNumber {
				// Применить конвертированное число
				convertedNumber.FirstPart = convertedIntegerObject.Result
			} else {
				// Если разделитель - дробная черта
				// Род числа всегда женский ('одна', 'две')
				// Применить конвертированное число
				convertedNumber.FirstPart = functions2.ConvertEachScaleToWords(
					functions2.NumberToScales(integerScalesArray),
					genders.FEMALE,
					appliedOptions.declension,
				).Result
			}
		}
		// Если нужно отображать валюту целой части числа
		if appliedOptions.showCurrency.Integer {
			// Если разделитель - не дробная черта
			if delimiter != constants.FractionalNumber {
				currencyWord := functions2.GetCurrencyWord(
					currencyObject,
					constants.DecimalNumber,
					convertedIntegerObject.UnitNameForm,
					convertedIntegerObject.LastScaleIsZero,
					appliedOptions.currency,
					appliedOptions.declension,
				)
				convertedNumber.FirstPartName = currencyWord
			}
		}
	}
	// Если нужно отображать дробную часть числа
	if appliedOptions.showNumberParts.Fractional {
		// По умолчанию число не конвертировано в слова
		convertedNumber.SecondPart = fractionalScalesArray
		// Получить результат конвертирования числа
		convertedFractionalObject := functions2.ConvertEachScaleToWords(
			functions2.NumberToScales(fractionalScalesArray),
			currencyObject.CurrencyNounGender.FractionalPart,
			appliedOptions.declension,
		)

		// Если нужно конвертировать число в слова
		if appliedOptions.convertNumberToWords.Fractional {
			// Если разделитель - дробная черта
			if delimiter == constants.FractionalNumber {
				convertedIntegerObject := functions2.ConvertEachScaleToWords(
					functions2.NumberToScales(integerScalesArray),
					currencyObject.CurrencyNounGender.Integer,
					appliedOptions.declension,
				)

				convertedNumber.SecondPart = functions2.ConvertEachScaleToWordsSlash(
					functions2.NumberToScales(fractionalScalesArray),
					convertedIntegerObject.UnitNameForm,
					appliedOptions.declension,
				)
			} else {
				// Если разделитель - не дробная черта
				// Применить конвертированное число
				convertedNumber.SecondPart = convertedFractionalObject.Result
			}
		} else { // Если не нужно конвертировать число в слова
			// Если валюта "number"
			if appliedOptions.currency == currency.NUMBER {
				// Если в дробной части есть цифры
				if len(convertedNumber.SecondPart) > 0 {
					// Удалить лишние нули перед числом
					convertedNumber.SecondPart = functions2.ClearFromString(convertedNumber.SecondPart, `^0+`)
					// Если после удаления лишних нулей не осталось цифр, то добавить один "0"
					if convertedNumber.SecondPart == "" && appliedOptions.roundNumber != 0 {
						convertedNumber.SecondPart = "0"
					}
				}
			}
		}
		// Если нужно отображать валюту дробной части числа
		if appliedOptions.showCurrency.Fractional {
			// Если валюта - не 'number'
			if appliedOptions.currency != currency.NUMBER {
				currencyWord := functions2.GetCurrencyWord(
					currencyObject,
					constants.FractionalNumber,
					convertedFractionalObject.UnitNameForm,
					convertedFractionalObject.LastScaleIsZero,
					appliedOptions.currency,
					appliedOptions.declension,
				)
				// Если определено число дробной части
				if convertedNumber.SecondPart != "" {
					// Добавить валюту к дробной части
					convertedNumber.SecondPartName = currencyWord
				}
			}
			// Если валюта указана как "number"
			if appliedOptions.currency == currency.NUMBER {
				// Если разделитель - не дробная черта
				if delimiter != constants.FractionalNumber {
					// Если имеет смысл добавлять название валюты
					if appliedOptions.roundNumber > 0 ||
						(appliedOptions.roundNumber < 0 && len(fractionalScalesArray) > 0) {
						runeFractionalScalesArray := []rune(fractionalScalesArray)
						digitToConvert :=
							objects.Digit(stringutilits.ToDigit(runeFractionalScalesArray[len(runeFractionalScalesArray)-1]))

						convertedNumber.SecondPartName = functions2.GetFractionalUnitCurrencyNumber(
							len(fractionalScalesArray)-1,
							digitToConvert,
							appliedOptions.declension,
							convertedFractionalObject.UnitNameForm,
						)
					}
				}
			}
			// Если разделитель - дробная черта
			if delimiter == constants.FractionalNumber {
				// Если указана валюта
				if appliedOptions.currency != currency.NUMBER {
					convertedNumber.SecondPartName =
						currencyObject.CurrencyNameDeclensions[declension.GENITIVE][0]
				}
			}
		}
	}
	// Объединить полученный результат
	convertedNumberResult := convertedNumber.Sign + " " +
		convertedNumber.FirstPart + " " + convertedNumber.FirstPartName + " " +
		convertedNumber.SecondPart + " " + convertedNumber.SecondPartName

	convertedNumberResult = strings.TrimSpace(functions2.ReplaceInString(convertedNumberResult, `\s+`, ` `))

	// Сделать первую букву заглавной
	r := []rune(convertedNumberResult)
	return string(append([]rune{unicode.ToUpper(r[0])}, r[1:]...))
}
