package option

import (
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/words"
)

type Options struct {
	/**
	* Select Language for converting
	 */
	Language words.Language

	/**
	 * Rounding\
	 * `-1` Rounding disabled\
	 * `0` and more. Precision of rounding\
	 * Default: `3`
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

	/**
	 * Show zero in decimal part
	 * `Object`
	 */
	ShowZeroInDecimalPart bool

	/**
	* Adding uppercase symbol of result number
	*
	 */
	AddUppercaseToFirstSymbol bool
}

func NewOptions(language words.Language, roundNumber int64,
	convertMinusSignToWord bool, showZeroInDecimalPart bool, addUppercaseTOFirstSymbol bool,
	showNumberParts, convertNumberToWords, showCurrency NumberPart,
) Options {
	return Options{
		Language:                  language,
		RoundNumber:               roundNumber,
		ConvertMinusSignToWord:    convertMinusSignToWord,
		ShowZeroInDecimalPart:     showZeroInDecimalPart,
		AddUppercaseToFirstSymbol: addUppercaseTOFirstSymbol,
		ShowCurrency:              showCurrency,
		ShowNumberParts:           showNumberParts,
		ConvertNumberToWords:      convertNumberToWords,
	}
}

func Default(language words.Language) Options {
	return Options{
		Language:                  language,
		RoundNumber:               3,
		ConvertMinusSignToWord:    true,
		ShowZeroInDecimalPart:     false,
		AddUppercaseToFirstSymbol: false,
		ShowCurrency:              NumberPart{true, true},
		ShowNumberParts:           NumberPart{true, true},
		ConvertNumberToWords:      NumberPart{true, true},
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
