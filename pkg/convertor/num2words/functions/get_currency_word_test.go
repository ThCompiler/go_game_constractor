package functions

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/currency"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/num2words/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words/declension"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words/genders"
	ts "github.com/ThCompiler/go_game_constractor/pkg/testing"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GetCurrencyWordSuite struct {
	ts.TestCasesSuite
	currency currency.CustomCurrency
	ActFunc  func(args ...interface{}) []interface{}
}

func (s *GetCurrencyWordSuite) SetupTest() {
	s.currency = currency.CustomCurrency{
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
	}
	s.ActFunc = func(args ...interface{}) []interface{} {
		res := GetCurrencyWord(args[0].(currency.CustomCurrency), args[1].(constants.NumberType),
			args[2].(objects.NumberForm), args[3].(bool), args[4].(bool), args[5].(declension.Declension))
		return []interface{}{res}
	}
}

//----------------------------------------------------------------------------------------------------------------------
// Decimal Number
//----------------------------------------------------------------------------------------------------------------------

func (s *GetCurrencyWordSuite) TestDecimalNumberWithFirstForm() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name: "Zeros triplet is false and declension is NOMINATIVE",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.FIRST_FORM,       // numberForm
				false,                    // lowestTripletIsZero
				false,                    // isNumber
				declension.NOMINATIVE,    // declension
			),
			Expected: ts.TTVE("рубль"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and declension is ACCUSATIVE",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.FIRST_FORM,       // numberForm
				false,                    // lowestTripletIsZero
				false,                    // isNumber
				declension.ACCUSATIVE,    // declensions
			),
			Expected: ts.TTVE("рубль"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and declension is DATIVE",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.FIRST_FORM,       // numberForm
				false,                    // lowestTripletIsZero
				false,                    // isNumber
				declension.DATIVE,        // declensions
			),
			Expected: ts.TTVE("рублю"),
		},
		ts.TestCase{
			Name: "Zeros triplet is true and declension is NOMINATIVE",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.FIRST_FORM,       // numberForm
				true,                     // lowestTripletIsZero
				false,                    // isNumber
				declension.NOMINATIVE,    // declensions
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Zeros triplet is true and it is number",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.FIRST_FORM,       // numberForm
				true,                     // lowestTripletIsZero
				true,                     // isNumber
				declension.NOMINATIVE,    // declensions
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and it is number",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.FIRST_FORM,       // numberForm
				false,                    // lowestTripletIsZero
				true,                     // isNumber
				declension.NOMINATIVE,    // declensions
			),
			Expected: ts.TTVE("рубль"),
		},
	)
}

func (s *GetCurrencyWordSuite) TestDecimalNumberWithSecondForm() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name: "Zeros triplet is false and declension is NOMINATIVE",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.SECOND_FORM,      // numberForm
				false,                    // lowestTripletIsZero
				false,                    // isNumber
				declension.NOMINATIVE,    // declension
			),
			Expected: ts.TTVE("рубля"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and declension is ACCUSATIVE",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.SECOND_FORM,      // numberForm
				false,                    // lowestTripletIsZero
				false,                    // isNumber
				declension.ACCUSATIVE,    // declensions
			),
			Expected: ts.TTVE("рубля"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and declension is DATIVE",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.SECOND_FORM,      // numberForm
				false,                    // lowestTripletIsZero
				false,                    // isNumber
				declension.DATIVE,        // declensions
			),
			Expected: ts.TTVE("рублям"),
		},
		ts.TestCase{
			Name: "Zeros triplet is true and declension is NOMINATIVE",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.SECOND_FORM,      // numberForm
				true,                     // lowestTripletIsZero
				false,                    // isNumber
				declension.NOMINATIVE,    // declensions
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Zeros triplet is true and it is number",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.SECOND_FORM,      // numberForm
				true,                     // lowestTripletIsZero
				true,                     // isNumber
				declension.NOMINATIVE,    // declensions
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and it is number and is NOMINATIVE",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.SECOND_FORM,      // numberForm
				false,                    // lowestTripletIsZero
				true,                     // isNumber
				declension.NOMINATIVE,    // declensions
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and it is number and is ACCUSATIVE",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.SECOND_FORM,      // numberForm
				false,                    // lowestTripletIsZero
				true,                     // isNumber
				declension.ACCUSATIVE,    // declensions
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and it is number and is DATIVE",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.SECOND_FORM,      // numberForm
				false,                    // lowestTripletIsZero
				true,                     // isNumber
				declension.DATIVE,        // declensions
			),
			Expected: ts.TTVE("рублям"),
		},
	)
}

func (s *GetCurrencyWordSuite) TestDecimalNumberWithThirdForm() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name: "Zeros triplet is false and declension is NOMINATIVE",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.THIRD_FORM,       // numberForm
				false,                    // lowestTripletIsZero
				false,                    // isNumber
				declension.NOMINATIVE,    // declension
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and declension is ACCUSATIVE",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.THIRD_FORM,       // numberForm
				false,                    // lowestTripletIsZero
				false,                    // isNumber
				declension.ACCUSATIVE,    // declensions
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and declension is DATIVE",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.THIRD_FORM,       // numberForm
				false,                    // lowestTripletIsZero
				false,                    // isNumber
				declension.DATIVE,        // declensions
			),
			Expected: ts.TTVE("рублям"),
		},
		ts.TestCase{
			Name: "Zeros triplet is true and declension is NOMINATIVE",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.THIRD_FORM,       // numberForm
				true,                     // lowestTripletIsZero
				false,                    // isNumber
				declension.NOMINATIVE,    // declensions
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Zeros triplet is true and it is number",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.THIRD_FORM,       // numberForm
				true,                     // lowestTripletIsZero
				true,                     // isNumber
				declension.NOMINATIVE,    // declensions
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and it is number and is NOMINATIVE",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.THIRD_FORM,       // numberForm
				false,                    // lowestTripletIsZero
				true,                     // isNumber
				declension.NOMINATIVE,    // declensions
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and it is number and is ACCUSATIVE",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.THIRD_FORM,       // numberForm
				false,                    // lowestTripletIsZero
				true,                     // isNumber
				declension.ACCUSATIVE,    // declensions
			),
			Expected: ts.TTVE("рублей"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and it is number and is DATIVE",
			Args: ts.TTA(
				s.currency,               // currencyObject
				constants.DECIMAL_NUMBER, // numberType
				objects.THIRD_FORM,       // numberForm
				false,                    // lowestTripletIsZero
				true,                     // isNumber
				declension.DATIVE,        // declensions
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
			Name: "Zeros triplet is false and declension is NOMINATIVE",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.FIRST_FORM,          // numberForm
				false,                       // lowestTripletIsZero
				false,                       // isNumber
				declension.NOMINATIVE,       // declension
			),
			Expected: ts.TTVE("копейка"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and declension is ACCUSATIVE",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.FIRST_FORM,          // numberForm
				false,                       // lowestTripletIsZero
				false,                       // isNumber
				declension.ACCUSATIVE,       // declensions
			),
			Expected: ts.TTVE("копейка"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and declension is DATIVE",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.FIRST_FORM,          // numberForm
				false,                       // lowestTripletIsZero
				false,                       // isNumber
				declension.DATIVE,           // declensions
			),
			Expected: ts.TTVE("копейке"),
		},
		ts.TestCase{
			Name: "Zeros triplet is true and declension is NOMINATIVE",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.FIRST_FORM,          // numberForm
				true,                        // lowestTripletIsZero
				false,                       // isNumber
				declension.NOMINATIVE,       // declensions
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Zeros triplet is true and it is number",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.FIRST_FORM,          // numberForm
				true,                        // lowestTripletIsZero
				true,                        // isNumber
				declension.NOMINATIVE,       // declensions
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and it is number",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.FIRST_FORM,          // numberForm
				false,                       // lowestTripletIsZero
				true,                        // isNumber
				declension.NOMINATIVE,       // declensions
			),
			Expected: ts.TTVE("копейка"),
		},
	)
}

func (s *GetCurrencyWordSuite) TestFractionalNumberWithSecondForm() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name: "Zeros triplet is false and declension is NOMINATIVE",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.SECOND_FORM,         // numberForm
				false,                       // lowestTripletIsZero
				false,                       // isNumber
				declension.NOMINATIVE,       // declension
			),
			Expected: ts.TTVE("копейки"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and declension is ACCUSATIVE",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.SECOND_FORM,         // numberForm
				false,                       // lowestTripletIsZero
				false,                       // isNumber
				declension.ACCUSATIVE,       // declensions
			),
			Expected: ts.TTVE("копейки"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and declension is DATIVE",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.SECOND_FORM,         // numberForm
				false,                       // lowestTripletIsZero
				false,                       // isNumber
				declension.DATIVE,           // declensions
			),
			Expected: ts.TTVE("копейкам"),
		},
		ts.TestCase{
			Name: "Zeros triplet is true and declension is NOMINATIVE",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.SECOND_FORM,         // numberForm
				true,                        // lowestTripletIsZero
				false,                       // isNumber
				declension.NOMINATIVE,       // declensions
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Zeros triplet is true and it is number",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.SECOND_FORM,         // numberForm
				true,                        // lowestTripletIsZero
				true,                        // isNumber
				declension.NOMINATIVE,       // declensions
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and it is number and is NOMINATIVE",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.SECOND_FORM,         // numberForm
				false,                       // lowestTripletIsZero
				true,                        // isNumber
				declension.NOMINATIVE,       // declensions
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and it is number and is ACCUSATIVE",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.SECOND_FORM,         // numberForm
				false,                       // lowestTripletIsZero
				true,                        // isNumber
				declension.ACCUSATIVE,       // declensions
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and it is number and is DATIVE",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.SECOND_FORM,         // numberForm
				false,                       // lowestTripletIsZero
				true,                        // isNumber
				declension.DATIVE,           // declensions
			),
			Expected: ts.TTVE("копейкам"),
		},
	)
}

func (s *GetCurrencyWordSuite) TestFractionalNumberWithThirdForm() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name: "Zeros triplet is false and declension is NOMINATIVE",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.THIRD_FORM,          // numberForm
				false,                       // lowestTripletIsZero
				false,                       // isNumber
				declension.NOMINATIVE,       // declension
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and declension is ACCUSATIVE",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.THIRD_FORM,          // numberForm
				false,                       // lowestTripletIsZero
				false,                       // isNumber
				declension.ACCUSATIVE,       // declensions
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and declension is DATIVE",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.THIRD_FORM,          // numberForm
				false,                       // lowestTripletIsZero
				false,                       // isNumber
				declension.DATIVE,           // declensions
			),
			Expected: ts.TTVE("копейкам"),
		},
		ts.TestCase{
			Name: "Zeros triplet is true and declension is NOMINATIVE",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.THIRD_FORM,          // numberForm
				true,                        // lowestTripletIsZero
				false,                       // isNumber
				declension.NOMINATIVE,       // declensions
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Zeros triplet is true and it is number",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.THIRD_FORM,          // numberForm
				true,                        // lowestTripletIsZero
				true,                        // isNumber
				declension.NOMINATIVE,       // declensions
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and it is number and is NOMINATIVE",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.THIRD_FORM,          // numberForm
				false,                       // lowestTripletIsZero
				true,                        // isNumber
				declension.NOMINATIVE,       // declensions
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and it is number and is ACCUSATIVE",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.THIRD_FORM,          // numberForm
				false,                       // lowestTripletIsZero
				true,                        // isNumber
				declension.ACCUSATIVE,       // declensions
			),
			Expected: ts.TTVE("копеек"),
		},
		ts.TestCase{
			Name: "Zeros triplet is false and it is number and is DATIVE",
			Args: ts.TTA(
				s.currency,                  // currencyObject
				constants.FRACTIONAL_NUMBER, // numberType
				objects.THIRD_FORM,          // numberForm
				false,                       // lowestTripletIsZero
				true,                        // isNumber
				declension.DATIVE,           // declensions
			),
			Expected: ts.TTVE("копейкам"),
		},
	)
}

func TestGetCurrencyWordSuite(t *testing.T) {
	suite.Run(t, new(GetCurrencyWordSuite))
}
