package stringutilits

import (
	"github.com/ThCompiler/ts"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ClearStringFromPunctuationSuite struct {
	ts.TestCasesSuite
}

func (s *ClearStringFromPunctuationSuite) TestStrWithPunctuation() {
	s.RunTest(
		ClearStringFromPunctuation,
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
		ClearStringFromPunctuation,
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
}
