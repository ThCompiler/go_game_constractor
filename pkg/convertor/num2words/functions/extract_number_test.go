package functions

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/num2words/objects"
	ts "github.com/ThCompiler/go_game_constractor/pkg/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

//----------------------------------------------------------------------------------------------------------------------
// Tests suite for "RemoveFromString" function
//----------------------------------------------------------------------------------------------------------------------

type RemoveFromStringSuite struct {
	suite.Suite
}

func (s *RemoveFromStringSuite) TestEmptyRemoveString() {
	res := RemoveFromString("word", "")
	assert.Equal(s.T(), "word", res)
}

func (s *RemoveFromStringSuite) TestEmptyBothString() {
	res := RemoveFromString("", "")
	assert.Equal(s.T(), "", res)
}

func (s *RemoveFromStringSuite) TestEmptyModifiedString() {
	res := RemoveFromString("", "word")
	assert.Equal(s.T(), "", res)
}

func (s *RemoveFromStringSuite) TestModifiedStringNotContainsRemoveString() {
	res := RemoveFromString("word", "sok")
	assert.Equal(s.T(), "word", res)
}

func (s *RemoveFromStringSuite) TestModifiedStringContainsRemoveString() {
	res := RemoveFromString("word", "rd")
	assert.Equal(s.T(), "wo", res)
}

func (s *RemoveFromStringSuite) TestModifiedStringContainsManyTimesRemoveString() {
	res := RemoveFromString("word can word be word", "rd")
	assert.Equal(s.T(), "wo can wo be wo", res)
}

func TestRemoveFromStringSuite(t *testing.T) {
	suite.Run(t, new(RemoveFromStringSuite))
}

//----------------------------------------------------------------------------------------------------------------------
// Tests suite for "ReplaceInString" function
//----------------------------------------------------------------------------------------------------------------------

type ReplaceInStringSuite struct {
	suite.Suite
}

func (s *ReplaceInStringSuite) TestNotEmptyModifiedStringAndEmptyRegexStringAndNotEmptyReplaceString() {
	res := ReplaceInString("word", "", "d")
	assert.Equal(s.T(), "word", res)
}

func (s *ReplaceInStringSuite) TestNotEmptyModifiedStringAndEmptyRegexStringAndEmptyReplaceString() {
	res := ReplaceInString("word", "", "")
	assert.Equal(s.T(), "word", res)
}

func (s *ReplaceInStringSuite) TestEmptyModifiedStringAndNotEmptyRegexStringAndEmptyReplaceString() {
	res := ReplaceInString("", "rd", "")
	assert.Equal(s.T(), "", res)
}

func (s *ReplaceInStringSuite) TestEmptyModifiedStringAndNotEmptyRegexStringAndNotEmptyReplaceString() {
	res := ReplaceInString("", "rd", "r")
	assert.Equal(s.T(), "", res)
}

func (s *ReplaceInStringSuite) TestEmptyModifiedStringAndEmptyRegexStringAndEmptyReplaceString() {
	res := ReplaceInString("", "", "")
	assert.Equal(s.T(), "", res)
}

func (s *ReplaceInStringSuite) TestEmptyModifiedStringAndEmptyRegexStringAndNotEmptyReplaceString() {
	res := ReplaceInString("", "", "r")
	assert.Equal(s.T(), "", res)
}

func (s *ReplaceInStringSuite) TestNotEmptyModifiedStringAndRegexContainsInStringAndEmptyReplaceString() {
	res := ReplaceInString("word", "wo", "")
	assert.Equal(s.T(), "rd", res)
}

func (s *ReplaceInStringSuite) TestNotEmptyModifiedStringAndRegexContainsInStringAndNotEmptyReplaceString() {
	res := ReplaceInString("word", "wo", "r")
	assert.Equal(s.T(), "rrd", res)
}

func (s *ReplaceInStringSuite) TestNotEmptyModifiedStringAndRegexManyContainsInStringAndEmptyReplaceString() {
	res := ReplaceInString("word can word be word", "wo", "")
	assert.Equal(s.T(), "rd can rd be rd", res)
}

func (s *ReplaceInStringSuite) TestNotEmptyModifiedStringAndRegexManyContainsInStringAndNotEmptyReplaceString() {
	res := ReplaceInString("word can word be word", "wo", "r")
	assert.Equal(s.T(), "rrd can rrd be rrd", res)
}

func (s *ReplaceInStringSuite) TestNotEmptyModifiedStringAndRegexNotContainsInStringAndEmptyReplaceString() {
	res := ReplaceInString("word", "ko", "")
	assert.Equal(s.T(), "word", res)
}

func (s *ReplaceInStringSuite) TestNotEmptyModifiedStringAndRegexNotContainsInStringAndNotEmptyReplaceString() {
	res := ReplaceInString("word", "ko", "r")
	assert.Equal(s.T(), "word", res)
}

func TestReplaceInStringSuite(t *testing.T) {
	suite.Run(t, new(ReplaceInStringSuite))
}

//----------------------------------------------------------------------------------------------------------------------
// Tests suite for "ReplaceInString" function
//----------------------------------------------------------------------------------------------------------------------

type ValidateNumberSuite struct {
	ts.TestCasesSuite
	ActFunc func(args ...interface{}) []interface{}
}

func (s *ValidateNumberSuite) SetupTest() {
	s.ActFunc = func(args ...interface{}) []interface{} {
		res := ValidateNumber(args[0].(string))
		return []interface{}{res}
	}
}

func (s *ValidateNumberSuite) TestCorrectNumber() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name:     "Empty",
			Args:     ts.TTA(``),
			Expected: ts.TTEE(nil),
		},
		ts.TestCase{
			Name:     "Zero",
			Args:     ts.TTA(`0`),
			Expected: ts.TTEE(nil),
		},
		ts.TestCase{
			Name:     "Positive number",
			Args:     ts.TTA(`12`),
			Expected: ts.TTEE(nil),
		},
		ts.TestCase{
			Name:     "Positive number with plus",
			Args:     ts.TTA(`+12`),
			Expected: ts.TTEE(nil),
		},
		ts.TestCase{
			Name:     "Negative number",
			Args:     ts.TTA(`-12`),
			Expected: ts.TTEE(nil),
		},
		ts.TestCase{
			Name:     "Fractional positive number",
			Args:     ts.TTA(`12/12`),
			Expected: ts.TTEE(nil),
		},
		ts.TestCase{
			Name:     "Fractional positive number with plus",
			Args:     ts.TTA(`+12/12`),
			Expected: ts.TTEE(nil),
		},
		ts.TestCase{
			Name:     "Fractional negative number",
			Args:     ts.TTA(`-12/12`),
			Expected: ts.TTEE(nil),
		},
		ts.TestCase{
			Name:     "Decimal positive number",
			Args:     ts.TTA(`12.2`),
			Expected: ts.TTEE(nil),
		},
		ts.TestCase{
			Name:     "Decimal positive number with plus",
			Args:     ts.TTA(`+12.2`),
			Expected: ts.TTEE(nil),
		},
		ts.TestCase{
			Name:     "Decimal negative number",
			Args:     ts.TTA(`-12.2`),
			Expected: ts.TTEE(nil),
		},
		ts.TestCase{
			Name:     "Decimal number with empty decimal part",
			Args:     ts.TTA(`12.`),
			Expected: ts.TTEE(nil),
		},
		ts.TestCase{
			Name:     "Number with zeros before number",
			Args:     ts.TTA(`000012`),
			Expected: ts.TTEE(nil),
		},
		ts.TestCase{
			Name:     "Number with spaces",
			Args:     ts.TTA(` 1 2 `),
			Expected: ts.TTEE(nil),
		},
	)
}

func (s *ValidateNumberSuite) TestIncorrectNumber() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name:     "Incorrect chars",
			Args:     ts.TTA(` !\"#$%&'()*+,-./:;<=>?@[\\]^_` + "`" + `{|}~ `),
			Expected: ts.TTEE(ErrorNumberContainsIncorrectChars),
		},
		ts.TestCase{
			Name:     "Number with many plus",
			Args:     ts.TTA(`+1+31+`),
			Expected: ts.TTEE(ErrorNumberHaveManySignChars),
		},
		ts.TestCase{
			Name:     "Number with many minus",
			Args:     ts.TTA(`-1-31-`),
			Expected: ts.TTEE(ErrorNumberHaveManySignChars),
		},
		ts.TestCase{
			Name:     "Number with minus not at beginning of number",
			Args:     ts.TTA(`131-`),
			Expected: ts.TTEE(ErrorNumberHaveSignCharNotInBegin),
		},
		ts.TestCase{
			Name:     "Number with plus not at beginning of number",
			Args:     ts.TTA(`131+`),
			Expected: ts.TTEE(ErrorNumberHaveSignCharNotInBegin),
		},
		ts.TestCase{
			Name:     "Number with many delimiter point",
			Args:     ts.TTA(`13..1`),
			Expected: ts.TTEE(ErrorNumberHaveManyDelimiterChars),
		},
		ts.TestCase{
			Name:     "Number with many delimiter comma",
			Args:     ts.TTA(`13,,1`),
			Expected: ts.TTEE(ErrorNumberHaveManyDelimiterChars),
		},
		ts.TestCase{
			Name:     "Number with many delimiter slash",
			Args:     ts.TTA(`13//1`),
			Expected: ts.TTEE(ErrorNumberHaveManyDelimiterChars),
		},
	)
}

func TestValidateNumberSuite(t *testing.T) {
	suite.Run(t, new(ValidateNumberSuite))
}

//----------------------------------------------------------------------------------------------------------------------
// Tests suite for "ExtractNumber" function
//----------------------------------------------------------------------------------------------------------------------

type ExtractNumberSuite struct {
	ts.TestCasesSuite
	ActFunc func(args ...interface{}) []interface{}
}

func (s *ExtractNumberSuite) SetupTest() {
	s.ActFunc = func(args ...interface{}) []interface{} {
		res, err := ExtractNumber(args[0].(string))
		return []interface{}{res, err}
	}
}

func (s *ExtractNumberSuite) TestIncorrectNumber() {
	_, err := ExtractNumber("d12")
	assert.Error(s.T(), err)
}

func (s *ExtractNumberSuite) TestEmptyNumber() {
	res, err := ExtractNumber("")
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), objects.Number{
		Divider:    constants.DECIMAL_NUMBER,
		FirstPart:  "0",
		SecondPart: "0",
		Sign:       "+",
	}, res)
}

func (s *ExtractNumberSuite) TestBaseNumber() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name: "Only number",
			Args: ts.TTA(`12`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "12",
				SecondPart: "0",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number with minus",
			Args: ts.TTA(`-12`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "12",
				SecondPart: "0",
				Sign:       "-",
			}),
		},
		ts.TestCase{
			Name: "Number with plus",
			Args: ts.TTA(`+12`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "12",
				SecondPart: "0",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number with spaces",
			Args: ts.TTA(` 1 2 `),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "12",
				SecondPart: "0",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Zero",
			Args: ts.TTA(`0`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "0",
				SecondPart: "0",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number with zeros in beginning of number",
			Args: ts.TTA(`00012`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "12",
				SecondPart: "0",
				Sign:       "+",
			}),
		},
	)
}

func (s *ExtractNumberSuite) TestDecimalNumberWithPoint() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name: "Only number",
			Args: ts.TTA(`1.2`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "1",
				SecondPart: "2",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number with minus",
			Args: ts.TTA(`-1.2`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "1",
				SecondPart: "2",
				Sign:       "-",
			}),
		},
		ts.TestCase{
			Name: "Number with plus",
			Args: ts.TTA(`+1.2`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "1",
				SecondPart: "2",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number with spaces",
			Args: ts.TTA(` 1 . 2 `),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "1",
				SecondPart: "2",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Zero",
			Args: ts.TTA(`0.0`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "0",
				SecondPart: "0",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number with zeros in integer part of number",
			Args: ts.TTA(`00012000.2`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "12000",
				SecondPart: "2",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number with zeros in decimal part of number",
			Args: ts.TTA(`12.00020000`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "12",
				SecondPart: "0002",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number without decimal part",
			Args: ts.TTA(`12.`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "12",
				SecondPart: "0",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number without integer part",
			Args: ts.TTA(`.2`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "0",
				SecondPart: "2",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number without integer and decimal part",
			Args: ts.TTA(`.`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "0",
				SecondPart: "0",
				Sign:       "+",
			}),
		},
	)
}

func (s *ExtractNumberSuite) TestDecimalNumberWithComma() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name: "Only number",
			Args: ts.TTA(`1,2`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "1",
				SecondPart: "2",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number with minus",
			Args: ts.TTA(`-1,2`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "1",
				SecondPart: "2",
				Sign:       "-",
			}),
		},
		ts.TestCase{
			Name: "Number with plus",
			Args: ts.TTA(`+1,2`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "1",
				SecondPart: "2",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number with spaces",
			Args: ts.TTA(` 1 , 2 `),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "1",
				SecondPart: "2",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Zero",
			Args: ts.TTA(`0,0`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "0",
				SecondPart: "0",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number with zeros in integer part of number",
			Args: ts.TTA(`00012000,2`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "12000",
				SecondPart: "2",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number with zeros in decimal part of number",
			Args: ts.TTA(`12,00020000`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "12",
				SecondPart: "0002",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number without decimal part",
			Args: ts.TTA(`12,`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "12",
				SecondPart: "0",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number without integer part",
			Args: ts.TTA(`,2`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "0",
				SecondPart: "2",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number without integer and decimal part",
			Args: ts.TTA(`,`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.DECIMAL_NUMBER,
				FirstPart:  "0",
				SecondPart: "0",
				Sign:       "+",
			}),
		},
	)
}

func (s *ExtractNumberSuite) TestFractionalNumber() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name: "Only number",
			Args: ts.TTA(`1/2`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.FRACTIONAL_NUMBER,
				FirstPart:  "1",
				SecondPart: "2",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number with minus",
			Args: ts.TTA(`-1/2`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.FRACTIONAL_NUMBER,
				FirstPart:  "1",
				SecondPart: "2",
				Sign:       "-",
			}),
		},
		ts.TestCase{
			Name: "Number with plus",
			Args: ts.TTA(`+1/2`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.FRACTIONAL_NUMBER,
				FirstPart:  "1",
				SecondPart: "2",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number with spaces",
			Args: ts.TTA(` 1 / 2 `),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.FRACTIONAL_NUMBER,
				FirstPart:  "1",
				SecondPart: "2",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Zero",
			Args: ts.TTA(`0/0`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.FRACTIONAL_NUMBER,
				FirstPart:  "0",
				SecondPart: "0",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number with zeros in numerator",
			Args: ts.TTA(`00012000/2`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.FRACTIONAL_NUMBER,
				FirstPart:  "12000",
				SecondPart: "2",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number with zeros in denominator",
			Args: ts.TTA(`12/00020000`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.FRACTIONAL_NUMBER,
				FirstPart:  "12",
				SecondPart: "20000",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number without denominator",
			Args: ts.TTA(`2/`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.FRACTIONAL_NUMBER,
				FirstPart:  "2",
				SecondPart: "0",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number without numerator",
			Args: ts.TTA(`/2`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.FRACTIONAL_NUMBER,
				FirstPart:  "0",
				SecondPart: "2",
				Sign:       "+",
			}),
		},
		ts.TestCase{
			Name: "Number without numerator and denominator",
			Args: ts.TTA(`/`),
			Expected: ts.TTVEWNE(objects.Number{
				Divider:    constants.FRACTIONAL_NUMBER,
				FirstPart:  "0",
				SecondPart: "0",
				Sign:       "+",
			}),
		},
	)
}

func TestExtractNumberSuite(t *testing.T) {
	suite.Run(t, new(ExtractNumberSuite))
}
