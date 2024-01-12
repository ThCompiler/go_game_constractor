package ru

import (
	core_constants "github.com/ThCompiler/go_game_constractor/pkg/converter/core/constants"
	core_objects "github.com/ThCompiler/go_game_constractor/pkg/converter/core/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/words"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/objects/declension"
)

// GetCurrencyAsWord Возвращает значение валюты в словах для указанной части числа
func (rs *Russian) GetCurrencyAsWord(numberPart core_constants.NumberPart,
	numberAsTriplets []core_objects.RuneDigitTriplet) string {
	if rs.words == nil {
		panic(ErrorLanguageNotLoaded)
	}
	// Если не "именительный" и не "винительный" падеж,
	// то всё что не один это множественное число

	numberForm := GetNumberForm(numberAsTriplets)
	currencyDeclensions := rs.declension

	wordForm := constants.SINGULAR_WORD
	if numberForm != constants.FIRST_FORM {
		wordForm = constants.PLURAL_WORD
	}

	// Если падеж "именительный" или "винительный" и множественное число
	if (currencyDeclensions == declension.NOMINATIVE ||
		currencyDeclensions == declension.ACCUSATIVE) && numberForm != constants.FIRST_FORM {
		wordForm = getCurrencyIntegerPartWordForm(rs.currencyName == words.NUMBER, numberForm)
		// Использовать родительный падеж.
		currencyDeclensions = declension.GENITIVE
	}

	// Если последний класс числа равен "000"
	if numberAsTriplets[len(numberAsTriplets)-1].IsZeros() {
		wordForm = constants.PLURAL_WORD
		// Всегда родительный падеж и множественное число
		currencyDeclensions = declension.GENITIVE
	}

	if numberPart == core_constants.FRACTIONAL_PART {
		return rs.GetCurrency().FractionalPartNameDeclensions[currencyDeclensions][wordForm]
	}
	return rs.GetCurrency().DecimalCurrencyNameDeclensions[currencyDeclensions][wordForm]
}

// GetCurrencyForFractionalNumber Возвращает значение валюты в случае если передана дробь
func (rs *Russian) GetCurrencyForFractionalNumber() string {
	if rs.words == nil {
		panic(ErrorLanguageNotLoaded)
	}
	return rs.GetCurrency().DecimalCurrencyNameDeclensions[declension.GENITIVE][constants.SINGULAR_WORD]
}
