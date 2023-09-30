package ru

import (
	core_constants "github.com/ThCompiler/go_game_constractor/pkg/convertor/core/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/core/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/core/words"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects/declension"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects/genders"
	"strings"
)

// ConvertNotLowestScaleToWords Возвращает название для дробной части если значащие(не 0) цифры начинаются с 1000
// и формирует соответствующие значения (тысячных, миллионных, ...)
func (rs *Russian) ConvertNotLowestScaleToWords(numberInfo words.NumberInfo, triplet objects.NumericDigitTriplet,
	tripletIndex int, isAloneScale bool, integerPart []objects.RuneDigitTriplet) string {
	if rs.words == nil {
		panic(ErrorLanguageNotLoaded)
	}

	numberDeclension, wordForm := getDeclensionAnsWordFormForFractionalPart(numberInfo.Declension,
		GetNumberForm(integerPart))

	// Собрать валюту в виде "двадцатипятитысячная".
	// Получить текущий класс для конвертирования.
	convertedScale := objects.StringDigitTriplet{
		Units:    "",
		Dozens:   rs.words.SlashNumberUnitPrefixes.Tens[triplet.Units],
		Hundreds: rs.words.SlashNumberUnitPrefixes.Hundreds[triplet.Hundreds],
	}

	if triplet.Dozens != 1 {
		convertedScale.Dozens = rs.words.SlashNumberUnitPrefixes.Dozens[triplet.Dozens]
		convertedScale.Units = rs.words.SlashNumberUnitPrefixes.Units[triplet.Units]
	}

	/* Если весь класс равен === 001
	и до него не было значений */
	if triplet == (objects.NumericDigitTriplet{Units: 1, Dozens: 0, Hundreds: 0}) && isAloneScale {
		// Получится "тысячная" вместо "однотысячная".
		convertedScale.Units = ""
	}

	if tripletIndex == 0 {
		return convertedScale.Hundreds + convertedScale.Dozens + convertedScale.Units + " "
	}

	// Получить корень названия класса числа ("тысяч")
	unitNameBeginnig := rs.words.FractionalUnit.FractionalUnitsBeginning[tripletIndex-1]
	if tripletIndex > len(rs.words.FractionalUnit.FractionalUnitsBeginning) {
		unitNameBeginnig = rs.words.UnitScalesNames.OtherBeginning[tripletIndex-2]
	}

	// Получить окончание названия класса числа с правильным падежом ("ная", "ной", "ных" и т.д.)
	unitNameEnding := rs.words.FractionalUnit.FractionalUnitEndings[numberDeclension][wordForm]
	// Добавить текст к общему результату
	return convertedScale.Hundreds + convertedScale.Dozens + convertedScale.Units +
		unitNameBeginnig + unitNameEnding + " "
}

// ConvertLowestScaleToWords Возвращает название для дробной части если значащие(не 0) цифры начинаются с 1000
// и формирует соответствующие значения (тысячных, миллионных, ...)
func (rs *Russian) ConvertLowestScaleToWords(numberInfo words.NumberInfo, triplet objects.NumericDigitTriplet,
	integerPart []objects.RuneDigitTriplet) string {
	if rs.words == nil {
		panic(ErrorLanguageNotLoaded)
	}

	numberDeclension, wordForm := getDeclensionAnsWordFormForFractionalPart(numberInfo.Declension,
		GetNumberForm(integerPart))

	res := ""

	/* Определить какой разряд в текущем классе последний
	(Любая последняя цифра, кроме нуля)
	0 - единицы, 1 - десятки, 2 - сотни */
	scaleForCommonConvert := triplet

	digitToConvert := triplet.Hundreds // Цифра для конвертирования

	// Получить объект для выбора формы последнего разряда
	lastDigitDeclensionsObject := rs.words.OrdinalNumbers.Hundreds

	switch {
	case triplet.Units != 0: // если разряд единиц
		scaleForCommonConvert.Units = 0
		digitToConvert = triplet.Units
		lastDigitDeclensionsObject = rs.words.OrdinalNumbers.Units
	case triplet.Dozens != 0: // если разряд десятков
		scaleForCommonConvert.Dozens = 0
		digitToConvert = triplet.Dozens
		lastDigitDeclensionsObject = rs.words.OrdinalNumbers.Dozens
	default: // если разряд сотен
		scaleForCommonConvert.Hundreds = 0
	}

	/* Получить класс без последнего разряда для конвертирования как обычного числа
	и если после этого в разряде десяток остается "1", то эту "1" тоже убрать,
	чтобы не отконвертировалось как 10-19. */
	if scaleForCommonConvert.Dozens == 1 {
		scaleForCommonConvert.Dozens = 0
	}

	// Если в классе остались цифры (не равен "000")
	if !scaleForCommonConvert.IsZeros() {
		// Конвертировать класс как обычное число
		res += rs.convertNotScaleNumbersToWords(scaleForCommonConvert) + " "
	}

	// Если в десятках цифра 1 (число 10-19)
	if triplet.Dozens == 1 {
		lastDigitDeclensionsObject = rs.words.OrdinalNumbers.Tens
		digitToConvert = triplet.Units
	}

	// Добавить текст к общему результату
	res += lastDigitDeclensionsObject[digitToConvert][genders.FEMALE][numberDeclension][wordForm] + " "

	return res
}

func (rs *Russian) convertNotScaleNumbersToWords(digits objects.NumericDigitTriplet) string {
	if rs.words == nil {
		panic(ErrorLanguageNotLoaded)
	}

	stringDigits := objects.StringDigitTriplet{Units: "", Dozens: "", Hundreds: ""}
	// Определить сотни
	stringDigits.Hundreds = convertDigitToWord(
		digits.Hundreds,
		rs.words.DigitWords.Hundreds,
		declension.NOMINATIVE,
		genders.FEMALE,
	)
	// Определить десятки и единицы
	// Если в разряде десятков стоит "1"
	if digits.Dozens == 1 {
		stringDigits.Dozens = convertDigitToWord(
			digits.Units,
			rs.words.DigitWords.Tens,
			declension.NOMINATIVE,
			genders.FEMALE,
		)
	} else { // Если в разряде десятков стоит не "1"
		stringDigits.Dozens = convertDigitToWord(
			digits.Dozens,
			rs.words.DigitWords.Dozens,
			declension.NOMINATIVE,
			genders.FEMALE,
		)

		stringDigits.Units = convertDigitToWord(
			digits.Units,
			rs.words.DigitWords.Units,
			declension.NOMINATIVE,
			genders.FEMALE,
		)
	}

	// Убрать ненужный "ноль"
	if digits.Units == 0 && (digits.Hundreds > 0 || digits.Dozens > 0) {
		stringDigits.Units = ""
	}

	return strings.TrimSpace(stringDigits.Hundreds) + " " +
		strings.TrimSpace(stringDigits.Dozens) + " " +
		strings.TrimSpace(stringDigits.Units)
}

// ConvertZeroToWordsForFractionalNumber Возвращает словесную форму числа для знаменателя числа
func (rs *Russian) ConvertZeroToWordsForFractionalNumber(numberInfo words.NumberInfo,
	integerPartTriplets []objects.RuneDigitTriplet) string {
	if rs.words == nil {
		panic(ErrorLanguageNotLoaded)
	}

	decl, wordForm := getDeclensionAnsWordFormForFractionalPart(numberInfo.Declension, GetNumberForm(integerPartTriplets))
	return rs.words.OrdinalNumbers.Units[0][genders.FEMALE][decl][wordForm]
}

// GetEndingOfDecimalNumberForFractionalPart Возвращает значение описывающие размерность десятичной
// части числа (тысячных, десятых, сотых и т.д.)
func (rs *Russian) GetEndingOfDecimalNumberForFractionalPart(countDigits int, lastDigit objects.Digit,
	decl words.Declension) string {

	numberDeclension, wordForm := getDeclensionAnsWordFormForFractionalPart(decl, GetNumberFormByDigit(lastDigit))

	if lastDigit == 0 {
		numberDeclension = declension.GENITIVE
		wordForm = constants.PLURAL_WORD
	}

	if countDigits <= 0 {
		return rs.words.FractionalUnit.FractionalUnitsDeclensions.Tens[numberDeclension][wordForm]
	}

	if countDigits == 1 {
		return rs.words.FractionalUnit.FractionalUnitsDeclensions.Hundreds[numberDeclension][wordForm]
	}

	// Определить класс числа
	// (0 - единицы, 1 - тысячи, 2 - миллионы и т.д.).
	numberScale := countDigits / core_constants.TripletSize

	// Получить разряд цифры в текущем классе
	// (0 - единицы, 1 - десятки, 2 - сотни).
	digitIndexInScale := countDigits - numberScale*core_constants.TripletSize

	// Получить корень названия класса числа
	unitNameBegin := rs.words.FractionalUnit.FractionalUnitsBeginning[min(
		len(rs.words.FractionalUnit.FractionalUnitsBeginning)-1,
		countDigits-2,
	)]

	// Получить приставку к числу
	unitNamePrefix := rs.words.FractionalUnit.FractionalUnitPrefixes[digitIndexInScale+1]

	// Составить объект с падежами
	return unitNamePrefix + unitNameBegin + rs.words.FractionalUnit.FractionalUnitEndings[numberDeclension][wordForm]
}

// CorrectNumberInfoForFractionalTriplets преобразует существующую информацию о числе в требуемую для описания
// чисел в знаменателе до последней значащей тройки (т.е. "123 231 123 000", параметры для описания 123 и 231)
func (rs *Russian) CorrectNumberInfoForFractionalTriplets(numberInfo words.NumberInfo) words.NumberInfo {
	if rs.words == nil {
		panic(ErrorLanguageNotLoaded)
	}

	return words.NumberInfo{CurrencyName: femaleCurrency, NumberType: numberInfo.NumberType,
		Declension: declension.NOMINATIVE}
}
