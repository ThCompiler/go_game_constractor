package words

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/objects"
)

type Declension string
type LanguageName string
type CurrencyName string

const (
	NUMBER = CurrencyName("number")
)

type NumberInfo struct {
	Declension   Declension
	NumberType   constants.NumberType
	CurrencyName CurrencyName
}

type Language interface {
	// GetMinusString Возвращает значение знака минус в словах
	GetMinusString() string

	// GetCurrencyAsWord Возвращает значение валюты в словах для указанной части числа
	GetCurrencyAsWord(numberInfo NumberInfo, numberAsTriplets []objects.RuneDigitTriplet) string

	// GetCurrencyForFractionalNumber Возвращает значение валюты в случае если передана дробь
	GetCurrencyForFractionalNumber(currencyName CurrencyName) string

	// For Integer Part

	// ConvertZeroToWordsForIntegerPart Возвращает словесную форму числа состоящего из нулей
	ConvertZeroToWordsForIntegerPart(declension Declension) string

	// ConvertTripletToWords Преобразует тройку в словесную форму
	ConvertTripletToWords(numberInfo NumberInfo, digits objects.NumericDigitTriplet, scale int) objects.StringDigitTriplet

	// GetWordScaleName Возвращает название текущего уровня тройки числа (тысячи, миллионы, ...)
	GetWordScaleName(scale int, numberInfo NumberInfo, digits objects.NumericDigitTriplet) string

	// For Fractional Part

	// ConvertNotLowestScaleToWords Возвращает название для дробной части если значащие(не 0) цифры начинаются с 1000
	// и формирует соответствующие значения (тысячных, миллионных, ...)
	ConvertNotLowestScaleToWords(numberInfo NumberInfo, triplet objects.NumericDigitTriplet, tripletIndex int, isAloneScale bool,
		integerPart []objects.RuneDigitTriplet) string

	// ConvertLowestScaleToWords Возвращает название для дробной части если значащие(не 0) цифры начинаются с 1000
	// и формирует соответствующие значения (тысячных, миллионных, ...)
	ConvertLowestScaleToWords(numberInfo NumberInfo, triplet objects.NumericDigitTriplet,
		integerPart []objects.RuneDigitTriplet) string

	// ConvertZeroToWordsForFractionalNumber Возвращает словесную форму числа для знаменателя числа
	ConvertZeroToWordsForFractionalNumber(numberInfo NumberInfo, integerPartTriplets []objects.RuneDigitTriplet) string

	// GetEndingOfDecimalNumberForFractionalPart Возвращает значение описывающие размерность десятичной
	// части числа (тысячных, десятых, сотых и т.д.)
	GetEndingOfDecimalNumberForFractionalPart(countDigits int, lastDigit objects.Digit,
		declension Declension) string

	// CorrectNumberInfoForFractionalTriplets преобразует существующую информацию о числе в требуемую для описания
	// чисел в знаменателе до последней значащей тройки (т.е. "123 231 123 000", параметры для описания 123 и 231)
	CorrectNumberInfoForFractionalTriplets(numberInfo NumberInfo) NumberInfo
}
