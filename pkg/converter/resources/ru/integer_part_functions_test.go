package ru

import (
	core_constants "github.com/ThCompiler/go_game_constractor/pkg/converter/core/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/convert"
	core_objects "github.com/ThCompiler/go_game_constractor/pkg/converter/core/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/words"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/objects/currency"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/objects/declension"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/objects/genders"
	"github.com/ThCompiler/ts"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConvertZeroToWordsForIntegerPartSuite struct {
	ts.TestCasesSuite
	rs *Russian
}

func (s *ConvertZeroToWordsForIntegerPartSuite) SetupTest() {
	s.rs = LoadRussian()
}

func (s *ConvertZeroToWordsForIntegerPartSuite) TestGetZero() {
	require.Equal(s.T(), "ноль", s.rs.GetZeroAsWordsForIntegerPart())
}

var TestCurrency = words.CurrencyName(uuid.New().String())

type ConvertTripletToWordsSuite struct {
	ts.TestCasesSuite
	rs *Russian
}

func (s *ConvertTripletToWordsSuite) SetupTest() {
	s.rs = LoadRussian(
		AddCurrency(TestCurrency, currency.Info{
			CurrencyNounGender: currency.NounGender{
				IntegerPart:    genders.MALE,
				FractionalPart: genders.MALE,
			},
		}),
		WithCurrency(TestCurrency),
		WithDeclension(declension.GENITIVE),
	)
}

func (s *ConvertTripletToWordsSuite) TestDecimalNumber() {
	s.RunTest(
		s.rs.ConvertTripletToWords,
		ts.TestCase{
			Name: "scale is units and use dozens in digits",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				core_objects.NumericDigitTriplet{
					Units:    1,
					Dozens:   2,
					Hundreds: 1,
				},
				convert.UnitsScale,
			),
			Expected: ts.TTVE(core_objects.StringDigitTriplet{
				Units:    "одного",
				Dozens:   "двадцати",
				Hundreds: "ста",
			}),
		},
		ts.TestCase{
			Name: "scale is units and use dozens in tens",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				core_objects.NumericDigitTriplet{
					Units:    1,
					Dozens:   1,
					Hundreds: 1,
				},
				convert.UnitsScale,
			),
			Expected: ts.TTVE(core_objects.StringDigitTriplet{
				Units:    "",
				Dozens:   "одиннадцати",
				Hundreds: "ста",
			}),
		},
		ts.TestCase{
			Name: "scale is thousands and use dozens in digits",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				core_objects.NumericDigitTriplet{
					Units:    1,
					Dozens:   2,
					Hundreds: 1,
				},
				convert.ThousandScale,
			),
			Expected: ts.TTVE(core_objects.StringDigitTriplet{
				Units:    "одной",
				Dozens:   "двадцати",
				Hundreds: "ста",
			}),
		},
		ts.TestCase{
			Name: "scale is thousands and use dozens in tens",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				core_objects.NumericDigitTriplet{
					Units:    1,
					Dozens:   1,
					Hundreds: 1,
				},
				convert.ThousandScale,
			),
			Expected: ts.TTVE(core_objects.StringDigitTriplet{
				Units:    "",
				Dozens:   "одиннадцати",
				Hundreds: "ста",
			}),
		},
		ts.TestCase{
			Name: "scale is more than thousands and use dozens in digits",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				core_objects.NumericDigitTriplet{
					Units:    1,
					Dozens:   2,
					Hundreds: 1,
				},
				20,
			),
			Expected: ts.TTVE(core_objects.StringDigitTriplet{
				Units:    "одного",
				Dozens:   "двадцати",
				Hundreds: "ста",
			}),
		},
		ts.TestCase{
			Name: "scale is more than thousands and use dozens in tens",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				core_objects.NumericDigitTriplet{
					Units:    1,
					Dozens:   1,
					Hundreds: 1,
				},
				20,
			),
			Expected: ts.TTVE(core_objects.StringDigitTriplet{
				Units:    "",
				Dozens:   "одиннадцати",
				Hundreds: "ста",
			}),
		},
	)
}

func (s *ConvertTripletToWordsSuite) TestFractionalNumber() {
	s.RunTest(
		s.rs.ConvertTripletToWords,
		ts.TestCase{
			Name: "scale is units and use dozens in digits",
			Args: ts.TTA(
				core_constants.FRACTIONAL_NUMBER,
				core_objects.NumericDigitTriplet{
					Units:    1,
					Dozens:   2,
					Hundreds: 1,
				},
				convert.UnitsScale,
			),
			Expected: ts.TTVE(core_objects.StringDigitTriplet{
				Units:    "одна",
				Dozens:   "двадцать",
				Hundreds: "сто",
			}),
		},
		ts.TestCase{
			Name: "scale is units and use dozens in tens",
			Args: ts.TTA(
				core_constants.FRACTIONAL_NUMBER,
				core_objects.NumericDigitTriplet{
					Units:    1,
					Dozens:   1,
					Hundreds: 1,
				},
				convert.UnitsScale,
			),
			Expected: ts.TTVE(core_objects.StringDigitTriplet{
				Units:    "",
				Dozens:   "одиннадцать",
				Hundreds: "сто",
			}),
		},
		ts.TestCase{
			Name: "scale is thousands and use dozens in digits",
			Args: ts.TTA(
				core_constants.FRACTIONAL_NUMBER,
				core_objects.NumericDigitTriplet{
					Units:    1,
					Dozens:   2,
					Hundreds: 1,
				},
				convert.ThousandScale,
			),
			Expected: ts.TTVE(core_objects.StringDigitTriplet{
				Units:    "одна",
				Dozens:   "двадцать",
				Hundreds: "сто",
			}),
		},
		ts.TestCase{
			Name: "scale is thousands and use dozens in tens",
			Args: ts.TTA(
				core_constants.FRACTIONAL_NUMBER,
				core_objects.NumericDigitTriplet{
					Units:    1,
					Dozens:   1,
					Hundreds: 1,
				},
				convert.ThousandScale,
			),
			Expected: ts.TTVE(core_objects.StringDigitTriplet{
				Units:    "",
				Dozens:   "одиннадцать",
				Hundreds: "сто",
			}),
		},
		ts.TestCase{
			Name: "scale is more than thousands and use dozens in digits",
			Args: ts.TTA(
				core_constants.FRACTIONAL_NUMBER,
				core_objects.NumericDigitTriplet{
					Units:    1,
					Dozens:   2,
					Hundreds: 1,
				},
				20,
			),
			Expected: ts.TTVE(core_objects.StringDigitTriplet{
				Units:    "один",
				Dozens:   "двадцать",
				Hundreds: "сто",
			}),
		},
		ts.TestCase{
			Name: "scale is more than thousands and use dozens in tens",
			Args: ts.TTA(
				core_constants.FRACTIONAL_NUMBER,
				core_objects.NumericDigitTriplet{
					Units:    1,
					Dozens:   1,
					Hundreds: 1,
				},
				20,
			),
			Expected: ts.TTVE(core_objects.StringDigitTriplet{
				Units:    "",
				Dozens:   "одиннадцать",
				Hundreds: "сто",
			}),
		},
	)
}

var firstFormDigits = core_objects.NumericDigitTriplet{
	Units:    1,
	Dozens:   0,
	Hundreds: 0,
}

var secondFormDigits = core_objects.NumericDigitTriplet{
	Units:    2,
	Dozens:   0,
	Hundreds: 0,
}

var thirdFormDigits = core_objects.NumericDigitTriplet{
	Units:    8,
	Dozens:   0,
	Hundreds: 0,
}

type GetWordScaleNameSuite struct {
	ts.TestCasesSuite
	rs *Russian
}

func (s *GetWordScaleNameSuite) SetupTest() {
	s.rs = LoadRussian()
}

func (s *GetWordScaleNameSuite) TestDecimalNumberInThousandsScale() {
	s.RunTest(
		func(numberType core_constants.NumberType,
			scale int, digits core_objects.NumericDigitTriplet, dec declension.Declension) string {
			s.rs.declension = dec
			return s.rs.GetWordScaleName(numberType, scale, digits)
		},
		ts.TestCase{
			Name: "first form digit and NOMINATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				convert.ThousandScale,
				firstFormDigits,
				declension.NOMINATIVE,
			),
			Expected: ts.TTVE("тысяча"),
		},
		ts.TestCase{
			Name: "second form digit and NOMINATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				convert.ThousandScale,
				secondFormDigits,
				declension.NOMINATIVE,
			),
			Expected: ts.TTVE("тысячи"),
		},
		ts.TestCase{
			Name: "third form digit and NOMINATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				convert.ThousandScale,
				thirdFormDigits,
				declension.NOMINATIVE,
			),
			Expected: ts.TTVE("тысяч"),
		},
		ts.TestCase{
			Name: "first form digit and ACCUSATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				convert.ThousandScale,
				firstFormDigits,
				declension.ACCUSATIVE,
			),
			Expected: ts.TTVE("тысячу"),
		},
		ts.TestCase{
			Name: "second form digit and ACCUSATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				convert.ThousandScale,
				secondFormDigits,
				declension.ACCUSATIVE,
			),
			Expected: ts.TTVE("тысячи"),
		},
		ts.TestCase{
			Name: "third form digit and ACCUSATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				convert.ThousandScale,
				thirdFormDigits,
				declension.ACCUSATIVE,
			),
			Expected: ts.TTVE("тысяч"),
		},
		ts.TestCase{
			Name: "first form digit and DATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				convert.ThousandScale,
				firstFormDigits,
				declension.DATIVE,
			),
			Expected: ts.TTVE("тысяче"),
		},
		ts.TestCase{
			Name: "second form digit and DATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				convert.ThousandScale,
				secondFormDigits,
				declension.DATIVE,
			),
			Expected: ts.TTVE("тысячам"),
		},
		ts.TestCase{
			Name: "third form digit and DATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				convert.ThousandScale,
				thirdFormDigits,
				declension.DATIVE,
			),
			Expected: ts.TTVE("тысячам"),
		},
	)
}

const MillionScale = 3

func (s *GetWordScaleNameSuite) TestDecimalNumberInMillionScale() {
	s.RunTest(
		func(numberType core_constants.NumberType,
			scale int, digits core_objects.NumericDigitTriplet, dec declension.Declension) string {
			s.rs.declension = dec
			return s.rs.GetWordScaleName(numberType, scale, digits)
		},
		ts.TestCase{
			Name: "first form digit and NOMINATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				MillionScale,
				firstFormDigits,
				declension.NOMINATIVE,
			),
			Expected: ts.TTVE("миллион"),
		},
		ts.TestCase{
			Name: "second form digit and NOMINATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				MillionScale,
				secondFormDigits,
				declension.NOMINATIVE,
			),
			Expected: ts.TTVE("миллиона"),
		},
		ts.TestCase{
			Name: "third form digit and NOMINATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				MillionScale,
				thirdFormDigits,
				declension.NOMINATIVE,
			),
			Expected: ts.TTVE("миллионов"),
		},
		ts.TestCase{
			Name: "first form digit and ACCUSATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				MillionScale,
				firstFormDigits,
				declension.NOMINATIVE,
			),
			Expected: ts.TTVE("миллион"),
		},
		ts.TestCase{
			Name: "second form digit and ACCUSATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				MillionScale,
				secondFormDigits,
				declension.NOMINATIVE,
			),
			Expected: ts.TTVE("миллиона"),
		},
		ts.TestCase{
			Name: "third form digit and ACCUSATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				MillionScale,
				thirdFormDigits,
				declension.NOMINATIVE,
			),
			Expected: ts.TTVE("миллионов"),
		},
		ts.TestCase{
			Name: "first form digit and DATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				MillionScale,
				firstFormDigits,
				declension.DATIVE,
			),
			Expected: ts.TTVE("миллиону"),
		},
		ts.TestCase{
			Name: "second form digit and DATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				MillionScale,
				secondFormDigits,
				declension.DATIVE,
			),
			Expected: ts.TTVE("миллионам"),
		},
		ts.TestCase{
			Name: "third form digit and DATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				MillionScale,
				thirdFormDigits,
				declension.DATIVE,
			),
			Expected: ts.TTVE("миллионам"),
		},
	)
}

func (s *GetWordScaleNameSuite) TestFractionalNumber() {
	s.RunTest(
		func(numberType core_constants.NumberType,
			scale int, digits core_objects.NumericDigitTriplet, dec declension.Declension) string {
			s.rs.declension = dec
			return s.rs.GetWordScaleName(numberType, scale, digits)
		},
		ts.TestCase{
			Name: "scale is million, first form digit and DATIVE",
			Args: ts.TTA(
				core_constants.FRACTIONAL_NUMBER,
				MillionScale,
				firstFormDigits,
				declension.DATIVE,
			),
			Expected: ts.TTVE("миллион"),
		},
		ts.TestCase{
			Name: "scale is million, second form digit and DATIVE",
			Args: ts.TTA(
				core_constants.FRACTIONAL_NUMBER,
				MillionScale,
				secondFormDigits,
				declension.DATIVE,
			),
			Expected: ts.TTVE("миллиона"),
		},
		ts.TestCase{
			Name: "scale is million, third form digit and DATIVE",
			Args: ts.TTA(
				core_constants.FRACTIONAL_NUMBER,
				MillionScale,
				thirdFormDigits,
				declension.DATIVE,
			),
			Expected: ts.TTVE("миллионов"),
		},
	)
}

func (s *GetWordScaleNameSuite) TestUnitsScale() {
	s.RunTest(
		func(numberType core_constants.NumberType,
			scale int, digits core_objects.NumericDigitTriplet, dec declension.Declension) string {
			s.rs.declension = dec
			return s.rs.GetWordScaleName(numberType, scale, digits)
		},
		ts.TestCase{
			Name: "first form digit, decimal number and DATIVE",
			Args: ts.TTA(
				core_constants.FRACTIONAL_NUMBER,
				convert.UnitsScale,
				firstFormDigits,
				declension.DATIVE,
			),
			Expected: ts.TTVE(""),
		},
		ts.TestCase{
			Name: "second form digit, fractional number and DATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				convert.UnitsScale,
				secondFormDigits,
				declension.DATIVE,
			),
			Expected: ts.TTVE(""),
		},
		ts.TestCase{
			Name: "third form digit, decimal number and ACCUSATIVE",
			Args: ts.TTA(
				core_constants.DECIMAL_NUMBER,
				convert.UnitsScale,
				thirdFormDigits,
				declension.ACCUSATIVE,
			),
			Expected: ts.TTVE(""),
		},
	)
}

func TestRussianIntegerPartFunctionsSuite(t *testing.T) {
	suite.Run(t, new(ConvertZeroToWordsForIntegerPartSuite))
	suite.Run(t, new(ConvertTripletToWordsSuite))
	suite.Run(t, new(GetWordScaleNameSuite))
}
