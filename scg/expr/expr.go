package expr

import (
	"fmt"
	"github.com/ThCompiler/go_game_constractor/pkg/cleanenv"
	"github.com/ThCompiler/go_game_constractor/scg/expr/scene"
)

type Script map[string]scene.Scene

type ScriptInfo struct {
	StartScene     string             `yaml:"startScene" json:"start_scene" xml:"startScene"`
	Name           string             `yaml:"name" json:"name" xml:"name"`
	GoodByeCommand string             `yaml:"goodByeCommand" json:"good_bye_command" xml:"goodByeCommand"`
	GoodByeScene   scene.GoodByeScene `yaml:"goodByeScene" json:"good_bye_scene" xml:"goodByeScene"`
	Script         Script             `yaml:"script" json:"script" xml:"script"`
}

func (si *ScriptInfo) IsValid() (bool, error) {
	is, err := si.GoodByeScene.IsValid()
	if !is {
		return is, err
	}

	for _, sc := range si.Script {
		if is, err = sc.IsValid(); !is {
			break
		}
	}
	if !is {
		return is, err
	}

	if _, is = si.Script[si.GoodByeScene.Name]; is {
		return false, errorNameSceneExists(si.GoodByeScene.Name)
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
