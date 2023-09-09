package functions

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/num2words/objects"
	ts "github.com/ThCompiler/go_game_constractor/pkg/testing"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RoundDigitSuite struct {
	ts.TestCasesSuite
	ActFunc func(args ...interface{}) []interface{}
}

func (s *RoundDigitSuite) SetupTest() {
	s.ActFunc = func(args ...interface{}) []interface{} {
		increase, res := roundDigit(args[0].(bool), args[1].([]rune), args[2].(int))
		return []interface{}{increase, res}
	}
}

func (s *RoundDigitSuite) TestWithRequireIncreasing() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name:     "Digit is zero",
			Args:     ts.TTA(true, []rune("98765432109"), 9),
			Expected: ts.TTVE(false, []rune("98765432119")),
		},
		ts.TestCase{
			Name:     "Digit is one",
			Args:     ts.TTA(true, []rune("98765432109"), 8),
			Expected: ts.TTVE(false, []rune("98765432209")),
		},
		ts.TestCase{
			Name:     "Digit is four",
			Args:     ts.TTA(true, []rune("98765432109"), 5),
			Expected: ts.TTVE(false, []rune("98765532109")),
		},
		ts.TestCase{
			Name:     "Digit is eight",
			Args:     ts.TTA(true, []rune("98765432109"), 1),
			Expected: ts.TTVE(false, []rune("99765432109")),
		},
		ts.TestCase{
			Name:     "Digit is nine",
			Args:     ts.TTA(true, []rune("98765432109"), 10),
			Expected: ts.TTVE(true, []rune("98765432100")),
		},
		ts.TestCase{
			Name:     "Digit is nine and first digit in number",
			Args:     ts.TTA(true, []rune("98765432109"), 0),
			Expected: ts.TTVE(true, []rune("108765432109")),
		},
	)
}

func (s *RoundDigitSuite) TestWithoutRequireIncreasing() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name:     "Digit is zero",
			Args:     ts.TTA(false, []rune("98765432109"), 9),
			Expected: ts.TTVE(false, []rune("98765432109")),
		},
		ts.TestCase{
			Name:     "Digit is one",
			Args:     ts.TTA(false, []rune("98765432109"), 8),
			Expected: ts.TTVE(false, []rune("98765432109")),
		},
		ts.TestCase{
			Name:     "Digit is four",
			Args:     ts.TTA(false, []rune("98765432109"), 5),
			Expected: ts.TTVE(false, []rune("98765432109")),
		},
		ts.TestCase{
			Name:     "Digit is five",
			Args:     ts.TTA(false, []rune("98765432109"), 4),
			Expected: ts.TTVE(true, []rune("98765432109")),
		},
		ts.TestCase{
			Name:     "Digit is nine",
			Args:     ts.TTA(false, []rune("98765432109"), 10),
			Expected: ts.TTVE(true, []rune("98765432109")),
		},
		ts.TestCase{
			Name:     "Digit is nine and first digit in number",
			Args:     ts.TTA(false, []rune("98765432109"), 0),
			Expected: ts.TTVE(true, []rune("98765432109")),
		},
	)
}

func TestRoundDigitSuite(t *testing.T) {
	suite.Run(t, new(RoundDigitSuite))
}

type RoundNumberSuite struct {
	ts.TestCasesSuite
	ActFunc func(args ...interface{}) []interface{}
}

func (s *RoundNumberSuite) SetupTest() {
	s.ActFunc = func(args ...interface{}) []interface{} {
		res := RoundNumber(args[0].(objects.Number), args[1].(int64))
		return []interface{}{res}
	}
}

func (s *RoundNumberSuite) TestWithoutRound() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name:     "Number is fractional",
			Args:     ts.TTA(objects.Number{Divider: constants.FractionalNumber}, int64(0)),
			Expected: ts.TTVE(objects.Number{Divider: constants.FractionalNumber}),
		},
		ts.TestCase{
			Name:     "Precision is negative",
			Args:     ts.TTA(objects.Number{Divider: constants.DecimalNumber}, int64(-1)),
			Expected: ts.TTVE(objects.Number{Divider: constants.DecimalNumber}),
		},
		ts.TestCase{
			Name: "Precision more then count numbers",
			Args: ts.TTA(objects.Number{
				Divider:    constants.DecimalNumber,
				SecondPart: "202",
			}, int64(5)),
			Expected: ts.TTVE(objects.Number{
				Divider:    constants.DecimalNumber,
				SecondPart: "202",
			}),
		},
		ts.TestCase{
			Name: "Precision equal count numbers",
			Args: ts.TTA(objects.Number{
				Divider:    constants.DecimalNumber,
				SecondPart: "202",
			}, int64(3)),
			Expected: ts.TTVE(objects.Number{
				Divider:    constants.DecimalNumber,
				SecondPart: "202",
			}),
		},
		ts.TestCase{
			Name: "Precision is zero and equal count numbers",
			Args: ts.TTA(objects.Number{
				Divider:    constants.DecimalNumber,
				SecondPart: "",
			}, int64(0)),
			Expected: ts.TTVE(objects.Number{
				Divider:    constants.DecimalNumber,
				SecondPart: "",
			}),
		},
	)
}

func (s *RoundNumberSuite) TestWithRound() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name: "Digit for round is less than five",
			Args: ts.TTA(objects.Number{
				Divider:    constants.DecimalNumber,
				SecondPart: "202",
			}, int64(2)),
			Expected: ts.TTVE(objects.Number{
				Divider:    constants.DecimalNumber,
				SecondPart: "2",
			}),
		},
		ts.TestCase{
			Name: "Digit for round is not less than five",
			Args: ts.TTA(objects.Number{
				Divider:    constants.DecimalNumber,
				SecondPart: "206",
			}, int64(2)),
			Expected: ts.TTVE(objects.Number{
				Divider:    constants.DecimalNumber,
				SecondPart: "21",
			}),
		},
		ts.TestCase{
			Name: "All digits is nine in decimal part",
			Args: ts.TTA(objects.Number{
				Divider:    constants.DecimalNumber,
				FirstPart:  "1",
				SecondPart: "99999",
			}, int64(3)),
			Expected: ts.TTVE(objects.Number{
				Divider:    constants.DecimalNumber,
				FirstPart:  "2",
				SecondPart: "0",
			}),
		},
		ts.TestCase{
			Name: "All digits is nine in decimal and integer part",
			Args: ts.TTA(objects.Number{
				Divider:    constants.DecimalNumber,
				FirstPart:  "2999",
				SecondPart: "99999",
			}, int64(3)),
			Expected: ts.TTVE(objects.Number{
				Divider:    constants.DecimalNumber,
				FirstPart:  "3000",
				SecondPart: "0",
			}),
		},
		ts.TestCase{
			Name: "All digits is nine in decimal part without integer part",
			Args: ts.TTA(objects.Number{
				Divider:    constants.DecimalNumber,
				SecondPart: "99999",
			}, int64(3)),
			Expected: ts.TTVE(objects.Number{
				Divider:    constants.DecimalNumber,
				FirstPart:  "1",
				SecondPart: "0",
			}),
		},
	)
}

func TestRoundNumberSuite(t *testing.T) {
	suite.Run(t, new(RoundNumberSuite))
}
