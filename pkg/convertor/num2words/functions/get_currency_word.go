package functions

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/currency"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/num2words/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words/declension"
)

func GetCurrencyWord(
	currencyObject currency.CustomCurrency, numberPart constants.NumberType,
	scaleNameForm objects.ScaleForm, lastScaleIsZero bool,
	curc currency.Currency, decl declension.Declension,
) string {
	declensionsObject := currencyObject.FractionalPartNameDeclensions
	if numberPart == constants.DecimalNumber {
		declensionsObject = currencyObject.CurrencyNameDeclensions
	}

	scaleForm := 1
	if scaleNameForm == 0 {
		scaleForm = 0
	}

	currentDeclension := decl

	// Если падеж "именительный" или "винительный" и множественное число
	if (decl == declension.NOMINATIVE || decl == declension.ACCUSATIVE) && scaleNameForm >= 1 {
		scaleForm = getScaleForm(curc, scaleNameForm)
		// Использовать родительный падеж.
		currentDeclension = declension.GENITIVE
	}
	// Если последний класс числа равен "000"
	if lastScaleIsZero {
		scaleForm = 1
		// Всегда родительный падеж и множественное число
		currentDeclension = declension.GENITIVE
	}

	return declensionsObject[currentDeclension][scaleForm]
}

func getScaleForm(curc currency.Currency, scaleNameForm objects.ScaleForm) int {
	// Если валюта указана как "number"
	if curc == currency.NUMBER {
		return 1
	}

	if scaleNameForm == 1 {
		return 0
	}

	return 1
}
