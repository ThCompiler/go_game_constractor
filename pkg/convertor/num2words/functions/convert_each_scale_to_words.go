package functions

import (
	"strings"

	"github.com/ThCompiler/go_game_constractor/pkg/convertor/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/num2words/objects"
	words2 "github.com/ThCompiler/go_game_constractor/pkg/convertor/words"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words/declension"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words/genders"
)

const (
	ThirdWordForm  = 2
	SecondWordForm = 1
	FirstWordForm  = 0
)

const defaultForm = ThirdWordForm

func ConvertEachScaleToWords(
	numberScalesArray []objects.RuneDigitTriplet,
	currencyNounGender genders.Gender,
	declension declension.Declension,
) objects.ConvertedScalesToWords {
	numberScalesArrayLen := len(numberScalesArray)

	convertedResult := ""
	// Форма падежа для названия класса единиц или валюты после (0 | 1 | 2).
	scaleNameForm := objects.ScaleForm(defaultForm)
	scaleIsZero := false // Равняется ли целый класс "000".

	// Для каждого класса числа
	for arrIndex, numberScale := range numberScalesArray {
		scaleNameForm = defaultForm // Падеж названия единиц измерения по умолчанию ("рублей")
		scaleIsZero = false
		// Определить порядковый номер текущего класса числа
		currentNumberScale := numberScalesArrayLen - arrIndex

		digits := numberScale.ToNumeric()
		stringDigits := objects.StringDigitTriplet{Units: "", Dozens: "", Hundreds: ""}

		// Если класс числа пустой (000)
		if digits.IsZeros() {
			scaleIsZero = true
			// Если нет классов выше
			if numberScalesArrayLen == 1 {
				convertedResult = convertDigitToWord(
					digits.Units,
					words2.WordConstants.N2w.DigitWords.Units,
					declension,
					genders.MALE,
				)
				scaleNameForm = defaultForm
				// Выйти из цикла
				break
			}
			// Пропустить этот пустой классы (000)
			continue
		}

		/* Определить род числа
		   если класс тысяч - то женский
		   если класс единиц - берем из валюты
		   иначе - мужской */
		gender := genders.MALE
		if currentNumberScale == constants.ThousandScale {
			// Если текущий класс - тысячи
			gender = genders.FEMALE
		} else if currentNumberScale == constants.UnitsScale {
			// Если текущий класс - единицы
			gender = currencyNounGender
		}

		// Определить сотни
		stringDigits.Hundreds = convertDigitToWord(
			digits.Hundreds,
			words2.WordConstants.N2w.DigitWords.Hundreds,
			declension,
			gender,
		)
		// Определить десятки и единицы
		// Если в разряде десятков стоит "1"
		if digits.Dozens == 1 {
			stringDigits.Dozens = convertDigitToWord(
				digits.Units,
				words2.WordConstants.N2w.DigitWords.Tens,
				declension,
				gender,
			)
		} else { // Если в разряде десятков стоит не "1"
			stringDigits.Dozens = convertDigitToWord(
				digits.Dozens,
				words2.WordConstants.N2w.DigitWords.Dozens,
				declension,
				gender,
			)

			stringDigits.Units = convertDigitToWord(
				digits.Units,
				words2.WordConstants.N2w.DigitWords.Units,
				declension,
				gender,
			)

			scaleNameForm = getDigitForm(digits.Units)
		}

		scaleName := getNumberFormScaleName(
			currentNumberScale-1,
			scaleNameForm,
			declension,
		)

		// Убрать ненужный "ноль"
		if digits.Units == 0 && (digits.Hundreds > 0 || digits.Dozens > 0) {
			stringDigits.Units = ""
		}

		// Соединить значения в одну строку
		scaleResult := strings.TrimSpace(stringDigits.Hundreds) + " " +
			strings.TrimSpace(stringDigits.Dozens) + " " +
			strings.TrimSpace(stringDigits.Units) + " " +
			strings.TrimSpace(scaleName)

		// Добавить текущий разобранный класс к общему результату
		convertedResult += " " + scaleResult
	}

	// Вернуть полученный результат и форму падежа для валюты
	return objects.ConvertedScalesToWords{
		Result:          strings.TrimSpace(convertedResult),
		UnitNameForm:    scaleNameForm,
		LastScaleIsZero: scaleIsZero,
	}
}

func convertDigitToWord(digit objects.Digit, digitWords words2.DeclensionNumbers,
	declension declension.Declension, gender genders.Gender,
) string {
	declensionValues := digitWords[declension]
	word := declensionValues[digit]

	if word.WithGender() {
		return word.GetGendersWord()[gender]
	}

	return word.GetWord()
}

// Определить форму названия единиц измерения (рубль/рубля/рублей)
func getDigitForm(digit objects.Digit) objects.ScaleForm {
	// Если цифра в разряде единиц от 1 до 4
	if digit >= 1 && digit <= 4 {
		// Если цифра в разряде единиц "1"
		if digit == 1 {
			// Получиться "рубль"
			return FirstWordForm
		}
		// Получиться "рубля"
		return SecondWordForm
	}

	return ThirdWordForm
}

func getNumberFormScaleName(scale int, scaleNameForm objects.ScaleForm, decl declension.Declension) string {
	if scale == FirstWordForm {
		// Класс единиц
		// Для них название не отображается.
		return ""
	}

	scaleDeclension := decl
	scaleForm := SecondWordForm

	if scaleNameForm == FirstWordForm {
		scaleForm = FirstWordForm
	}

	// Если падеж "именительный" или "винительный" и множественное число
	if (decl == declension.NOMINATIVE || decl == declension.ACCUSATIVE) && scaleNameForm >= SecondWordForm {
		// Для множественного числа именительного падежа используется родительный падеж.
		scaleDeclension = declension.GENITIVE
		scaleForm = SecondWordForm

		if scaleNameForm == SecondWordForm {
			scaleForm = FirstWordForm
		}
	}

	if scale == SecondWordForm {
		// Класс тысяч
		return words2.WordConstants.N2w.UnitScalesNames.Thousands[scaleDeclension][scaleForm]
	}

	// Остальные классы
	ending := words2.WordConstants.N2w.UnitScalesNames.OtherEnding[scaleDeclension][scaleForm]
	base := words2.WordConstants.N2w.UnitScalesNames.OtherBeginning[scale-2]

	return base + ending
}
