package stringutilits

import (
	"testing"

	"github.com/stretchr/testify/suite"

	ts "github.com/ThCompiler/go_game_constractor/pkg/testing"
)

type StringFormatSuite struct {
	ts.TestCasesSuite
	RunFunc func(args ...interface{}) []interface{}
}

func (s *StringFormatSuite) SetupTest() {
	s.RunFunc = func(args ...interface{}) []interface{} {
		res := StringFormat(args[0].(string), args[1:]...)
		return []interface{}{res}
	}
}

func (s *StringFormatSuite) TestNoArguments() {
	s.RunTest(
		s.RunFunc,
		ts.TestCase{
			Name:     "Correct",
			Args:     ts.ToTestArgs("Hello world"),
			Expected: ts.ToTestValuesExpected("Hello world"),
		},
		ts.TestCase{
			Name:     "Required fields",
			Args:     ts.ToTestArgs("Hello. {name}"),
			Expected: ts.ToTestValuesExpected("Hello. {name}"),
		},
	)
}

func (s *StringFormatSuite) TestInCorrect() {
	s.RunTest(
		s.RunFunc,
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

func TestStringFormatSuite(t *testing.T) {
	suite.Run(t, new(StringFormatSuite))
}
