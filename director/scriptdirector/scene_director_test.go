package scriptdirector

import (
	ts "github.com/ThCompiler/go_game_constractor/pkg/testing"
)

type SceneDirectorSuite struct {
	ts.TestCasesSuite
	RunFunc func(args ...interface{}) []interface{}
}

/*func (s *SceneDirectorSuite) SetupTest() {
	s.RunFunc = func(args ...interface{}) []interface{} {
		res := ClearStringFromPunctuation(args[0].(string))
		return []interface{}{res}
	}
}

func (s *SceneDirectorSuite) TestStrWithPunctuation() {
	s.RunTest(
		s.RunFunc,
		ts.TestCase{
			Name:     "All punctuation characters",
			Args:     ts.TTA(`!\"#$%&'()*+,-./:;<=>?@[\\]^_` + "`" + `{|}~`),
			Expected: ts.TTVE(""),
		},
		ts.TestCase{
			Name:     "All punctuation characters with spaces",
			Args:     ts.TTA(` !\"#$%&'()*+,-./:;<=>?@[\\]^_` + "`" + `{|}~ `),
			Expected: ts.TTVE("  "),
		},
		ts.TestCase{
			Name:     "All punctuation characters with spaces and text",
			Args:     ts.TTA(` !\"#$%&'()*+,-./:;<=>?@[\\]^_` + "`" + `{|}~ text `),
			Expected: ts.TTVE("  text "),
		},
		ts.TestCase{
			Name:     "All punctuation characters with spaces and text",
			Args:     ts.TTA(` !\"#$%&'()*+,-./:;<=>?@[\\]^_` + "`" + `{|}~ text !\"#$%&'()*+,-./:;<=>?@[\\]^_` + "`" + `{|}~ `),
			Expected: ts.TTVE("  text  "),
		},
		ts.TestCase{
			Name: "All punctuation characters with spaces and text between",
			Args: ts.TTA(`text !\"#$%&'()*+,-./:;<=>?@[\\]^_` + "`" + `{|}~ text !\"#$%&'()*+,` +
				`-./:;<=>?@[\\]^_` + "`" + `{|}~ text`),
			Expected: ts.TTVE("text  text  text"),
		},
	)
}

func (s *ClearStringFromPunctuationSuite) TestStrWithoutPunctuation() {
	s.RunTest(
		s.RunFunc,
		ts.TestCase{
			Name:     "Single word",
			Args:     ts.TTA(`text`),
			Expected: ts.TTVE("text"),
		},
		ts.TestCase{
			Name:     "Single word with spaces",
			Args:     ts.TTA(` text `),
			Expected: ts.TTVE(" text "),
		},
		ts.TestCase{
			Name:     "Multiple words",
			Args:     ts.TTA(`text text text`),
			Expected: ts.TTVE("text text text"),
		},
		ts.TestCase{
			Name:     "Multiple words with spaces",
			Args:     ts.TTA(` text text text `),
			Expected: ts.TTVE(" text text text "),
		},
	)
}

func TestClearStringFromPunctuationSuite(t *testing.T) {
	suite.Run(t, new(ClearStringFromPunctuationSuite))
}*/
