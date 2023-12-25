package ru

import (
	core_constants "github.com/ThCompiler/go_game_constractor/pkg/convertor/core/constants"
	core_objects "github.com/ThCompiler/go_game_constractor/pkg/convertor/core/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/core/words"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/currency"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects/declension"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects/genders"
	"github.com/ThCompiler/ts"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RussianUtilityFunctionsSuite struct {
	ts.TestCasesSuite
	rs      Russian
	ActFunc func(args ...interface{}) []interface{}
}

func (s *RussianUtilityFunctionsSuite) SetupTest() {
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
			NumberType:   args[1].(core_constants.NumberType),
			Declension:   args[2].(words.Declension),
			CurrencyName: args[0].(words.CurrencyName),
		}
		res := s.rs.GetCurrencyAsWord(numberInfo, args[3].([]core_objects.RuneDigitTriplet))
		return []interface{}{res}
	}
}

func TestRussianUtilityFunctionsSuite(t *testing.T) {
	suite.Run(t, new(RussianUtilityFunctionsSuite))
}
