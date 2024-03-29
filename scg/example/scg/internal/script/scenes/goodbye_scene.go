// Code generated by scg 1, .
//
// EchoGame-SceneStructs
//
// Command:
// scg
//.

package scenes

import (
	"github.com/ThCompiler/go_game_constractor/director"
	"github.com/ThCompiler/go_game_constractor/director/scriptdirector/scene"
	"github.com/ThCompiler/go_game_constractor/pkg/logger/context"
	"github.com/ThCompiler/go_game_constractor/scg/example/scg/internal/texts/manager"
)

// Goodbye scene
type Goodbye struct {
	context.LogObject
	TextManager manager.TextManager
	NextScene   SceneName
}

// React function of actions after scene has been played
func (sc *Goodbye) React(_ *scene.Context) scene.Command {
	// TODO Write the actions after Goodbye scene has been played

	sc.NextScene = GoodbyeScene // TODO: manually set next scene after reaction

	return scene.FinishScene
}

// Next function returning next scene
func (sc *Goodbye) Next() scene.Scene {
	if sc.NextScene == GoodbyeScene {
		return &Goodbye{
			TextManager: sc.TextManager,
		}
	}

	return &Goodbye{
		TextManager: sc.TextManager,
	}
}

// GetSceneInfo function returning info about scene
func (sc *Goodbye) GetSceneInfo(ctx *scene.Context) (scene.Info, bool) {
	// TODO Write some actions for get data for texts

	text, err := sc.TextManager.GetGoodbyeText()
	if err != nil {
		sc.Log(ctx).Error(err, "error while getting text for Goodbye scene")

		return scene.Info{}, false
	}

	return scene.Info{
		Text:             text,
		ExpectedMessages: []scene.MessageMatcher{},
		Buttons:          []director.Button{},
	}, true
}
