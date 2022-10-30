package scene

import (
	"github.com/ThCompiler/go_game_constractor/scg/script/errors"
	"github.com/ThCompiler/go_game_constractor/scg/script/matchers"
)

type Scene struct {
	Text        Text              `yaml:"text" json:"text" xml:"text"`
	NextScenes  []string          `yaml:"nextScenes,omitempty" json:"next_scenes,omitempty" xml:"nextScenes,omitempty"`
	IsInfoScene bool              `yaml:"isInfoScene,omitempty" json:"is_info_scene,omitempty" xml:"isInfoScene,omitempty"`
	Matchers    []Matcher         `yaml:"matchers,omitempty" json:"matchers,omitempty" xml:"matchers,omitempty"`
	Error       Error             `yaml:"error,omitempty" json:"error,omitempty" xml:"error,omitempty"`
	Buttons     map[string]Button `yaml:"buttons,omitempty" json:"buttons,omitempty" xml:"buttons,omitempty"`
}

func (s *Scene) IsValid() (bool, error) {

	is, err := s.isMatchersValid()

	if !is {
		return is, err
	}

	is, err = s.isErrorsValid()

	if !is {
		return is, err
	}

	return s.Text.IsValid()
}

func (s *Scene) isMatchersValid() (bool, error) {
	err := error(nil)
	for _, matcher := range s.Matchers {
		if matcher.IsDefaultMatcher() {
			name, _ := matcher.GetStandardMatcher()
			if !matchers.IsCorrectNameOfMather(name) {
				err = errorNotSupportedMatherType(name)
				break
			}
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
