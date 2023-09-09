package stringutilits

import (
	"testing"

	"github.com/stretchr/testify/suite"

	ts "github.com/ThCompiler/go_game_constractor/pkg/testing"
)

type ToDigitSuite struct {
	ts.TestCasesSuite
	ActFunc func(args ...interface{}) []interface{}
}

func (s *ToDigitSuite) SetupTest() {
	s.ActFunc = func(args ...interface{}) []interface{} {
		res := ToDigit(args[0].(rune))
		return []interface{}{res}
	}
}

func (s *ToDigitSuite) TestCorrect() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name:     "0",
			Args:     ts.ToTestArgs('0'),
			Expected: ts.ToTestValuesExpected(0),
		},
		ts.TestCase{
			Name:     "1",
			Args:     ts.ToTestArgs('1'),
			Expected: ts.ToTestValuesExpected(1),
		},
		ts.TestCase{
			Name:     "2",
			Args:     ts.ToTestArgs('2'),
			Expected: ts.ToTestValuesExpected(2),
		},
		ts.TestCase{
			Name:     "3",
			Args:     ts.ToTestArgs('3'),
			Expected: ts.ToTestValuesExpected(3),
		},
		ts.TestCase{
			Name:     "4",
			Args:     ts.ToTestArgs('4'),
			Expected: ts.ToTestValuesExpected(4),
		},
		ts.TestCase{
			Name:     "5",
			Args:     ts.ToTestArgs('5'),
			Expected: ts.ToTestValuesExpected(5),
		},
		ts.TestCase{
			Name:     "6",
			Args:     ts.ToTestArgs('6'),
			Expected: ts.ToTestValuesExpected(6),
		},
		ts.TestCase{
			Name:     "7",
			Args:     ts.ToTestArgs('7'),
			Expected: ts.ToTestValuesExpected(7),
		},
		ts.TestCase{
			Name:     "8",
			Args:     ts.ToTestArgs('8'),
			Expected: ts.ToTestValuesExpected(8),
		},
		ts.TestCase{
			Name:     "9",
			Args:     ts.ToTestArgs('9'),
			Expected: ts.ToTestValuesExpected(9),
		},
	)
}

func (s *ToDigitSuite) TestInCorrect() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name:     "character",
			Args:     ts.ToTestArgs('h'),
			Expected: ts.ToTestValuesExpected(-1),
		},
		ts.TestCase{
			Name:     "special character",
			Args:     ts.ToTestArgs('\t'),
			Expected: ts.ToTestValuesExpected(-1),
		},
	)
}

type ToRuneSuite struct {
	ts.TestCasesSuite
	ActFunc func(args ...interface{}) []interface{}
}

func (s *ToRuneSuite) SetupTest() {
	s.ActFunc = func(args ...interface{}) []interface{} {
		res := ToRune(args[0].(int8))
		return []interface{}{res}
	}
}

func (s *ToRuneSuite) TestCorrect() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name:     "0",
			Args:     ts.ToTestArgs(int8(0)),
			Expected: ts.ToTestValuesExpected('0'),
		},
		ts.TestCase{
			Name:     "1",
			Args:     ts.ToTestArgs(int8(1)),
			Expected: ts.ToTestValuesExpected('1'),
		},
		ts.TestCase{
			Name:     "2",
			Args:     ts.ToTestArgs(int8(2)),
			Expected: ts.ToTestValuesExpected('2'),
		},
		ts.TestCase{
			Name:     "3",
			Args:     ts.ToTestArgs(int8(3)),
			Expected: ts.ToTestValuesExpected('3'),
		},
		ts.TestCase{
			Name:     "4",
			Args:     ts.ToTestArgs(int8(4)),
			Expected: ts.ToTestValuesExpected('4'),
		},
		ts.TestCase{
			Name:     "5",
			Args:     ts.ToTestArgs(int8(5)),
			Expected: ts.ToTestValuesExpected('5'),
		},
		ts.TestCase{
			Name:     "6",
			Args:     ts.ToTestArgs(int8(6)),
			Expected: ts.ToTestValuesExpected('6'),
		},
		ts.TestCase{
			Name:     "7",
			Args:     ts.ToTestArgs(int8(7)),
			Expected: ts.ToTestValuesExpected('7'),
		},
		ts.TestCase{
			Name:     "8",
			Args:     ts.ToTestArgs(int8(8)),
			Expected: ts.ToTestValuesExpected('8'),
		},
		ts.TestCase{
			Name:     "9",
			Args:     ts.ToTestArgs(int8(9)),
			Expected: ts.ToTestValuesExpected('9'),
		},
	)
}

func (s *ToRuneSuite) TestInCorrect() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name:     "positive number",
			Args:     ts.ToTestArgs(int8(30)),
			Expected: ts.ToTestValuesExpected(rune(0)),
		},
		ts.TestCase{
			Name:     "negative number",
			Args:     ts.ToTestArgs(int8(-20)),
			Expected: ts.ToTestValuesExpected(rune(0)),
		},
	)
}

func TestToDigitSuite(t *testing.T) {
	suite.Run(t, new(ToDigitSuite))
}

func TestToRuneSuite(t *testing.T) {
	suite.Run(t, new(ToRuneSuite))
}
