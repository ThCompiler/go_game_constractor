package convertor

import (
	"strings"
	"unicode"

	"github.com/ThCompiler/go_game_constractor/pkg/convertor/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/currency"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/num2words/functions"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/num2words/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words/declension"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words/genders"
	"github.com/ThCompiler/go_game_constractor/pkg/stringutilits"
)

const twoSignAfterRound = 2

func combineResultData(number objects.Number, appliedOptions Options) string {
	/* Пример convertedNumberArr:
	   ['минус', 'двадцать два', 'рубля', 'сорок одна', 'копейка'] */
	convertedNumber := objects.ResultNumberT{}
	modifiedNumber := number

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
	modifiedNumber = functions.RoundNumber(number, appliedOptions.roundNumber)
	// Если указана валюта
	if appliedOptions.currency != "" && appliedOptions.currency != currency.NUMBER {
		// Округлить число до 2 знаков после запятой
		modifiedNumber = functions.RoundNumber(modifiedNumber, twoSignAfterRound)
	}

	// Если нужно отображать целую часть числа
	if appliedOptions.showNumberParts.Integer {
		convertedNumber = combineIntegerPart(convertedNumber, modifiedNumber.FirstPart, modifiedNumber.Divider,
			appliedOptions)
	}
	// Если нужно отображать дробную часть числа
	if appliedOptions.showNumberParts.Fractional {
		convertedNumber = combineFractionalPart(convertedNumber, modifiedNumber.FirstPart,
			modifiedNumber.SecondPart, modifiedNumber.Divider,
			appliedOptions)
	}
	// Объединить полученный результат
	convertedNumberResult := convertedNumber.Sign + " " +
		convertedNumber.FirstPart + " " + convertedNumber.FirstPartName + " " +
		convertedNumber.SecondPart + " " + convertedNumber.SecondPartName

	convertedNumberResult = strings.TrimSpace(functions.ReplaceInString(convertedNumberResult, `\s+`, ` `))

	// Сделать первую букву заглавной
	r := []rune(convertedNumberResult)

	return string(append([]rune{unicode.ToUpper(r[0])}, r[1:]...))
}

func combineIntegerPart(convertedNumber objects.ResultNumberT, integerPart string,
	delimiter constants.NumberType, appliedOptions Options,
) objects.ResultNumberT {
	// По умолчанию число не конвертировано в слова
	convertedNumber.FirstPart = integerPart
	// Получить результат конвертирования числа
	convertedIntegerObject := functions.ConvertEachScaleToWords(
		functions.SplitNumberIntoThrees(integerPart),
		appliedOptions.getCurrencyObject().CurrencyNounGender.Integer,
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
			convertedNumber.FirstPart = functions.ConvertEachScaleToWords(
				functions.SplitNumberIntoThrees(integerPart),
				genders.FEMALE,
				appliedOptions.declension,
			).Result
		}
	}
	// Если нужно отображать валюту целой части числа
	if appliedOptions.showCurrency.Integer {
		// Если разделитель - не дробная черта
		if delimiter != constants.FractionalNumber {
			currencyWord := functions.GetCurrencyWord(
				appliedOptions.getCurrencyObject(),
				constants.DecimalNumber,
				convertedIntegerObject.UnitNameForm,
				convertedIntegerObject.LastScaleIsZero,
				appliedOptions.currency,
				appliedOptions.declension,
			)
			convertedNumber.FirstPartName = currencyWord
		}
	}

	return convertedNumber
}

func combineFractionalPart(convertedNumber objects.ResultNumberT, integerPart, fractionalPart string,
	delimiter constants.NumberType, appliedOptions Options,
) objects.ResultNumberT {
	// По умолчанию число не конвертировано в слова
	convertedNumber.SecondPart = fractionalPart
	// Получить результат конвертирования числа
	convertedFractionalObject := functions.ConvertEachScaleToWords(
		functions.SplitNumberIntoThrees(fractionalPart),
		appliedOptions.getCurrencyObject().CurrencyNounGender.FractionalPart,
		appliedOptions.declension,
	)

	// Если нужно конвертировать число в слова
	if appliedOptions.convertNumberToWords.Fractional {
		convertedNumber = convertNumberToWord(convertedNumber, integerPart, fractionalPart, delimiter, appliedOptions,
			convertedFractionalObject)
	} else { // Если не нужно конвертировать число в слова
		convertedNumber = convertNumberToNumberString(convertedNumber, fractionalPart, appliedOptions)
	}
	// Если нужно отображать валюту дробной части числа
	if appliedOptions.showCurrency.Fractional {
		convertedNumber = addCurrencyToFractionalPart(convertedNumber, fractionalPart, convertedFractionalObject,
			appliedOptions, delimiter)
	}

	return convertedNumber
}

func convertNumberToWord(convertedNumber objects.ResultNumberT, integerPart, fractionalPart string,
	delimiter constants.NumberType, appliedOptions Options, defaultPart objects.ConvertedScalesToWords,
) objects.ResultNumberT {
	// Если разделитель - не дробная черта
	// Применить конвертированное число
	if delimiter != constants.FractionalNumber {
		convertedNumber.SecondPart = defaultPart.Result

		return convertedNumber
	}

	// Если разделитель - дробная черта
	convertedIntegerObject := functions.ConvertEachScaleToWords(
		functions.SplitNumberIntoThrees(integerPart),
		appliedOptions.getCurrencyObject().CurrencyNounGender.Integer,
		appliedOptions.declension,
	)

	convertedNumber.SecondPart = functions.ConvertEachScaleToWordsSlash(
		functions.SplitNumberIntoThrees(fractionalPart),
		convertedIntegerObject.UnitNameForm,
		appliedOptions.declension,
	)

	return convertedNumber
}

func addCurrencyToFractionalPart(convertedNumber objects.ResultNumberT, fractionalPart string,
	fractionalObject objects.ConvertedScalesToWords, appliedOptions Options,
	delimiter constants.NumberType,
) objects.ResultNumberT {
	// Если валюта - не 'number'
	if appliedOptions.currency != currency.NUMBER {
		convertedNumber = setNotNumberCurrency(convertedNumber, fractionalObject, appliedOptions)
	}

	// Если валюта указана как "number"
	if appliedOptions.currency == currency.NUMBER {
		// Если разделитель - не дробная черта
		convertedNumber = setNumberCurrency(convertedNumber, fractionalPart, fractionalObject,
			appliedOptions, delimiter)
	}

	// Если разделитель - дробная черта
	if delimiter == constants.FractionalNumber {
		// Если указана валюта
		if appliedOptions.currency != currency.NUMBER {
			convertedNumber.SecondPartName = appliedOptions.getCurrencyObject().DecimalCurrencyNameDeclensions[declension.GENITIVE][0]
		}
	}

	return convertedNumber
}

func convertNumberToNumberString(convertedNumber objects.ResultNumberT,
	fractionalPart string, appliedOptions Options,
) objects.ResultNumberT {
	// Если валюта "number"
	if appliedOptions.currency == currency.NUMBER {
		// Если в дробной части есть цифры
		if len(fractionalPart) > 0 {
			// Удалить лишние нули перед числом
			fractionalPart = functions.RemoveFromString(fractionalPart, `^0+`)
			// Если после удаления лишних нулей не осталось цифр, то добавить один "0"
			if fractionalPart == "" && appliedOptions.roundNumber != 0 {
				fractionalPart = "0"
			}
		}

		convertedNumber.SecondPart = fractionalPart
	}

	return convertedNumber
}

func setNotNumberCurrency(convertedNumber objects.ResultNumberT, fractionalObject objects.ConvertedScalesToWords,
	appliedOptions Options,
) objects.ResultNumberT {
	currencyWord := functions.GetCurrencyWord(
		appliedOptions.getCurrencyObject(),
		constants.FractionalNumber,
		fractionalObject.UnitNameForm,
		fractionalObject.LastScaleIsZero,
		appliedOptions.currency,
		appliedOptions.declension,
	)
	// Если определено число дробной части
	if convertedNumber.SecondPart != "" {
		// Добавить валюту к дробной части
		convertedNumber.SecondPartName = currencyWord
	}

	return convertedNumber
}

func setNumberCurrency(convertedNumber objects.ResultNumberT, fractionalPart string,
	fractionalObject objects.ConvertedScalesToWords, appliedOptions Options,
	delimiter constants.NumberType,
) objects.ResultNumberT {
	// Если разделитель - не дробная черта
	if delimiter == constants.FractionalNumber {
		return convertedNumber
	}

	// Если имеет смысл добавлять название валюты
	if appliedOptions.roundNumber > 0 ||
		(appliedOptions.roundNumber < 0 && len(fractionalPart) > 0) {
		runeFractionalScalesArray := []rune(fractionalPart)
		digitToConvert := objects.Digit(
			stringutilits.ToDigit(
				runeFractionalScalesArray[len(runeFractionalScalesArray)-1],
			),
		)

		convertedNumber.SecondPartName = functions.GetFractionalUnitCurrencyNumber(
			len(fractionalPart)-1,
			digitToConvert,
			appliedOptions.declension,
			fractionalObject.UnitNameForm,
		)
	}

	return convertedNumber
}
