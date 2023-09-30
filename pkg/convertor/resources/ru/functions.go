package ru

import (
	core_objects "github.com/ThCompiler/go_game_constractor/pkg/convertor/core/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/core/words"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects/declension"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects/genders"
)

func GetNumberForm(triplets []core_objects.RuneDigitTriplet) constants.NumberForm {
	if len(triplets) == 0 {
		return constants.THIRD_FORM
	}

	return GetNumberFormByTriplet(triplets[len(triplets)-1].ToNumeric())
}

func GetNumberFormByTriplet(triplet core_objects.NumericDigitTriplet) constants.NumberForm {
	return GetNumberFormByDigit(triplet.Units)
}

func GetNumberFormByDigit(digit core_objects.Digit) constants.NumberForm {
	if digit == 1 {
		return constants.FIRST_FORM
	}

	if digit == 0 || digit > 4 {
		return constants.THIRD_FORM
	}

	return constants.SECOND_FORM
}

func convertDigitToWord(digit core_objects.Digit, digitWords objects.DeclensionNumbers,
	declension words.Declension, gender genders.Gender,
) string {
	declensionValues := digitWords[declension]
	word := declensionValues[digit]

	if word.WithGender() {
		return word.GetWordsByGender()[gender]
	}

	return word.GetWord()
}

func getWordForm(isNumber bool, numberForm constants.NumberForm) constants.WordForm {
	if isNumber {
		return constants.PLURAL_WORD
	}

	if numberForm == constants.SECOND_FORM {
		return constants.SINGULAR_WORD
	}

	return constants.PLURAL_WORD
}

func getDeclensionAnsWordFormForFractionalPart(decl words.Declension,
	numberForm constants.NumberForm) (words.Declension, constants.WordForm) {
	wordForm := constants.SINGULAR_WORD

	if numberForm != constants.FIRST_FORM {
		wordForm = constants.PLURAL_WORD
	}

	// Если падеж "именительный" или "винительный" и множественное число
	if numberForm != constants.FIRST_FORM && (decl == declension.NOMINATIVE || decl == declension.ACCUSATIVE) {
		decl = declension.GENITIVE
	}

	return decl, wordForm
}
