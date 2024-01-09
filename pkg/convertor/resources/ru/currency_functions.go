package ru

import (
	core_constants "github.com/ThCompiler/go_game_constractor/pkg/convertor/core/constants"
	core_objects "github.com/ThCompiler/go_game_constractor/pkg/convertor/core/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/core/words"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects/declension"
)

// GetCurrencyAsWord Возвращает значение валюты в словах для указанной части числа
func (rs *Russian) GetCurrencyAsWord(numberInfo words.NumberInfo, numberAsTriplets []core_objects.RuneDigitTriplet) string {
	// Если не "именительный" и не "винительный" падеж,
	// то всё что не один это множественное число

	numberForm := GetNumberForm(numberAsTriplets)
	currencyDeclensions := numberInfo.Declension

	wordForm := constants.SINGULAR_WORD
	if numberForm != constants.FIRST_FORM {
		wordForm = constants.PLURAL_WORD
	}

	// Если падеж "именительный" или "винительный" и множественное число
	if (currencyDeclensions == declension.NOMINATIVE ||
		currencyDeclensions == declension.ACCUSATIVE) && numberForm != constants.FIRST_FORM {
		wordForm = getCurrencyIntegerPartWordForm(numberInfo.CurrencyName == words.NUMBER, numberForm)
		// Использовать родительный падеж.
		currencyDeclensions = declension.GENITIVE
	}

	// Если последний класс числа равен "000"
	if numberAsTriplets[len(numberAsTriplets)-1].IsZeros() {
		wordForm = constants.PLURAL_WORD
		// Всегда родительный падеж и множественное число
		currencyDeclensions = declension.GENITIVE
	}

	if numberInfo.NumberType == core_constants.FRACTIONAL_NUMBER {
		return rs.GetCurrencyByName(numberInfo.CurrencyName).FractionalPartNameDeclensions[currencyDeclensions][wordForm]
	}
	return rs.GetCurrencyByName(numberInfo.CurrencyName).DecimalCurrencyNameDeclensions[currencyDeclensions][wordForm]
}

// GetCurrencyForFractionalNumber Возвращает значение валюты в случае если передана дробь
func (rs *Russian) GetCurrencyForFractionalNumber(currencyName words.CurrencyName) string {
	return rs.GetCurrencyByName(currencyName).DecimalCurrencyNameDeclensions[declension.GENITIVE][constants.SINGULAR_WORD]
}
