package expr

import (
	"fmt"
	"github.com/ThCompiler/go_game_constractor/pkg/cleanenv"
	"github.com/ThCompiler/go_game_constractor/scg/expr/scene"
)

type Script map[string]scene.Scene

type ScriptInfo struct {
	StartScene     string                   `yaml:"startScene" json:"start_scene" xml:"startScene"`
	Name           string                   `yaml:"name" json:"name" xml:"name"`
	GoodByeCommand string                   `yaml:"goodByeCommand" json:"good_bye_command" xml:"goodByeCommand"`
	GoodByeScene   string                   `yaml:"goodByeScene" json:"good_bye_scene" xml:"goodByeScene"`
	UserMatchers   map[string]scene.Matcher `yaml:"matchers,omitempty" json:"matchers,omitempty" xml:"matchers,omitempty"`
	Script         Script                   `yaml:"script" json:"script" xml:"script"`
}

func (si *ScriptInfo) IsValid() (is bool, err error) {
	unknownScene := ""
up:
	for _, sc := range si.Script {
		if is, err = sc.IsValid(si.UserMatchers); !is {
			break
		}

		if _, is = si.Script[sc.NextScene]; sc.IsInfoScene && !is {
			unknownScene = sc.NextScene
			break
		}

		for _, name := range sc.NextScenes {
			if _, is = si.Script[name]; !is {
				unknownScene = name
				break up
			}
		}
	}
	if !is {
		if unknownScene == "" {
			return is, err
		}
		return is, errorNameSceneNotFound(unknownScene)
	}

	if _, is = si.Script[si.GoodByeScene]; !is {
		return false, StartSceneNotFoundError
	}

	if _, is = si.Script[si.StartScene]; !is {
		return false, StartSceneNotFoundError
	}

	return false, nil
}

func LoadScriptInfo(path string) (*ScriptInfo, error) {
	si := ScriptInfo{}
	err := cleanenv.ReadConfig(path, &si)
	if err != nil {
		return nil, fmt.Errorf("error load script info: %w", err)
	}

	_, err = si.IsValid()
	if err != nil {
		return nil, fmt.Errorf("this script config is not correct: %w", err)
	}
	return &si, nil
}
