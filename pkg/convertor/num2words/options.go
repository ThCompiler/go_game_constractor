package convertor

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/currency"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words/declension"
)

var DefaultOption = NewOptions(currency.RUB, declension.NOMINATIVE, -1,
	true, NumberPart{Fractional: true, Integer: true},
	NumberPart{Fractional: true, Integer: true},
	NumberPart{Fractional: true, Integer: true}, nil)

type Options struct {
	/**
	* Select currency\
	* `'rub'` Russian ruble. 124 рубля 42 копейки\
	* `'usd'` Dollar. 124 доллара 42 цента\
	* `'eur'` Euro. 124 евро 42 цента\
	* `'number'` Number without currency. 124 целых 42 сотых\
	* `Object` Custom currency. 124 юаня 42 фыня\
	* Default: `'rub'`
	 */
	currency currency.Currency

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
	declension declension.Declension

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
	currencyInfo *currency.CustomCurrency
}

func NewOptions(currency currency.Currency, declension declension.Declension,
	roundNumber int64, convertMinusSignToWord bool,
	showNumberParts, convertNumberToWords, showCurrency NumberPart,
	currencyInfo *currency.CustomCurrency,
) Options {
	return Options{
		currency:               currency,
		declension:             declension,
		roundNumber:            roundNumber,
		convertMinusSignToWord: convertMinusSignToWord,
		showCurrency:           showCurrency,
		showNumberParts:        showNumberParts,
		convertNumberToWords:   convertNumberToWords,
		currencyInfo:           currencyInfo,
	}
}

func (o *Options) getCurrencyObject() currency.CustomCurrency {
	if o.currency == currency.CUSTOM {
		return *o.currencyInfo
	}

	return words.WordConstants.N2w.CurrencyStrings.Currencies[o.currency]
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
