package convertor

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/currency"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/words"
)

type Options struct {
	/**
	* Select language for converting
	 */
	language words.Language

	/**
	* Select currency\
	* `'rub'` Russian ruble. 124 рубля 42 копейки\
	* `'usd'` Dollar. 124 доллара 42 цента\
	* `'eur'` Euro. 124 евро 42 цента\
	* `'number'` Number without currency. 124 целых 42 сотых\
	* `Object` Custom currency. 124 юаня 42 фыня\
	* Default: `'rub'`
	 */
	currencyName currency.Name

	/**
	 * Select declension\
	 * `'nominative'` Одна тысяча два рубля\
	 * `'genitive'` Одной тысячи двух рублей\
	 * `'dative'` Одной тысяче двум рублям\
	 * `'accusative'` Одну тысячу два рубля\
	 * `'instrumental'` Одной тысячей двумя рублями\
	 * `'prepositional'` Одной тысяче двух рублях\
	 * Default: `nominative`
	 */
	declension words.Declension

	/**
	 * Rounding\
	 * `-1` Rounding disabled\
	 * `0` and more. Precision of rounding\
	 * Default: `-1`
	 */
	roundNumber int64

	/**
	 * Convert minus sign to word\
	 * `true` Минус\
	 * `false` -\
	 * Default: `true`
	 */
	convertMinusSignToWord bool

	/**
	 * Show number parts
	 */
	showNumberParts NumberPart

	/**
	 * Convert number parts to words
	 */
	convertNumberToWords NumberPart

	/**
	 * Show currency
	 * `Object`
	 */
	showCurrency NumberPart
}

func NewOptionsWithCurrency(language words.Language, declension words.Declension,
	roundNumber int64, convertMinusSignToWord bool,
	showNumberParts, convertNumberToWords, showCurrency NumberPart,
	currencyName currency.Name,
) Options {
	return Options{
		language:               language,
		declension:             declension,
		roundNumber:            roundNumber,
		convertMinusSignToWord: convertMinusSignToWord,
		showCurrency:           showCurrency,
		showNumberParts:        showNumberParts,
		convertNumberToWords:   convertNumberToWords,
		currencyName:           currencyName,
	}
}

func NewOptionsWithoutCurrency(language words.Language, declension words.Declension,
	roundNumber int64, convertMinusSignToWord bool,
	showNumberParts, convertNumberToWords, showCurrency NumberPart,
) Options {
	return Options{
		language:               language,
		declension:             declension,
		roundNumber:            roundNumber,
		convertMinusSignToWord: convertMinusSignToWord,
		showCurrency:           showCurrency,
		showNumberParts:        showNumberParts,
		convertNumberToWords:   convertNumberToWords,
		currencyName:           currency.NUMBER,
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
