// Code generated by scg 1, .
//
// EchoGame-SceneStructs
//
// Command:
// scg
//.

package scenes

import (
	"github.com/ThCompiler/go_game_constractor/director/scene"
	"github.com/ThCompiler/go_game_constractor/scg/example/scg/echogame/manager"
)

type Goodbye struct {
	TextManager manager.TextManager
	NextScene   SceneName
}

func (sc *Goodbye) React(_ *scene.Context) scene.Command {
	// TODO

	sc.NextScene = GoodbyeScene // TODO: manually set next scene after reaction
	return scene.NoCommand
}

func (sc *Goodbye) Next() scene.Scene {
	switch sc.NextScene {
	case GoodbyeScene:
		return &Goodbye{
			TextManager: sc.TextManager,
		}
	}

	return &Goodbye{
		TextManager: sc.TextManager,
	}
}

func (sc *Goodbye) GetSceneInfo(ctx *scene.Context) (scene.Info, bool) {

	//TODO

	text, _ := sc.TextManager.GetGoodbyeText()
	return scene.Info{
		Text:             text,
		ExpectedMessages: []scene.MessageMatcher{},
		Buttons:          []scene.Button{},
	}, true
}