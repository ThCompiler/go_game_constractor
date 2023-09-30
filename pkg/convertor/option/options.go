package option

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/core/words"
)

type Options struct {
	/**
	* Select Language for converting
	 */
	Language words.Language

	/**
	* Select currency\
	* `'rub'` Russian ruble. 124 рубля 42 копейки\
	* `'usd'` Dollar. 124 доллара 42 цента\
	* `'eur'` Euro. 124 евро 42 цента\
	* `'number'` Number without currency. 124 целых 42 сотых\
	* `Object` Custom currency. 124 юаня 42 фыня\
	* Default: `'rub'`
	 */
	CurrencyName words.CurrencyName

	/**
	 * Select Declension\
	 * `'nominative'` Одна тысяча два рубля\
	 * `'genitive'` Одной тысячи двух рублей\
	 * `'dative'` Одной тысяче двум рублям\
	 * `'accusative'` Одну тысячу два рубля\
	 * `'instrumental'` Одной тысячей двумя рублями\
	 * `'prepositional'` Одной тысяче двух рублях\
	 * Default: `nominative`
	 */
	Declension words.Declension

	/**
	 * Rounding\
	 * `-1` Rounding disabled\
	 * `0` and more. Precision of rounding\
	 * Default: `-1`
	 */
	RoundNumber int64

	/**
	 * Convert minus sign to word\
	 * `true` Минус\
	 * `false` -\
	 * Default: `true`
	 */
	ConvertMinusSignToWord bool

	/**
	 * Show number parts
	 */
	ShowNumberParts NumberPart

	/**
	 * Convert number parts to words
	 */
	ConvertNumberToWords NumberPart

	/**
	 * Show currency
	 * `Object`
	 */
	ShowCurrency NumberPart
}

func NewOptionsWithCurrency(language words.Language, declension words.Declension,
	roundNumber int64, convertMinusSignToWord bool,
	showNumberParts, convertNumberToWords, showCurrency NumberPart,
	currencyName words.CurrencyName,
) Options {
	return Options{
		Language:               language,
		Declension:             declension,
		RoundNumber:            roundNumber,
		ConvertMinusSignToWord: convertMinusSignToWord,
		ShowCurrency:           showCurrency,
		ShowNumberParts:        showNumberParts,
		ConvertNumberToWords:   convertNumberToWords,
		CurrencyName:           currencyName,
	}
}

func NewOptionsWithoutCurrency(language words.Language, declension words.Declension,
	roundNumber int64, convertMinusSignToWord bool,
	showNumberParts, convertNumberToWords, showCurrency NumberPart,
) Options {
	return Options{
		Language:               language,
		Declension:             declension,
		RoundNumber:            roundNumber,
		ConvertMinusSignToWord: convertMinusSignToWord,
		ShowCurrency:           showCurrency,
		ShowNumberParts:        showNumberParts,
		ConvertNumberToWords:   convertNumberToWords,
		CurrencyName:           words.NUMBER,
	}
}

type NumberPart struct {
	/**
	 * Show fractional part of number\
	 * `true` Два рубля **пять копеек**\
	 * `false` Два рубля\
	 * Default: `true`
	 */
	Fractional bool
	/**
	 * Show integer part of number\
	 * `true` **Два рубля** пять копеек\
	 * `false` Пять копеек\
	 * Default: `true`
	 */
	Integer bool
}
