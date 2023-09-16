package convertor

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/currency"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/functions"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/stringutilits"
)

type Converter struct {
	options Options
}

func (c *Converter) ConvertIntegerPart(convertedNumber objects.ResultNumberT, integerPart string,
	delimiter constants.NumberType) objects.ResultNumberT {
	numberAsTriplets := functions.SplitNumberIntoThrees(integerPart)

	// По умолчанию число не конвертировано в слова
	convertedNumber.FirstPart = integerPart

	// Если нужно конвертировать число в слова
	if c.options.convertNumberToWords.Integer {
		convertedNumber.FirstPart = c.options.language.
			ConvertIntegerPartTripletsToWords(c.options.currencyName, numberAsTriplets, delimiter, c.options.declension)
	}

	// Если нужно отображать валюту целой части числа
	if c.options.showCurrency.Integer {
		// Если разделитель - не дробная черта
		if delimiter != constants.FRACTIONAL_NUMBER {
			currencyWord := c.options.language.GetCurrencyAsWord(
				c.options.currencyName,
				constants.DECIMAL_NUMBER,
				numberAsTriplets,
				c.options.declension,
			)
			convertedNumber.FirstPartName = currencyWord
		}
	}

	return convertedNumber
}

func (c *Converter) ConvertFractionalPart(convertedNumber objects.ResultNumberT, integerPart, fractionalPart string,
	delimiter constants.NumberType) objects.ResultNumberT {
	numberAsTriplets := functions.SplitNumberIntoThrees(fractionalPart)

	// По умолчанию число не конвертировано в слова
	convertedNumber.SecondPart = fractionalPart

	// Если нужно конвертировать число в слова
	if c.options.convertNumberToWords.Fractional {
		convertedNumber.SecondPart = c.options.language.ConvertFractionalPartTripletsToWords(
			c.options.currencyName,
			numberAsTriplets,
			delimiter,
			c.options.declension,
		)
	} else { // Если не нужно конвертировать число в слова
		convertedNumber = c.convertNumberToNumberString(convertedNumber, fractionalPart)
	}

	// Если нужно отображать валюту дробной части числа
	if c.options.showCurrency.Fractional {
		convertedNumber = c.addCurrencyToFractionalPart(convertedNumber, fractionalPart, numberAsTriplets, delimiter)
	}

	return convertedNumber
}

func (c *Converter) addCurrencyToFractionalPart(convertedNumber objects.ResultNumberT, fractionalPart string,
	numberAsTriplets []objects.RuneDigitTriplet, delimiter constants.NumberType) objects.ResultNumberT {
	// Если валюта - не 'number'
	if c.options.currencyName != currency.NUMBER {
		// Если разделитель - дробная черта
		if delimiter == constants.FRACTIONAL_NUMBER {
			convertedNumber.SecondPartName = c.options.language.GetCurrencyForFractionalNumber()
			return convertedNumber
		}

		currencyWord := c.options.language.GetCurrencyAsWord(
			c.options.currencyName,
			constants.FRACTIONAL_NUMBER,
			numberAsTriplets,
			c.options.declension,
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
	if c.options.currencyName == currency.NUMBER {
		// Если в дробной части есть цифры
		if len(fractionalPart) > 0 {
			// Удалить лишние нули перед числом
			fractionalPart = functions.RemoveFromString(fractionalPart, `^0+`)
			// Если после удаления лишних нулей не осталось цифр, то добавить один "0"
			if fractionalPart == "" && c.options.roundNumber != 0 {
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
	if c.options.roundNumber > 0 ||
		(c.options.roundNumber < 0 && len(fractionalPart) > 0) {
		runeFractionalScalesArray := []rune(fractionalPart)
		digitToConvert := objects.Digit(
			stringutilits.ToDigit(
				runeFractionalScalesArray[len(runeFractionalScalesArray)-1],
			),
		)

		convertedNumber.SecondPartName = c.options.language.GetEndingOfDecimalNumberForFractionalPart(
			len(fractionalPart)-1,
			digitToConvert,
			c.options.declension,
		)
	}

	return convertedNumber
}
