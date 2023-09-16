package functions

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/currency"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/num2words/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words/declension"
)

func GetCurrencyWord(
	currencyObject currency.CustomCurrency, numberType constants.NumberType,
	numberForm objects.NumberForm, lowestTripletIsZero bool,
	isNumber bool, currentDeclension declension.Declension,
) string {
	// Если не "именительный" и не "винительный" падеж,
	// то всё что не один это множественное число
	wordForm := objects.SINGULAR_WORD
	if numberForm != objects.FIRST_FORM {
		wordForm = objects.PLURAL_WORD
	}

	// Если падеж "именительный" или "винительный" и множественное число
	if (currentDeclension == declension.NOMINATIVE ||
		currentDeclension == declension.ACCUSATIVE) && numberForm != objects.FIRST_FORM {
		wordForm = getWordForm(isNumber, numberForm)
		// Использовать родительный падеж.
		currentDeclension = declension.GENITIVE
	}

	// Если последний класс числа равен "000"
	if lowestTripletIsZero {
		wordForm = objects.PLURAL_WORD
		// Всегда родительный падеж и множественное число
		currentDeclension = declension.GENITIVE
	}

	if numberType == constants.FRACTIONAL_NUMBER {
		return currencyObject.FractionalPartNameDeclensions[currentDeclension][wordForm]
	}
	return currencyObject.IntegerPartNameDeclensions[currentDeclension][wordForm]
}

func getWordForm(isNumber bool, numberForm objects.NumberForm) objects.WordForm {
	if isNumber {
		return objects.PLURAL_WORD
	}

	if numberForm == objects.SECOND_FORM {
		return objects.SINGULAR_WORD
	}

	return objects.PLURAL_WORD
}
