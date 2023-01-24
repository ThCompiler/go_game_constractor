package scene

import (
	"golang.org/x/exp/slices"

	"github.com/ThCompiler/go_game_constractor/scg/script/errors"
	"github.com/ThCompiler/go_game_constractor/scg/script/matchers"
)

type Scene struct {
	IsEnd       bool              `yaml:"isEnd,omitempty" json:"is_end,omitempty" xml:"isEnd,omitempty"`
	IsInfoScene bool              `yaml:"isInfoScene,omitempty" json:"is_info_scene,omitempty" xml:"isInfoScene,omitempty"`
	NextScene   string            `yaml:"nextScene,omitempty" json:"next_scene,omitempty" xml:"nextScene,omitempty"`
	Text        Text              `yaml:"text" json:"text" xml:"text"`
	Error       Error             `yaml:"error,omitempty" json:"error,omitempty" xml:"error,omitempty"`
	Context     Context           `yaml:"context" json:"context" xml:"context"`
	NextScenes  []string          `yaml:"nextScenes,omitempty" json:"next_scenes,omitempty" xml:"nextScenes,omitempty"`
	Matchers    []Matcher         `yaml:"matchers,omitempty" json:"matchers,omitempty" xml:"matchers,omitempty"`
	Buttons     map[string]Button `yaml:"buttons,omitempty" json:"buttons,omitempty" xml:"buttons,omitempty"`
}

func (s *Scene) IsValid(userMatchers map[string]ScriptMatcher) (bool, error) {
	if is, err := s.checkEmptyField(); !is {
		return is, err
	}

	if is, err := s.isMatchersValid(userMatchers); !is {
		return is, err
	}

	if is, err := s.isErrorsValid(); !is {
		return is, err
	}

	if is, err := s.isButtonValid(); !is {
		return is, err
	}

	if err := s.Context.checkValuesType(); err != nil {
		return false, err
	}

	return s.Text.IsValid()
}

func (s *Scene) checkEmptyField() (bool, error) {
	if s.IsInfoScene && len(s.NextScenes) != 0 && s.NextScene == "" {
		return false, ErrorEmptyNextSceneWithInfoScene
	}

	if !s.IsInfoScene && len(s.NextScenes) == 0 && s.NextScene != "" {
		return false, ErrorEmptyNextScenesWithNoInfoScene
	}

	return true, nil
}

func (s *Scene) isMatchersValid(userMatchers map[string]ScriptMatcher) (bool, error) {
	err := error(nil)

	for _, matcher := range s.Matchers {
		if matchers.IsCorrectNameOfMather(matcher.Name) {
			continue
		}

		if _, is := userMatchers[matcher.Name]; !is {
			err = errorNotSupportedMatherType(matcher.Name)

			break
		}

		if is := slices.Contains(s.NextScenes, matcher.ToScene); !is && matcher.ToScene != "" {
			err = errorNotFoundToSceneInMather(matcher.ToScene, matcher.Name)

			break
		}
	}

	return err == nil, err
}

func (s *Scene) isErrorsValid() (bool, error) {
	if s.Error.IsBase() {
		if !errors.IsCorrectNameOfError(s.Error.Base) {
			return false, errorNotSupportedErrorType(s.Error.Base)
		}
	}

	return true, nil
}

func (s *Scene) isButtonValid() (bool, error) {
	err := error(nil)

	for name, button := range s.Buttons {
		if is := slices.Contains(s.NextScenes, button.ToScene); !is && button.ToScene != "" {
			err = errorNotFoundToSceneInButton(button.ToScene, name)

			break
		}
	}

	return err == nil, err
}

type GoodByeScene struct {
	Scene
	Name string `yaml:"name" json:"name" xml:"name"`
}
