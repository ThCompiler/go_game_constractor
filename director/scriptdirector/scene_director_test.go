package scriptdirector

/*
import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/ThCompiler/go_game_constractor/director"
	"github.com/ThCompiler/go_game_constractor/director/scriptdirector/scene"
	mock_scene "github.com/ThCompiler/go_game_constractor/director/scriptdirector/scene/mocks"
	ts "github.com/ThCompiler/go_game_constractor/pkg/testing"
)

type SceneDirectorSuite struct {
	ts.TestCasesSuite
	ActFunc func(args ...interface{}) []interface{}
}

func (s *SceneDirectorSuite) TestStrWithPunctuation() {
	s.ActFunc = func(args ...interface{}) []interface{} {
		scriptDirector := ScriptDirector{
			currentScene:  args[0].(scene.Scene),
			isEndOfScript: args[0].(scene.Scene),
			cf:            args[0].(scene.Scene),
		}

		res := scriptDirector.PlayScene(args[0].(director.SceneRequest))
		return []interface{}{res}
	}

	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name:     "All punctuation characters",
			Args:     ts.TTA(`!\"#$%&'()*+,-./:;<=>?@[\\]^_` + "`" + `{|}~`),
			Expected: ts.TTVE(""),
			InitMocks: func(ctrl *gomock.Controller) []interface{} {
				return []interface{}{
					mock_scene.NewMockScene(ctrl),
					mock_scene.NewMockError(ctrl),
					mock_scene.NewMockMessageMatcher(ctrl),
				}
			},
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
		s.ActFunc,
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
*/
