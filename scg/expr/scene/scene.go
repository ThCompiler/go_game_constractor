package scene

import (
	"github.com/ThCompiler/go_game_constractor/scg/script/errors"
	"github.com/ThCompiler/go_game_constractor/scg/script/matchers"
)

type Scene struct {
	Text        Text              `yaml:"text" json:"text" xml:"text"`
	NextScene   string            `yaml:"nextScene,omitempty" json:"next_scene,omitempty" xml:"nextScene,omitempty"`
	NextScenes  []string          `yaml:"nextScenes,omitempty" json:"next_scenes.md,omitempty" xml:"nextScenes,omitempty"`
	IsInfoScene bool              `yaml:"isInfoScene,omitempty" json:"is_info_scene,omitempty" xml:"isInfoScene,omitempty"`
	Matchers    []string          `yaml:"matchers,omitempty" json:"matchers,omitempty" xml:"matchers,omitempty"`
	Error       Error             `yaml:"error,omitempty" json:"error,omitempty" xml:"error,omitempty"`
	Buttons     map[string]Button `yaml:"buttons,omitempty" json:"buttons,omitempty" xml:"buttons,omitempty"`
}

func (s *Scene) IsValid(userMatchers map[string]Matcher) (bool, error) {
	if s.IsInfoScene && len(s.NextScenes) != 0 && s.NextScene == "" {
		return false, errorEmptyNextSceneWithInfoScene
	}

	if !s.IsInfoScene && len(s.NextScenes) == 0 && s.NextScene != "" {
		return false, errorEmptyNextScenesWithNoInfoScene
	}

	is, err := s.isMatchersValid(userMatchers)

	if !is {
		return is, err
	}

	is, err = s.isErrorsValid()

	if !is {
		return is, err
	}

	return s.Text.IsValid()
}

func (s *Scene) isMatchersValid(userMatchers map[string]Matcher) (bool, error) {
	err := error(nil)
	for _, matcher := range s.Matchers {
		if matchers.IsCorrectNameOfMather(matcher) {
			continue
		}

		if _, is := userMatchers[matcher]; !is {
			err = errorNotSupportedMatherType(matcher)
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

type GoodByeScene struct {
	Scene
	Name string `yaml:"name" json:"name" xml:"name"`
}
