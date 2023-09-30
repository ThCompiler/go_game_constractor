package ru

import (
	core_constants "github.com/ThCompiler/go_game_constractor/pkg/convertor/core/constants"
	core_objects "github.com/ThCompiler/go_game_constractor/pkg/convertor/core/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/core/words"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects/declension"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects/genders"
)

const (
	UnitsScale    = 1
	ThousandScale = 2
)

// ConvertZeroToWordsForIntegerPart Возвращает словесную форму числа состоящего из нулей
func (rs *Russian) ConvertZeroToWordsForIntegerPart(declension words.Declension) string {
	if rs.words == nil {
		panic(ErrorLanguageNotLoaded)
	}

	return convertDigitToWord(0, rs.words.DigitWords.Units, declension, genders.FEMALE)
}

// ConvertTripletToWords Преобразует тройку в словесную форму
func (rs *Russian) ConvertTripletToWords(numberInfo words.NumberInfo,
	digits core_objects.NumericDigitTriplet, scale int) core_objects.StringDigitTriplet {
	if rs.words == nil {
		panic(ErrorLanguageNotLoaded)
	}

	unitNounGender := rs.GetCurrencyByName(numberInfo.CurrencyName).CurrencyNounGender.IntegerPart
	if numberInfo.NumberType == core_constants.FRACTIONAL_NUMBER {
		unitNounGender = genders.FEMALE
	}

	stringDigits := core_objects.StringDigitTriplet{Units: "", Dozens: "", Hundreds: ""}

	gender := genders.MALE
	if scale == ThousandScale {
		// Если текущий класс - тысячи
		gender = genders.FEMALE
	} else if scale == UnitsScale {
		// Если текущий класс - единицы
		gender = unitNounGender
	}

	// Определить сотни
	stringDigits.Hundreds = convertDigitToWord(
		digits.Hundreds,
		rs.words.DigitWords.Hundreds,
		numberInfo.Declension,
		gender,
	)
	// Определить десятки и единицы
	// Если в разряде десятков стоит "1"
	if digits.Dozens == 1 {
		stringDigits.Dozens = convertDigitToWord(
			digits.Units,
			rs.words.DigitWords.Tens,
			numberInfo.Declension,
			gender,
		)
	} else { // Если в разряде десятков стоит не "1"
		stringDigits.Dozens = convertDigitToWord(
			digits.Dozens,
			rs.words.DigitWords.Dozens,
			numberInfo.Declension,
			gender,
		)

		stringDigits.Units = convertDigitToWord(
			digits.Units,
			rs.words.DigitWords.Units,
			numberInfo.Declension,
			gender,
		)
	}

	return stringDigits
}

// GetWordScaleName Возвращает название текущего уровня тройки числа (тысячи, миллионы, ...)
func (rs *Russian) GetWordScaleName(scale int, numberInfo words.NumberInfo, digits core_objects.NumericDigitTriplet) string {
	if rs.words == nil {
		panic(ErrorLanguageNotLoaded)
	}

	if scale == 1 {
		// Класс единиц
		// Для них название не отображается.
		return ""
	}

	scaleDeclension := numberInfo.Declension
	wordForm := constants.PLURAL_WORD
	numberForm := GetNumberFormByTriplet(digits)

	if numberForm == constants.FIRST_FORM {
		wordForm = constants.SINGULAR_WORD
	}

	// Если падеж "именительный" или "винительный" и множественное число
	if (numberInfo.Declension == declension.NOMINATIVE ||
		numberInfo.Declension == declension.ACCUSATIVE) &&
		numberForm >= constants.SECOND_FORM {
		// Для множественного числа именительного падежа используется родительный падеж.
		scaleDeclension = declension.GENITIVE
		wordForm = constants.SINGULAR_WORD

		if numberForm == constants.SECOND_FORM {
			wordForm = constants.PLURAL_WORD
		}
	}

	// Класс тысяч
	if scale == 2 {
		return rs.words.UnitScalesNames.Thousands[scaleDeclension][wordForm]
	}

	// Остальные классы
	base := rs.words.UnitScalesNames.OtherBeginning[scale-2]
	ending := rs.words.UnitScalesNames.OtherEnding[scaleDeclension][wordForm]

	return base + ending
}
