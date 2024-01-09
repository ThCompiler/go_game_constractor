package ru

import (
	"encoding/json"
	coreobjects "github.com/ThCompiler/go_game_constractor/pkg/convertor/core/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects/declension"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects/genders"
	"github.com/ThCompiler/ts"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type RussianUtilityFunctionsSuite struct {
	ts.TestCasesSuite
	rs      Russian
	ActFunc func(args ...interface{}) []interface{}
}

func (r *RussianUtilityFunctionsSuite) TestGetNumberFormByDigitFunction() {
	r.RunTest(
		GetNumberFormByDigit,
		ts.TestCase{
			Name:     "zero input number",
			Args:     ts.TTA(0),
			Expected: ts.TTVE(constants.THIRD_FORM),
		},
		ts.TestCase{
			Name:     "first form number (1)",
			Args:     ts.TTA(1),
			Expected: ts.TTVE(constants.FIRST_FORM),
		},
		ts.TestCase{
			Name:     "second form number (2)",
			Args:     ts.TTA(2),
			Expected: ts.TTVE(constants.SECOND_FORM),
		},
		ts.TestCase{
			Name:     "second form number (4)",
			Args:     ts.TTA(4),
			Expected: ts.TTVE(constants.SECOND_FORM),
		},
		ts.TestCase{
			Name:     "third form number (5)",
			Args:     ts.TTA(5),
			Expected: ts.TTVE(constants.THIRD_FORM),
		},
		ts.TestCase{
			Name:     "third form number (9)",
			Args:     ts.TTA(9),
			Expected: ts.TTVE(constants.THIRD_FORM),
		},
		ts.TestCase{
			Name:     "invalid positive number",
			Args:     ts.TTA(10),
			Expected: ts.TTVE(constants.INVALID_FORM),
		},
		ts.TestCase{
			Name:     "invalid negative number",
			Args:     ts.TTA(-1),
			Expected: ts.TTVE(constants.INVALID_FORM),
		},
	)
}

func (r *RussianUtilityFunctionsSuite) TestGetNumberFormByTripletFunction() {
	r.RunTest(
		GetNumberFormByTriplet,
		ts.TestCase{
			Name:     "zero triplet",
			Args:     ts.TTA(coreobjects.NumericDigitTriplet{Hundreds: 2, Dozens: 1, Units: 0}),
			Expected: ts.TTVE(constants.THIRD_FORM),
		},
		ts.TestCase{
			Name:     "first form triplet",
			Args:     ts.TTA(coreobjects.NumericDigitTriplet{Hundreds: 2, Dozens: 1, Units: 1}),
			Expected: ts.TTVE(constants.FIRST_FORM),
		},
		ts.TestCase{
			Name:     "second form triplet",
			Args:     ts.TTA(coreobjects.NumericDigitTriplet{Hundreds: 2, Dozens: 1, Units: 3}),
			Expected: ts.TTVE(constants.SECOND_FORM),
		},
		ts.TestCase{
			Name:     "third form triplet",
			Args:     ts.TTA(coreobjects.NumericDigitTriplet{Hundreds: 2, Dozens: 1, Units: 8}),
			Expected: ts.TTVE(constants.THIRD_FORM),
		},
		ts.TestCase{
			Name:     "invalid triplet",
			Args:     ts.TTA(coreobjects.NumericDigitTriplet{Hundreds: 127, Dozens: 127, Units: 127}),
			Expected: ts.TTVE(constants.INVALID_FORM),
		},
		ts.TestCase{
			Name:     "check not dependent number form from hundreds and dozens (they are incorrect)",
			Args:     ts.TTA(coreobjects.NumericDigitTriplet{Hundreds: 127, Dozens: 127, Units: 1}),
			Expected: ts.TTVE(constants.FIRST_FORM),
		},
	)
}

func (r *RussianUtilityFunctionsSuite) TestGetNumberFormFunction() {
	r.RunTest(
		GetNumberForm,
		ts.TestCase{
			Name:     "zero triplets",
			Args:     ts.TTA([]coreobjects.RuneDigitTriplet{{'2', '1', '0'}}),
			Expected: ts.TTVE(constants.THIRD_FORM),
		},
		ts.TestCase{
			Name:     "first form triplets",
			Args:     ts.TTA([]coreobjects.RuneDigitTriplet{{'2', '1', '1'}}),
			Expected: ts.TTVE(constants.FIRST_FORM),
		},
		ts.TestCase{
			Name:     "second form triplets",
			Args:     ts.TTA([]coreobjects.RuneDigitTriplet{{'2', '1', '3'}}),
			Expected: ts.TTVE(constants.SECOND_FORM),
		},
		ts.TestCase{
			Name:     "third form triplets",
			Args:     ts.TTA([]coreobjects.RuneDigitTriplet{{'2', '1', '8'}}),
			Expected: ts.TTVE(constants.THIRD_FORM),
		},
		ts.TestCase{
			Name:     "invalid triplets",
			Args:     ts.TTA([]coreobjects.RuneDigitTriplet{{'a', 'a', 'a'}}),
			Expected: ts.TTVE(constants.INVALID_FORM),
		},
		ts.TestCase{
			Name:     "check not dependent number form from hundreds, dozens and other triplet (they are incorrect)",
			Args:     ts.TTA([]coreobjects.RuneDigitTriplet{{'a', 'a', 'a'}, {'a', 'a', '1'}}),
			Expected: ts.TTVE(constants.FIRST_FORM),
		},
		ts.TestCase{
			Name:     "empty triplets",
			Args:     ts.TTA([]coreobjects.RuneDigitTriplet{}),
			Expected: ts.TTVE(constants.THIRD_FORM),
		},
		ts.TestCase{
			Name:     "nil triplets",
			Args:     ts.TTA([]coreobjects.RuneDigitTriplet(nil)),
			Expected: ts.TTVE(constants.THIRD_FORM),
		},
	)
}

func (r *RussianUtilityFunctionsSuite) TestGetCurrencyIntegerPartWordFormFunction() {
	r.RunTest(
		getCurrencyIntegerPartWordForm,
		ts.TestCase{
			Name:     "currency will be number",
			Args:     ts.TTA(true, constants.INVALID_FORM),
			Expected: ts.TTVE(constants.PLURAL_WORD),
		},
		ts.TestCase{
			Name:     "number is in second form",
			Args:     ts.TTA(false, constants.SECOND_FORM),
			Expected: ts.TTVE(constants.SINGULAR_WORD),
		},
		ts.TestCase{
			Name:     "number is in first form",
			Args:     ts.TTA(false, constants.FIRST_FORM),
			Expected: ts.TTVE(constants.PLURAL_WORD),
		},
		ts.TestCase{
			Name:     "number is in third form",
			Args:     ts.TTA(false, constants.THIRD_FORM),
			Expected: ts.TTVE(constants.PLURAL_WORD),
		},
	)
}

func (r *RussianUtilityFunctionsSuite) TestGetCurrencyFractionalPartWordFormFunction() {
	r.RunTest(
		getCurrencyFractionalPartWordForm,
		ts.TestCase{
			Name:     "number is in first form",
			Args:     ts.TTA(constants.FIRST_FORM),
			Expected: ts.TTVE(constants.SINGULAR_WORD),
		},
		ts.TestCase{
			Name:     "number is in second form",
			Args:     ts.TTA(constants.SECOND_FORM),
			Expected: ts.TTVE(constants.PLURAL_WORD),
		},
		ts.TestCase{
			Name:     "number is in third form",
			Args:     ts.TTA(constants.THIRD_FORM),
			Expected: ts.TTVE(constants.PLURAL_WORD),
		},
	)
}

func (r *RussianUtilityFunctionsSuite) TestGetCurrencyFractionalPartDeclensionFunction() {
	r.RunTest(
		getCurrencyFractionalPartDeclension,
		ts.TestCase{
			Name:     "number is in second form and NOMINATIVE",
			Args:     ts.TTA(declension.NOMINATIVE, constants.SECOND_FORM),
			Expected: ts.TTVE(declension.GENITIVE),
		},
		ts.TestCase{
			Name:     "number is in second form and ACCUSATIVE",
			Args:     ts.TTA(declension.ACCUSATIVE, constants.SECOND_FORM),
			Expected: ts.TTVE(declension.GENITIVE),
		},
		ts.TestCase{
			Name:     "number is in third form and NOMINATIVE",
			Args:     ts.TTA(declension.NOMINATIVE, constants.THIRD_FORM),
			Expected: ts.TTVE(declension.GENITIVE),
		},
		ts.TestCase{
			Name:     "number is in third form and ACCUSATIVE",
			Args:     ts.TTA(declension.ACCUSATIVE, constants.THIRD_FORM),
			Expected: ts.TTVE(declension.GENITIVE),
		},
		ts.TestCase{
			Name:     "number is in first form and ACCUSATIVE",
			Args:     ts.TTA(declension.ACCUSATIVE, constants.FIRST_FORM),
			Expected: ts.TTVE(declension.ACCUSATIVE),
		},
		ts.TestCase{
			Name:     "number is in first form and NOMINATIVE",
			Args:     ts.TTA(declension.NOMINATIVE, constants.FIRST_FORM),
			Expected: ts.TTVE(declension.NOMINATIVE),
		},
		ts.TestCase{
			Name:     "number is in second form and DATIVE",
			Args:     ts.TTA(declension.DATIVE, constants.SECOND_FORM),
			Expected: ts.TTVE(declension.DATIVE),
		},
		ts.TestCase{
			Name:     "number is in third form and DATIVE",
			Args:     ts.TTA(declension.DATIVE, constants.THIRD_FORM),
			Expected: ts.TTVE(declension.DATIVE),
		},
	)
}

func (r *RussianUtilityFunctionsSuite) TestConvertDigitToWordFunction() {
	jsonDeclensionNumbers := strings.NewReader("{\"nominative\":[ \"ноль\", { \"male\": \"один\", \"neuter\": \"одно\", \"female\": \"одна\"}," +
		" { \"male\": \"два\", \"neuter\": \"два\", \"female\": \"две\"}, " +
		"\"три\", \"четыре\", \"пять\", \"шесть\", \"семь\", \"восемь\", \"девять\"]}")
	decoder := json.NewDecoder(jsonDeclensionNumbers)

	declensionNumbers := objects.DeclensionNumbers{}
	require.NoError(r.T(), decoder.Decode(&declensionNumbers))

	r.RunTest(
		convertDigitToWord,
		ts.TestCase{
			Name:     "word with gender",
			Args:     ts.TTA(1, declensionNumbers, declension.NOMINATIVE, genders.FEMALE),
			Expected: ts.TTVE("одна"),
		},
		ts.TestCase{
			Name:     "word without gender",
			Args:     ts.TTA(3, declensionNumbers, declension.NOMINATIVE, genders.FEMALE),
			Expected: ts.TTVE("три"),
		},
	)
}

func TestRussianUtilityFunctionsSuite(t *testing.T) {
	suite.Run(t, new(RussianUtilityFunctionsSuite))
}
