package ru

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/words"
)

// ConvertZeroToWordsForIntegerPart Возвращает словесную форму числа состоящего из нулей
func (rs *Russian) ConvertZeroToWordsForIntegerPart(declension words.Declension) string {
	return ""
}

// ConvertTripletToWords Преобразует тройку в словесную форму
func (rs *Russian) ConvertTripletToWords(numberInfo words.NumberInfo,
	digits objects.NumericDigitTriplet, scale int) objects.StringDigitTriplet {
	return objects.StringDigitTriplet{}
}

// GetWordScaleName Возвращает название текущего уровня тройки числа (тысячи, миллионы, ...)
func (rs *Russian) GetWordScaleName(scale int, numberInfo words.NumberInfo, digits objects.NumericDigitTriplet) string {
	return ""
}
