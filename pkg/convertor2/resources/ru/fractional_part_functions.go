package ru

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/words"
)

// ConvertNotLowestScaleToWords Возвращает название для дробной части если значащие(не 0) цифры начинаются с 1000
// и формирует соответствующие значения (тысячных, миллионных, ...)
func (rs *Russian) ConvertNotLowestScaleToWords(numberInfo words.NumberInfo, triplet objects.NumericDigitTriplet,
	tripletIndex int, isAloneScale bool, integerPart []objects.RuneDigitTriplet) string {
	return ""
}

// ConvertLowestScaleToWords Возвращает название для дробной части если значащие(не 0) цифры начинаются с 1000
// и формирует соответствующие значения (тысячных, миллионных, ...)
func (rs *Russian) ConvertLowestScaleToWords(numberInfo words.NumberInfo, triplet objects.NumericDigitTriplet,
	integerPart []objects.RuneDigitTriplet) string {
	return ""
}

// ConvertZeroToWordsForFractionalNumber Возвращает словесную форму числа для знаменателя числа
func (rs *Russian) ConvertZeroToWordsForFractionalNumber(numberInfo words.NumberInfo,
	integerPartTriplets []objects.RuneDigitTriplet) string {
	return ""
}

// GetEndingOfDecimalNumberForFractionalPart Возвращает значение описывающие размерность десятичной
// части числа (тысячных, десятых, сотых и т.д.)
func (rs *Russian) GetEndingOfDecimalNumberForFractionalPart(countDigits int, lastDigit objects.Digit,
	declension words.Declension) string {
	return ""
}

// CorrectNumberInfoForFractionalTriplets преобразует существующую информацию о числе в требуемую для описания
// чисел в знаменателе до последней значащей тройки (т.е. "123 231 123 000", параметры для описания 123 и 231)
func (rs *Russian) CorrectNumberInfoForFractionalTriplets(numberInfo words.NumberInfo) words.NumberInfo {
	return words.NumberInfo{}
}
