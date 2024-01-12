package words

import (
	core_constants "github.com/ThCompiler/go_game_constractor/pkg/converter/core/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/objects"
)

type LanguageName string
type CurrencyName string

const (
	NUMBER = CurrencyName("number")
)

type Language interface {
	// GetMinusString Возвращает значение знака минус в словах
	GetMinusString() string

	// GetCurrencyAsWord Возвращает значение валюты в словах для указанной части числа
	GetCurrencyAsWord(numberPart core_constants.NumberPart, numberAsTriplets []objects.RuneDigitTriplet) string

	// GetCurrencyForFractionalNumber Возвращает значение валюты в случае если передана дробь
	GetCurrencyForFractionalNumber() string

	// For Integer Part

	// GetZeroAsWordsForIntegerPart Возвращает словесную форму числа состоящего из нулей
	GetZeroAsWordsForIntegerPart() string

	// ConvertTripletToWords Преобразует тройку в словесную форму
	ConvertTripletToWords(numberType core_constants.NumberType, digits objects.NumericDigitTriplet,
		scale int) objects.StringDigitTriplet

	// GetWordScaleName Возвращает название текущего уровня тройки числа (тысячи, миллионы, ...)
	GetWordScaleName(numberType core_constants.NumberType, scale int, digits objects.NumericDigitTriplet) string

	// For Fractional Part

	// ConvertNotLowestScaleToWords Возвращает название для дробной части если значащие(не 0) цифры начинаются с 1000
	// и формирует соответствующие значения (тысячных, миллионных, ...)
	ConvertNotLowestScaleToWords(triplet objects.NumericDigitTriplet, tripletIndex int, isAloneScale bool,
		integerPart []objects.RuneDigitTriplet) string

	// ConvertLowestScaleToWords Возвращает название для дробной части если значащие(не 0) цифры начинаются с 1000
	// и формирует соответствующие значения (тысячных, миллионных, ...)
	ConvertLowestScaleToWords(triplet objects.NumericDigitTriplet,
		integerPart []objects.RuneDigitTriplet) string

	// GetZeroAsWordsForFractionalNumber Возвращает словесную форму нуля для знаменателя числа
	GetZeroAsWordsForFractionalNumber(integerPartTriplets []objects.RuneDigitTriplet) string

	// GetEndingOfDecimalNumberForFractionalPart Возвращает значение описывающие размерность десятичной
	// части числа (тысячных, десятых, сотых и т.д.)
	GetEndingOfDecimalNumberForFractionalPart(countDigits int, lastDigit objects.Digit) string

	// IsCurrency Сообщает необходимо ли перевести число с указанием валюты
	IsCurrency() bool

	// IsNumber Сообщает необходимо ли перевести число как число без указание валюты
	IsNumber() bool
}
