package ru

import (
	core_constants "github.com/ThCompiler/go_game_constractor/pkg/converter/core/constants"
	core_objects "github.com/ThCompiler/go_game_constractor/pkg/converter/core/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/words"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/objects/currency"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/objects/declension"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/objects/genders"
	"github.com/ThCompiler/ts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConvertNotLowestScaleToWordsSuite struct {
	ts.TestCasesSuite
	rs      Russian
	ActFunc func(words.CurrencyName, core_constants.NumberPart,
		declension.Declension, []core_objects.RuneDigitTriplet) string
}

func (s *ConvertNotLowestScaleToWordsSuite) SetupTest() {
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

	s.ActFunc = func(currencyName words.CurrencyName, numberPart core_constants.NumberPart,
		decl declension.Declension, runeTriplets []core_objects.RuneDigitTriplet) string {
		s.rs.currencyName = currencyName
		s.rs.declension = decl
		return s.rs.GetCurrencyAsWord(numberPart, runeTriplets)
	}
}

func (s *ConvertNotLowestScaleToWordsSuite) TestGetCurrencyForFractionalNumber() {
	s.rs.currencyName = testCurrency
	res := s.rs.GetCurrencyForFractionalNumber()
	assert.Equal(s.T(), s.rs.words.CurrencyStrings.Currencies[testCurrency].
		DecimalCurrencyNameDeclensions[declension.GENITIVE][constants.SINGULAR_WORD], res)
}

func (s *ConvertNotLowestScaleToWordsSuite) TestZeroNumber() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name: "Zero decimal number with currency and declension is NOMINATIVE",
			Args: ts.TTA(
				testCurrency,                // currencyName
				core_constants.INTEGER_PART, // numberType
				declension.NOMINATIVE,       // declension
				zeroTriplet,                 // number triplet
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Zero decimal number without currency and declension is NOMINATIVE",
			Args: ts.TTA(
				words.NUMBER,                // currencyName
				core_constants.INTEGER_PART, // numberType
				declension.NOMINATIVE,       // declension
				zeroTriplet,                 // number triplet
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Zero fractional number with currency and declension is NOMINATIVE",
			Args: ts.TTA(
				testCurrency,                   // currencyName
				core_constants.FRACTIONAL_PART, // numberType
				declension.NOMINATIVE,          // declension
				zeroTriplet,                    // number triplet
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Zero fractional number without currency and declension is NOMINATIVE",
			Args: ts.TTA(
				words.NUMBER,                   // currencyName
				core_constants.FRACTIONAL_PART, // numberType
				declension.NOMINATIVE,          // declension
				zeroTriplet,                    // number triplet
			),
			Expected: ts.TTVE("копеек"),
		},
	)
}

func TestRussianFractionalFunctions(t *testing.T) {
	suite.Run(t, new(ConvertNotLowestScaleToWordsSuite))
}
