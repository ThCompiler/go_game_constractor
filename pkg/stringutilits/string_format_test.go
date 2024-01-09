package stringutilits

import (
	"github.com/ThCompiler/ts"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StringFormatSuite struct {
	ts.TestCasesSuite
	ActFunc func(args ...interface{}) string
}

func (s *StringFormatSuite) SetupTest() {
	s.ActFunc = func(args ...interface{}) string {
		res := StringFormat(args[0].(string), args[1:]...)
		return res
	}
}

func (s *StringFormatSuite) TestCorrectNoFields() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name:     "No Arguments",
			Args:     ts.TTA("Hello world"),
			Expected: ts.TTVE("Hello world"),
		},
		ts.TestCase{
			Name:     "Required fields",
			Args:     ts.TTA("Hello. {name}"),
			Expected: ts.TTVE("Hello. {name}"),
		},
	)
}

func (s *StringFormatSuite) TestCorrectWithFields() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name:     "With two field in order 'greeting' and 'name'",
			Args:     ts.TTA("{greeting}. {name}", "greeting", "Hello", "name", "ThCompiler"),
			Expected: ts.TTVE("Hello. ThCompiler"),
		},
		ts.TestCase{
			Name:     "With two field in order 'name' and 'greeting'",
			Args:     ts.TTA("{greeting}. {name}", "name", "ThCompiler", "greeting", "Hello"),
			Expected: ts.TTVE("Hello. ThCompiler"),
		},
		ts.TestCase{
			Name:     "With one field",
			Args:     ts.TTA("Hello. {name}", "name", "ThCompiler"),
			Expected: ts.TTVE("Hello. ThCompiler"),
		},
	)
}

func (s *StringFormatSuite) TestInCorrect() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name:     "Without value of field",
			Args:     ts.TTA("Hello. {name}", "name"),
			Expected: ts.TTPEE("strings.NewReplacer: odd argument count"),
		},
	)
}

func TestStringFormatSuite(t *testing.T) {
	suite.Run(t, new(StringFormatSuite))
}
