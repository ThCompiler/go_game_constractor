package functions

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/num2words/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words/declension"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words/genders"
	"strings"
)

func ConvertEachScaleToWordsSlash(
	numberScalesArray []objects.RuneDigitTriplet,
	scaleForm objects.ScaleForm,
	decl declension.Declension,
) string {
	if len(numberScalesArray) < 1 {
		return ""
	}
	convertedResult := ""

	// Получить последний разряд (в виде текста) в правильном падеже
	numberDecl, numberScaleForm := selectDeclensionsParamsByDeclension(decl, scaleForm != 0)

	// Удалить лишние нули в начале числа
	updatedNumberScalesArray := removeEmptyScalesFromBeginning(numberScalesArray)
	updatedNumberScalesArrayLen := len(updatedNumberScalesArray)

	/* Определить индекс класса, который является последним.
	   После него могут быть только классы с "000".
	   0 - единицы, 1 - тысячи, 2 - миллионы и т.д. */
	lastScaleWihNumber := indexOfLastNotNullScalesByEnd(updatedNumberScalesArray)
	// Индекс последнего класса в массиве.
	lastScaleWihNumberIndex := updatedNumberScalesArrayLen - lastScaleWihNumber - 1

	// Если нет ни одного не пустого класса
	if lastScaleWihNumber == -1 {
		// Вернуть ноль
		zeroValuesObject := words.WordConstants.N2w.OrdinalNumbers.Units[0]
		return zeroValuesObject[genders.FEMALE][numberDecl][numberScaleForm]
	}

	/* Если есть не пустые классы до последнего не пустого класса,
	   то конвертировать как обычное число */
	if lastScaleWihNumberIndex > 0 {
		// Получить массив классов, в котором последний класс будет пустым.
		numberScalesArrForCommonConvert := make([]objects.RuneDigitTriplet, len(updatedNumberScalesArray))
		copy(numberScalesArrForCommonConvert, updatedNumberScalesArray)
		numberScalesArrForCommonConvert[lastScaleWihNumberIndex] = objects.RuneDigitTriplet{
			Units:    '0',
			Dozens:   '0',
			Hundreds: '0',
		}

		// Конвертировать классы как обычное число
		convertedResult += ConvertEachScaleToWords(
			numberScalesArrForCommonConvert,
			genders.FEMALE,
			declension.NOMINATIVE,
		).Result + " "
	}

	// Если последний класс для конвертирования - тысячи или больше
	if lastScaleWihNumber >= 1 {
		convertedResult += convertOtherScaleToWordsSlash(updatedNumberScalesArray[lastScaleWihNumberIndex].ToNumeric(),
			lastScaleWihNumber, updatedNumberScalesArrayLen-1 == lastScaleWihNumber, numberDecl, numberScaleForm)
	}
	// Если последний класс для конвертирования - единицы
	if lastScaleWihNumber == 0 {
		convertedResult +=
			convertLastScaleToWordsSlash(updatedNumberScalesArray[lastScaleWihNumberIndex].ToNumeric(), numberDecl, numberScaleForm)
	}

	return strings.TrimSpace(convertedResult)
}

func convertOtherScaleToWordsSlash(scaleToConvert objects.NumericDigitTriplet, scaleNumber int, isAloneScale bool,
	numberDeclension declension.Declension, numberScaleForm objects.ScaleForm) string {
	// Собрать валюту в виде "двадцатипятитысячная".
	// Получить текущий класс для конвертирования.

	convertedScale := objects.StringDigitTriplet{
		Units:    "",
		Dozens:   words.WordConstants.N2w.SlashNumberUnitPrefixes.Tens[scaleToConvert.Units],
		Hundreds: words.WordConstants.N2w.SlashNumberUnitPrefixes.Hundreds[scaleToConvert.Hundreds],
	}
	if scaleToConvert.Dozens != 1 {
		convertedScale.Dozens =
			words.WordConstants.N2w.SlashNumberUnitPrefixes.Dozens[scaleToConvert.Dozens]
		convertedScale.Units = words.WordConstants.N2w.SlashNumberUnitPrefixes.Units[scaleToConvert.Units]
	}
	/* Если весь класс равен === 001
	   и до него не было значений */
	if scaleToConvert == (objects.NumericDigitTriplet{1, 0, 0}) && isAloneScale {
		// Получится "тысячная" вместо "однотысячная".
		convertedScale.Units = ""
	}

	if scaleNumber != 0 {
		// Получить корень названия класса числа ("тысяч")
		unitNameBeginnig := words.WordConstants.N2w.FractionalUnit.FractionalUnitsBeginning[scaleNumber-1]
		if scaleNumber > len(words.WordConstants.N2w.FractionalUnit.FractionalUnitsBeginning) {
			unitNameBeginnig = words.WordConstants.N2w.UnitScalesNames.OtherBeginning[scaleNumber-2]
		}

		// Получить окончание названия класса числа с правильным падежом ("ная", "ной", "ных" и т.д.)
		unitNameEnding :=
			words.WordConstants.N2w.FractionalUnit.FractionalUnitEndings[numberDeclension][numberScaleForm]
		// Добавить текст к общему результату
		return convertedScale.Hundreds + convertedScale.Dozens + convertedScale.Units +
			unitNameBeginnig + unitNameEnding + " "
	}
	return convertedScale.Hundreds + convertedScale.Dozens + convertedScale.Units + " "
}

func convertLastScaleToWordsSlash(scaleToConvert objects.NumericDigitTriplet,
	numberDeclension declension.Declension, numberScaleForm objects.ScaleForm) (res string) {
	res = ""

	/* Определить какой разряд в текущем классе последний
	   (Любая последняя цифра, кроме нуля)
	   0 - единицы, 1 - десятки, 2 - сотни */
	scaleForCommonConvert := scaleToConvert

	digitToConvert := scaleToConvert.Hundreds // Цифра для конвертирования

	// Получить объект для выбора формы последнего разряда
	lastDigitDeclensionsObject := words.WordConstants.N2w.OrdinalNumbers.Hundreds
	if scaleToConvert.Units != 0 { // если разряд сотен
		scaleForCommonConvert.Units = 0
		digitToConvert = scaleToConvert.Units
		lastDigitDeclensionsObject = words.WordConstants.N2w.OrdinalNumbers.Units
	} else if scaleToConvert.Dozens != 0 { // если разряд десятков
		scaleForCommonConvert.Dozens = 0
		digitToConvert = scaleToConvert.Dozens
		lastDigitDeclensionsObject = words.WordConstants.N2w.OrdinalNumbers.Dozens
	} else {
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
		res += ConvertEachScaleToWords(
			[]objects.RuneDigitTriplet{scaleForCommonConvert.ToRune()},
			genders.FEMALE,
			declension.NOMINATIVE).Result + " "
	}

	// Если в десятках цифра 1 (число 10-19)
	if scaleToConvert.Dozens == 1 {
		lastDigitDeclensionsObject = words.WordConstants.N2w.OrdinalNumbers.Tens
		digitToConvert = scaleToConvert.Units
	}

	// Добавить текст к общему результату
	res += lastDigitDeclensionsObject[digitToConvert][genders.FEMALE][numberDeclension][numberScaleForm] + " "
	return res
}

func removeEmptyScalesFromBeginning(numberTripletsArray []objects.RuneDigitTriplet) []objects.RuneDigitTriplet {
	for index, triplet := range numberTripletsArray {
		if !triplet.IsZeros() {
			numberTripletsArray = numberTripletsArray[index:]
			break
		}
	}
	return numberTripletsArray
}

func indexOfLastNotNullScalesByEnd(numberTripletsArray []objects.RuneDigitTriplet) (res int) {
	res = -1
	for i := len(numberTripletsArray) - 1; i >= 0; i-- {
		if !numberTripletsArray[i].IsZeros() {
			res = len(numberTripletsArray) - i - 1
			break
		}
	}
	return res
}

func selectDeclensionsParamsByDeclension(decl declension.Declension,
	isPlural bool) (resDecl declension.Declension, scaleForm objects.ScaleForm) {
	resDecl = decl
	scaleForm = 0
	if isPlural {
		scaleForm = 1
	}

	// Если падеж "именительный" или "винительный" и множественное число
	if isPlural && (decl == declension.NOMINATIVE || decl == declension.ACCUSATIVE) {
		resDecl = declension.GENITIVE
		scaleForm = 1
	}
	return
}
