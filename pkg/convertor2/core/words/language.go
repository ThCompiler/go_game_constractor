package words

import (
	currency2 "github.com/ThCompiler/go_game_constractor/pkg/convertor/currency"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/currency"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/objects"
)

type Gender string
type Declension string
type LanguageName string

type Language interface {
	GetMinusString() string
	GetCurrency(currency currency.Name) currency2.CustomCurrency
	GetCurrencyAsWord(currencyName currency.Name, numberType constants.NumberType,
		numberAsTriplets []objects.RuneDigitTriplet, declension Declension) string
	GetCurrencyForFractionalNumber() string
	ConvertIntegerPartTripletsToWords(currencyName currency.Name, convertNumberAsTriplets []objects.RuneDigitTriplet,
		numberType constants.NumberType, declension Declension) string
	ConvertFractionalPartTripletsToWords(currencyName currency.Name, convertNumberAsTriplets []objects.RuneDigitTriplet,
		numberType constants.NumberType, declension Declension) string
	GetEndingOfDecimalNumberForFractionalPart(countDigits int, lastDigit objects.Digit,
		declension Declension) string
}
