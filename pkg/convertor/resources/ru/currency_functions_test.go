package ru

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/core/constants"
	core_objects "github.com/ThCompiler/go_game_constractor/pkg/convertor/core/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/core/words"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/currency"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects/declension"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects/genders"
	ts "github.com/ThCompiler/go_game_constractor/pkg/testing"
	"github.com/stretchr/testify/suite"
	"testing"
)

const (
	testCurrency = words.CurrencyName("TEST")
)

var (
	zeroTriplet = []core_objects.RuneDigitTriplet{
		{
			Units:    '0',
			Dozens:   '0',
			Hundreds: '0',
		},
	}

	firstFormNumber = []core_objects.RuneDigitTriplet{
		{
			Units:    '1',
			Dozens:   '1',
			Hundreds: '1',
		},
	}

	secondFormNumber = []core_objects.RuneDigitTriplet{
		{
			Units:    '3',
			Dozens:   '1',
			Hundreds: '4',
		},
	}
	thirdFormNumber = []core_objects.RuneDigitTriplet{
		{
			Units:    '6',
			Dozens:   '1',
			Hundreds: '1',
		},
	}
)

type GetCurrencyWordSuite struct {
	ts.TestCasesSuite
	rs      Russian
	ActFunc func(args ...interface{}) []interface{}
}

func (s *GetCurrencyWordSuite) SetupTest() {
	s.rs = Russian{
		words: &wordsConstantsForNumbers{
			CurrencyStrings: objects.CurrencyWords{
				Currencies: map[words.CurrencyName]currency.Info{
					testCurrency: {
						DecimalCurrencyNameDeclensions: currency.Declensions{
							declension.NOMINATIVE:    [2]string{"рубль", ""},
							declension.GENITIVE:      [2]string{"рубля", "рублей"},
							declension.DATIVE:        [2]string{"рублю", "рублям"},
							declension.ACCUSATIVE:    [2]string{"рубль", ""},
							declension.INSTRUMENTAL:  [2]string{"рублём", "рублями"},
							declension.PREPOSITIONAL: [2]string{"рубле", "рублях"},
						},
						FractionalPartNameDeclensions: currency.Declensions{
							declension.NOMINATIVE:    [2]string{"копейка", ""},
							declension.GENITIVE:      [2]string{"копейки", "копеек"},
							declension.DATIVE:        [2]string{"копейке", "копейкам"},
							declension.ACCUSATIVE:    [2]string{"копейка", ""},
							declension.INSTRUMENTAL:  [2]string{"копейкой", "копейками"},
							declension.PREPOSITIONAL: [2]string{"копейке", "копейках"},
						},
						CurrencyNounGender: currency.NounGender{
							IntegerPart:    genders.MALE,
							FractionalPart: genders.FEMALE,
						},
					},
					words.NUMBER: {
						DecimalCurrencyNameDeclensions: currency.Declensions{
							declension.NOMINATIVE:    [2]string{"рубль", ""},
							declension.GENITIVE:      [2]string{"рубля", "рублей"},
							declension.DATIVE:        [2]string{"рублю", "рублям"},
							declension.ACCUSATIVE:    [2]string{"рубль", ""},
							declension.INSTRUMENTAL:  [2]string{"рублём", "рублями"},
							declension.PREPOSITIONAL: [2]string{"рубле", "рублях"},
						},
						FractionalPartNameDeclensions: currency.Declensions{
							declension.NOMINATIVE:    [2]string{"копейка", ""},
							declension.GENITIVE:      [2]string{"копейки", "копеек"},
							declension.DATIVE:        [2]string{"копейке", "копейкам"},
							declension.ACCUSATIVE:    [2]string{"копейка", ""},
							declension.INSTRUMENTAL:  [2]string{"копейкой", "копейками"},
							declension.PREPOSITIONAL: [2]string{"копейке", "копейках"},
						},
						CurrencyNounGender: currency.NounGender{
							IntegerPart:    genders.MALE,
							FractionalPart: genders.FEMALE,
						},
					},
				},
			},
		},
	}

	s.ActFunc = func(args ...interface{}) []interface{} {
		numberInfo := words.NumberInfo{
			NumberType:   args[1].(constants.NumberType),
			Declension:   args[2].(words.Declension),
			CurrencyName: args[0].(words.CurrencyName),
		}
		res := s.rs.GetCurrencyAsWord(numberInfo, args[3].([]core_objects.RuneDigitTriplet))
		return []interface{}{res}
	}
}

func (s *GetCurrencyWordSuite) TestZeroNumber() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name: "Zero decimal number with currency and declension is NOMINATIVE",
			Args: ts.TTA(
				testCurrency,             // currencyName
				constants.DECIMAL_NUMBER, // numberType
				declension.NOMINATIVE,    // declension
				zeroTriplet,              // number triplet
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Zero decimal number without currency and declension is NOMINATIVE",
			Args: ts.TTA(
				words.NUMBER,             // currencyName
				constants.DECIMAL_NUMBER, // numberType
				declension.NOMINATIVE,    // declension
				zeroTriplet,              // number triplet
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Zero fractional number with currency and declension is NOMINATIVE",
			Args: ts.TTA(
				testCurrency,                // currencyName
				constants.FRACTIONAL_NUMBER, // numberType
				declension.NOMINATIVE,       // declension
				zeroTriplet,                 // number triplet
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Zero fractional number without currency and declension is NOMINATIVE",
			Args: ts.TTA(
				words.NUMBER,                // currencyName
				constants.FRACTIONAL_NUMBER, // numberType
				declension.NOMINATIVE,       // declension
				zeroTriplet,                 // number triplet
			),
			Expected: ts.TTVE("копеек"),
		},
	)
}

//----------------------------------------------------------------------------------------------------------------------
// Decimal Number
//----------------------------------------------------------------------------------------------------------------------

func (s *GetCurrencyWordSuite) TestDecimalNumberWithFirstForm() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name: "First form number with currency and declension is NOMINATIVE",
			Args: ts.TTA(
				testCurrency,             // currencyName
				constants.DECIMAL_NUMBER, // numberType
				declension.NOMINATIVE,    // declension
				firstFormNumber,          // number triplet
			),
			Expected: ts.TTVE("рубль"),
		},
		ts.TestCase{
			Name: "First form number with currency and declension is ACCUSATIVE",
			Args: ts.TTA(
				testCurrency,             // currencyName
				constants.DECIMAL_NUMBER, // numberType
				declension.ACCUSATIVE,    // declension
				firstFormNumber,          // number triplet
			),
			Expected: ts.TTVE("рубль"),
		},
		ts.TestCase{
			Name: "First form number with currency and declension is DATIVE",
			Args: ts.TTA(
				testCurrency,             // currencyName
				constants.DECIMAL_NUMBER, // numberType
				declension.DATIVE,        // declension
				firstFormNumber,          // number triplet
			),
			Expected: ts.TTVE("рублю"),
		},
		ts.TestCase{
			Name: "First form number without currency and declension is NOMINATIVE",
			Args: ts.TTA(
				words.NUMBER,             // currencyName
				constants.DECIMAL_NUMBER, // numberType
				declension.NOMINATIVE,    // declension
				firstFormNumber,          // number triplet
			),
			Expected: ts.TTVE("рубль"),
		},
	)
}

func (s *GetCurrencyWordSuite) TestDecimalNumberWithSecondForm() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name: "Second form number with currency and declension is NOMINATIVE",
			Args: ts.TTA(
				testCurrency,             // currencyName
				constants.DECIMAL_NUMBER, // numberType
				declension.NOMINATIVE,    // declension
				secondFormNumber,         // number triplet
			),
			Expected: ts.TTVE("рубля"),
		},
		ts.TestCase{
			Name: "Second form number with currency and declension is ACCUSATIVE",
			Args: ts.TTA(
				testCurrency,             // currencyName
				constants.DECIMAL_NUMBER, // numberType
				declension.ACCUSATIVE,    // declension
				secondFormNumber,         // number triplet
			),
			Expected: ts.TTVE("рубля"),
		},
		ts.TestCase{
			Name: "Second form number with currency and declension is DATIVE",
			Args: ts.TTA(
				testCurrency,             // currencyName
				constants.DECIMAL_NUMBER, // numberType
				declension.DATIVE,        // declension
				secondFormNumber,         // number triplet
			),
			Expected: ts.TTVE("рублям"),
		},
		ts.TestCase{
			Name: "Second form number without currency and declension is NOMINATIVE",
			Args: ts.TTA(
				words.NUMBER,             // currencyName
				constants.DECIMAL_NUMBER, // numberType
				declension.NOMINATIVE,    // declension
				secondFormNumber,         // number triplet
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Second form number without currency and declension is ACCUSATIVE",
			Args: ts.TTA(
				words.NUMBER,             // currencyName
				constants.DECIMAL_NUMBER, // numberType
				declension.ACCUSATIVE,    // declension
				secondFormNumber,         // number triplet
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Second form number without currency and declension is DATIVE",
			Args: ts.TTA(
				words.NUMBER,             // currencyName
				constants.DECIMAL_NUMBER, // numberType
				declension.DATIVE,        // declension
				secondFormNumber,         // number triplet
			),
			Expected: ts.TTVE("рублям"),
		},
	)
}

func (s *GetCurrencyWordSuite) TestDecimalNumberWithThirdForm() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name: "Third form number with currency and declension is NOMINATIVE",
			Args: ts.TTA(
				testCurrency,             // currencyName
				constants.DECIMAL_NUMBER, // numberType
				declension.NOMINATIVE,    // declension
				thirdFormNumber,          // number triplet
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Third form number with currency and declension is ACCUSATIVE",
			Args: ts.TTA(
				testCurrency,             // currencyName
				constants.DECIMAL_NUMBER, // numberType
				declension.ACCUSATIVE,    // declension
				thirdFormNumber,          // number triplet
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Third form number with currency and declension is DATIVE",
			Args: ts.TTA(
				testCurrency,             // currencyName
				constants.DECIMAL_NUMBER, // numberType
				declension.DATIVE,        // declension
				thirdFormNumber,          // number triplet
			),
			Expected: ts.TTVE("рублям"),
		},
		ts.TestCase{
			Name: "Third form number without currency and declension is NOMINATIVE",
			Args: ts.TTA(
				words.NUMBER,             // currencyName
				constants.DECIMAL_NUMBER, // numberType
				declension.NOMINATIVE,    // declension
				thirdFormNumber,          // number triplet
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Third form number without currency and declension is ACCUSATIVE",
			Args: ts.TTA(
				words.NUMBER,             // currencyName
				constants.DECIMAL_NUMBER, // numberType
				declension.ACCUSATIVE,    // declension
				thirdFormNumber,          // number triplet
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Third form number without currency and declension is DATIVE",
			Args: ts.TTA(
				words.NUMBER,             // currencyName
				constants.DECIMAL_NUMBER, // numberType
				declension.DATIVE,        // declension
				thirdFormNumber,          // number triplet
			),
			Expected: ts.TTVE("рублям"),
		},
	)
}

//----------------------------------------------------------------------------------------------------------------------
// Fractional Number
//----------------------------------------------------------------------------------------------------------------------

func (s *GetCurrencyWordSuite) TestFractionalNumberWithFirstForm() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name: "First form number with currency and declension is NOMINATIVE",
			Args: ts.TTA(
				testCurrency,                // currencyName
				constants.FRACTIONAL_NUMBER, // numberType
				declension.NOMINATIVE,       // declension
				firstFormNumber,             // number triplet
			),
			Expected: ts.TTVE("копейка"),
		},
		ts.TestCase{
			Name: "First form number with currency and declension is ACCUSATIVE",
			Args: ts.TTA(
				testCurrency,                // currencyName
				constants.FRACTIONAL_NUMBER, // numberType
				declension.ACCUSATIVE,       // declension
				firstFormNumber,             // number triplet
			),
			Expected: ts.TTVE("копейка"),
		},
		ts.TestCase{
			Name: "First form number with currency and declension is DATIVE",
			Args: ts.TTA(
				testCurrency,                // currencyName
				constants.FRACTIONAL_NUMBER, // numberType
				declension.DATIVE,           // declension
				firstFormNumber,             // number triplet
			),
			Expected: ts.TTVE("копейке"),
		},
		ts.TestCase{
			Name: "First form number without currency and declension is NOMINATIVE",
			Args: ts.TTA(
				words.NUMBER,                // currencyName
				constants.FRACTIONAL_NUMBER, // numberType
				declension.NOMINATIVE,       // declension
				firstFormNumber,             // number triplet
			),
			Expected: ts.TTVE("копейка"),
		},
	)
}

func (s *GetCurrencyWordSuite) TestFractionalNumberWithSecondForm() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name: "Second form number with currency and declension is NOMINATIVE",
			Args: ts.TTA(
				testCurrency,                // currencyName
				constants.FRACTIONAL_NUMBER, // numberType
				declension.NOMINATIVE,       // declension
				secondFormNumber,            // number triplet
			),
			Expected: ts.TTVE("копейки"),
		},
		ts.TestCase{
			Name: "Second form number with currency and declension is ACCUSATIVE",
			Args: ts.TTA(
				testCurrency,                // currencyName
				constants.FRACTIONAL_NUMBER, // numberType
				declension.ACCUSATIVE,       // declension
				secondFormNumber,            // number triplet
			),
			Expected: ts.TTVE("копейки"),
		},
		ts.TestCase{
			Name: "Second form number with currency and declension is DATIVE",
			Args: ts.TTA(
				testCurrency,                // currencyName
				constants.FRACTIONAL_NUMBER, // numberType
				declension.DATIVE,           // declension
				secondFormNumber,            // number triplet
			),
			Expected: ts.TTVE("копейкам"),
		},
		ts.TestCase{
			Name: "Second form number without currency and declension is NOMINATIVE",
			Args: ts.TTA(
				words.NUMBER,                // currencyName
				constants.FRACTIONAL_NUMBER, // numberType
				declension.NOMINATIVE,       // declension
				secondFormNumber,            // number triplet
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Second form number without currency and declension is ACCUSATIVE",
			Args: ts.TTA(
				words.NUMBER,                // currencyName
				constants.FRACTIONAL_NUMBER, // numberType
				declension.ACCUSATIVE,       // declension
				secondFormNumber,            // number triplet
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Second form number without currency and declension is DATIVE",
			Args: ts.TTA(
				words.NUMBER,                // currencyName
				constants.FRACTIONAL_NUMBER, // numberType
				declension.DATIVE,           // declension
				secondFormNumber,            // number triplet
			),
			Expected: ts.TTVE("копейкам"),
		},
	)
}

func (s *GetCurrencyWordSuite) TestFractionalNumberWithThirdForm() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name: "Third form number with currency and declension is NOMINATIVE",
			Args: ts.TTA(
				testCurrency,                // currencyName
				constants.FRACTIONAL_NUMBER, // numberType
				declension.NOMINATIVE,       // declension
				thirdFormNumber,             // number triplet
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Third form number with currency and declension is ACCUSATIVE",
			Args: ts.TTA(
				testCurrency,                // currencyName
				constants.FRACTIONAL_NUMBER, // numberType
				declension.ACCUSATIVE,       // declension
				thirdFormNumber,             // number triplet
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Third form number with currency and declension is DATIVE",
			Args: ts.TTA(
				testCurrency,                // currencyName
				constants.FRACTIONAL_NUMBER, // numberType
				declension.DATIVE,           // declension
				thirdFormNumber,             // number triplet
			),
			Expected: ts.TTVE("копейкам"),
		},
		ts.TestCase{
			Name: "Third form number without currency and declension is NOMINATIVE",
			Args: ts.TTA(
				words.NUMBER,                // currencyName
				constants.FRACTIONAL_NUMBER, // numberType
				declension.NOMINATIVE,       // declension
				thirdFormNumber,             // number triplet
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Third form number without currency and declension is ACCUSATIVE",
			Args: ts.TTA(
				words.NUMBER,                // currencyName
				constants.FRACTIONAL_NUMBER, // numberType
				declension.ACCUSATIVE,       // declension
				thirdFormNumber,             // number triplet
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Third form number without currency and declension is DATIVE",
			Args: ts.TTA(
				words.NUMBER,                // currencyName
				constants.FRACTIONAL_NUMBER, // numberType
				declension.DATIVE,           // declension
				thirdFormNumber,             // number triplet
			),
			Expected: ts.TTVE("копейкам"),
		},
	)
}

func TestGetCurrencyWordSuite(t *testing.T) {
	suite.Run(t, new(GetCurrencyWordSuite))
}
