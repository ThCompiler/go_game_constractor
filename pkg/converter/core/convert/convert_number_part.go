package convertor

import (
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/functions"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/option"
	"github.com/ThCompiler/go_game_constractor/pkg/stringutilits"
	"strings"
)

type Converter struct {
	options option.Options
}

func NewConverter(options option.Options) Converter {
	return Converter{options: options}
}

func (c *Converter) ConvertIntegerPart(convertedNumber objects.ResultNumberT, integerPart string,
	delimiter constants.NumberType) objects.ResultNumberT {
	numberAsTriplets := functions.SplitNumberIntoThrees(integerPart)

	// По умолчанию число не конвертировано в слова
	convertedNumber.FirstPart = integerPart

	// Если нужно конвертировать число в слова
	if c.options.ConvertNumberToWords.Integer {
		convertedNumber.FirstPart = c.convertTripletsToWords(
			numberAsTriplets,
			delimiter,
		)
	}

	// Если нужно отображать валюту целой части числа
	if c.options.ShowCurrency.Integer {
		// Если разделитель - не дробная черта
		if delimiter != constants.FRACTIONAL_NUMBER {
			currencyWord := c.options.Language.GetCurrencyAsWord(
				constants.INTEGER_PART,
				numberAsTriplets,
			)
			convertedNumber.FirstPartName = currencyWord
		}
	}

	return convertedNumber
}

func (c *Converter) ConvertFractionalPart(convertedNumber objects.ResultNumberT, integerPart, fractionalPart string,
	delimiter constants.NumberType) objects.ResultNumberT {
	numberAsTriplets := functions.SplitNumberIntoThrees(fractionalPart)
	IntegerPartAsTriplets := functions.SplitNumberIntoThrees(integerPart)
	// По умолчанию число не конвертировано в слова
	convertedNumber.SecondPart = fractionalPart

	// Если нужно конвертировать число в слова

	if delimiter == constants.FRACTIONAL_NUMBER {
		if c.options.ConvertNumberToWords.Fractional {
			convertedNumber.SecondPart = c.convertFractionalTripletsToWords(
				IntegerPartAsTriplets,
				numberAsTriplets,
			)
		} else { // Если не нужно конвертировать число в слова
			convertedNumber = c.convertNumberToNumberString(convertedNumber, fractionalPart)
		}
	} else {
		fractionalPart = strings.TrimRight(fractionalPart, "0")
		if fractionalPart == "" {
			fractionalPart = "0"
		}
		// По умолчанию число не конвертировано в слова
		convertedNumber.SecondPart = fractionalPart

		// Если нужно конвертировать число в слова
		if c.options.ConvertNumberToWords.Fractional {
			convertedNumber.SecondPart = c.convertTripletsToWords(
				numberAsTriplets,
				delimiter,
			)
		}

		if fractionalPart == "0" && !c.options.ShowZeroInDecimalPart {
			convertedNumber.SecondPart = ""
		}
	}

	// Если нужно отображать валюту дробной части числа
	if c.options.ShowCurrency.Fractional {
		convertedNumber = c.addCurrencyToFractionalPart(convertedNumber, fractionalPart, numberAsTriplets, delimiter)
	}

	return convertedNumber
}

func (c *Converter) addCurrencyToFractionalPart(convertedNumber objects.ResultNumberT, fractionalPart string,
	numberAsTriplets []objects.RuneDigitTriplet, delimiter constants.NumberType) objects.ResultNumberT {
	// Если валюта - не 'number'
	if c.options.Language.IsCurrency() {
		// Если разделитель - дробная черта
		if delimiter == constants.FRACTIONAL_NUMBER {
			convertedNumber.SecondPartName = c.options.Language.GetCurrencyForFractionalNumber()
			return convertedNumber
		}

		currencyWord := c.options.Language.GetCurrencyAsWord(
			constants.FRACTIONAL_PART,
			numberAsTriplets,
		)

		// Если определено число дробной части
		if convertedNumber.SecondPart != "" {
			// Добавить валюту к дробной части
			convertedNumber.SecondPartName = currencyWord
		}

		return convertedNumber
	}
	// Если валюта указана как "number"

	// Если разделитель дробная черта
	if delimiter == constants.FRACTIONAL_NUMBER {
		return convertedNumber
	}

	return c.setDecimalNumberCurrency(convertedNumber, fractionalPart)
}

func (c *Converter) convertNumberToNumberString(convertedNumber objects.ResultNumberT,
	fractionalPart string,
) objects.ResultNumberT {
	// Если валюта "number"
	if c.options.Language.IsNumber() {
		// Если в дробной части есть цифры
		if len(fractionalPart) > 0 {
			// Удалить лишние нули перед числом
			fractionalPart = strings.TrimLeft(fractionalPart, "0")
			//fractionalPart = functions.RemoveFromString(fractionalPart, `^0+`) // TODO: check if it works
			// Если после удаления лишних нулей не осталось цифр, то добавить один "0"
			if fractionalPart == "" && c.options.RoundNumber != 0 {
				fractionalPart = "0"
			}
		}

		convertedNumber.SecondPart = fractionalPart
	}

	return convertedNumber
}

func (c *Converter) setDecimalNumberCurrency(convertedNumber objects.ResultNumberT,
	fractionalPart string) objects.ResultNumberT {
	// Если имеет смысл добавлять название валюты
	if len(fractionalPart) > 0 && convertedNumber.SecondPart != "" {
		runeFractionalScalesArray := []rune(fractionalPart)
		digitToConvert := objects.Digit(
			stringutilits.ToDigit(
				runeFractionalScalesArray[len(runeFractionalScalesArray)-1],
			),
		)

		convertedNumber.SecondPartName = c.options.Language.GetEndingOfDecimalNumberForFractionalPart(
			len(fractionalPart)-1,
			digitToConvert,
		)
	}

	return convertedNumber
}

func (c *Converter) convertTripletsToWords(
	numberByTriplets []objects.RuneDigitTriplet,
	delimiter constants.NumberType,
) string {
	if len(numberByTriplets) == 0 || (len(numberByTriplets) == 1 && numberByTriplets[0].IsZeros()) {
		return strings.TrimSpace(c.options.Language.ConvertZeroToWordsForIntegerPart())
	}
	numberScalesArrayLen := len(numberByTriplets)
	convertedResult := ""

	// Для каждого класса числа
	for arrIndex, numberTriplet := range numberByTriplets {
		// Определить порядковый номер текущего класса числа
		currentNumberScale := numberScalesArrayLen - arrIndex

		digits := numberTriplet.ToNumeric()

		// Если класс числа пустой (000)
		if digits.IsZeros() {
			// Если нет классов выше
			if numberScalesArrayLen == 1 {
				// Выйти из цикла
				break
			}
			// Пропустить этот пустой класс (000)
			continue
		}

		stringDigits := c.options.Language.ConvertTripletToWords(
			delimiter,
			digits,
			currentNumberScale,
		)

		scaleName := c.options.Language.GetWordScaleName(delimiter, currentNumberScale-1, digits)

		// Убрать ненужный "ноль"
		if digits.Units == 0 && (digits.Hundreds > 0 || digits.Dozens > 0) {
			stringDigits.Units = ""
		}

		// Соединить значения в одну строку
		scaleResult := strings.TrimSpace(stringDigits.Hundreds) + " " +
			strings.TrimSpace(stringDigits.Dozens) + " " +
			strings.TrimSpace(stringDigits.Units) + " " +
			strings.TrimSpace(scaleName)

		// Добавить текущий разобранный класс к общему результату
		convertedResult += " " + scaleResult
	}

	// Вернуть полученный результат и форму падежа для валюты
	return strings.TrimSpace(convertedResult)
}

func (c *Converter) convertFractionalTripletsToWords(
	integerPartTriplets []objects.RuneDigitTriplet,
	fractionalPartTriplets []objects.RuneDigitTriplet,
) string {
	if len(fractionalPartTriplets) < 1 {
		return ""
	}

	convertedResult := ""

	// Удалить лишние нули в начале числа
	updatedNumberTriplets := functions.RemoveZeroTripletFromBeginning(fractionalPartTriplets)

	/* Определить индекс класса, который является последним.
	После него могут быть только классы с "000".
	0 - единицы, 1 - тысячи, 2 - миллионы и т.д. */
	lastNotNullTriplet := functions.IndexByEndOfLastNotZeroTriplet(updatedNumberTriplets)

	// Если нет ни одного не пустого класса
	if lastNotNullTriplet == -1 {
		// Вернуть ноль
		return c.options.Language.ConvertZeroToWordsForFractionalNumber(integerPartTriplets)
	}

	// Индекс последнего класса в массиве.
	lastTripletIndex := len(updatedNumberTriplets) - lastNotNullTriplet - 1
	/* Если есть не пустые классы до последнего не пустого класса,
	то конвертировать как обычное число */
	if lastTripletIndex > 0 {
		// Получить массив классов, в котором последний класс будет пустым.
		numberScalesArrForCommonConvert := make([]objects.RuneDigitTriplet, len(updatedNumberTriplets))
		copy(numberScalesArrForCommonConvert, updatedNumberTriplets)
		numberScalesArrForCommonConvert[lastTripletIndex] = objects.RuneDigitTriplet{
			Units:    '0',
			Dozens:   '0',
			Hundreds: '0',
		}

		// Конвертировать классы как обычное число
		convertedResult += c.convertTripletsToWords(
			numberScalesArrForCommonConvert,
			constants.FRACTIONAL_NUMBER, // TODO: изменить падежи для дробных числе! //c.options.Language.CorrectNumberInfoForFractionalTriplets(),
		) + " "
	}

	// Если последний класс для конвертирования - тысячи или больше
	if lastNotNullTriplet >= 1 {
		convertedResult += c.options.Language.ConvertNotLowestScaleToWords(
			updatedNumberTriplets[lastTripletIndex].ToNumeric(),
			lastNotNullTriplet,
			(len(updatedNumberTriplets)-1) == lastNotNullTriplet,
			integerPartTriplets,
		)
	}

	// Если последний класс для конвертирования - единицы
	if lastNotNullTriplet == 0 {
		convertedResult += c.options.Language.ConvertLowestScaleToWords(
			updatedNumberTriplets[lastTripletIndex].ToNumeric(),
			integerPartTriplets,
		)
	}

	return strings.TrimSpace(convertedResult)
}
