package functions

import (
	"math"

	"github.com/ThCompiler/go_game_constractor/pkg/convertor/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/num2words/objects"
	words2 "github.com/ThCompiler/go_game_constractor/pkg/convertor/words"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words/declension"
)

const two = 2.0

func GetFractionalUnitCurrencyNumber(scaleIndex int, digitToConvert objects.Digit,
	decl declension.Declension, unitNameForm objects.ScaleForm,
) string {
	if scaleIndex < 0 {
		scaleIndex = 0
	}

	res := ""
	unitDeclensionsObject := words2.WordConstants.N2w.FractionalUnit.FractionalUnitsDeclensions.Tens

	if scaleIndex == 1 {
		unitDeclensionsObject = words2.WordConstants.N2w.FractionalUnit.FractionalUnitsDeclensions.Hundreds
	} else if scaleIndex > 1 {
		// Определить класс числа
		// (0 - единицы, 1 - тысячи, 2 - миллионы и т.д.).
		numberScale := int(math.Ceil((float64(scaleIndex)+two)/constants.TripletSize) - 1.0)
		if numberScale == 0 {
			return ""
		}

		// Получить разряд цифры в текущем классе
		// (0 - единицы, 1 - десятки, 2 - сотни).
		digitIndexInScale := scaleIndex - numberScale*3 + 1
		// Получить корень названия класса числа
		unitNameBegin := words2.WordConstants.N2w.FractionalUnit.FractionalUnitsBeginning[numberScale-1]

		if numberScale > len(words2.WordConstants.N2w.FractionalUnit.FractionalUnitsBeginning) {
			unitNameBegin = words2.WordConstants.N2w.UnitScalesNames.OtherBeginning[numberScale-2]
		}

		// Получить приставку к числу
		unitNamePrefix := words2.WordConstants.N2w.FractionalUnit.FractionalUnitPrefixes[digitIndexInScale]

		// Составить объект с падежами
		unitDeclensionsObject = words2.DeclensionFractionalUnits{}

		for key, value := range words2.WordConstants.N2w.FractionalUnit.FractionalUnitEndings {
			newWords := [constants.CountScaleNumberNameForms]string{}
			for i, val := range value {
				newWords[i] = unitNamePrefix + unitNameBegin + val
			}
			unitDeclensionsObject[key] = newWords
		}
	}

	// Выбрать правильную форму слова
	numberDecl, numberScaleForm := selectDeclensionsParamsByDeclension(decl, unitNameForm != 0)
	res = unitDeclensionsObject[numberDecl][numberScaleForm]

	// Если цифра для конвертирования === 0
	if digitToConvert == 0 {
		// Использовать родительный падеж.
		res = unitDeclensionsObject[declension.GENITIVE][1]
	}

	return res
}
