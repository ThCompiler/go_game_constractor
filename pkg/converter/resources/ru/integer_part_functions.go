package ru

import (
	core_constants "github.com/ThCompiler/go_game_constractor/pkg/converter/core/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/convert"
	core_objects "github.com/ThCompiler/go_game_constractor/pkg/converter/core/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/objects/declension"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/objects/genders"
)

// GetZeroAsWordsForIntegerPart Возвращает словесную форму числа состоящего из нулей
func (rs *Russian) GetZeroAsWordsForIntegerPart() string {
	if rs.words == nil {
		panic(ErrorLanguageNotLoaded)
	}

	return convertDigitToWord(0, rs.words.DigitWords.Units, rs.declension, genders.FEMALE)
}

// ConvertTripletToWords Преобразует тройку в словесную форму
func (rs *Russian) ConvertTripletToWords(numberType core_constants.NumberType,
	digits core_objects.NumericDigitTriplet, scale int) core_objects.StringDigitTriplet {
	if rs.words == nil {
		panic(ErrorLanguageNotLoaded)
	}

	unitNounGender := rs.GetCurrency().CurrencyNounGender.IntegerPart
	decl := rs.declension
	if numberType == core_constants.FRACTIONAL_NUMBER {
		unitNounGender = genders.FEMALE
		decl = declension.NOMINATIVE
	}

	stringDigits := core_objects.StringDigitTriplet{Units: "", Dozens: "", Hundreds: ""}

	gender := genders.MALE
	if scale == convert.ThousandScale {
		// Если текущий класс - тысячи
		gender = genders.FEMALE
	} else if scale == convert.UnitsScale {
		// Если текущий класс - единицы
		gender = unitNounGender
	}

	// Определить сотни
	stringDigits.Hundreds = convertDigitToWord(
		digits.Hundreds,
		rs.words.DigitWords.Hundreds,
		decl,
		gender,
	)
	// Определить десятки и единицы
	// Если в разряде десятков стоит "1"
	if digits.Dozens == 1 {
		stringDigits.Dozens = convertDigitToWord(
			digits.Units,
			rs.words.DigitWords.Tens,
			decl,
			gender,
		)
	} else { // Если в разряде десятков стоит не "1"
		stringDigits.Dozens = convertDigitToWord(
			digits.Dozens,
			rs.words.DigitWords.Dozens,
			decl,
			gender,
		)

		stringDigits.Units = convertDigitToWord(
			digits.Units,
			rs.words.DigitWords.Units,
			decl,
			gender,
		)
	}

	return stringDigits
}

// GetWordScaleName Возвращает название текущего уровня тройки числа (тысячи, миллионы, ...)
func (rs *Russian) GetWordScaleName(numberType core_constants.NumberType,
	scale int, digits core_objects.NumericDigitTriplet) string {
	if rs.words == nil {
		panic(ErrorLanguageNotLoaded)
	}

	if scale == convert.UnitsScale {
		// Класс единиц
		// Для них название не отображается.
		return ""
	}

	scaleDeclension := rs.declension
	if numberType == core_constants.FRACTIONAL_NUMBER {
		scaleDeclension = declension.NOMINATIVE
	}
	wordForm := constants.PLURAL_WORD
	numberForm := GetNumberFormByTriplet(digits)

	if numberForm == constants.FIRST_FORM {
		wordForm = constants.SINGULAR_WORD
	}

	// Если падеж "именительный" или "винительный" и множественное число
	if (scaleDeclension == declension.NOMINATIVE ||
		scaleDeclension == declension.ACCUSATIVE) &&
		numberForm >= constants.SECOND_FORM {
		// Для множественного числа именительного падежа используется родительный падеж.
		scaleDeclension = declension.GENITIVE

		if numberForm == constants.SECOND_FORM {
			wordForm = constants.SINGULAR_WORD
		}
	}

	// Класс тысяч
	if scale == convert.ThousandScale {
		return rs.words.UnitScalesNames.Thousands[scaleDeclension][wordForm]
	}

	// Остальные классы
	base := rs.words.UnitScalesNames.OtherBeginning[scale-3]
	ending := rs.words.UnitScalesNames.OtherEnding[scaleDeclension][wordForm]

	return base + ending
}
