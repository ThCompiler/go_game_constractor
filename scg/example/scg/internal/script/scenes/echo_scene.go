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
	base_matchers "github.com/ThCompiler/go_game_constractor/director/scriptdirector/matchers"
	"github.com/ThCompiler/go_game_constractor/director/scriptdirector/scene"
	loghttp "github.com/ThCompiler/go_game_constractor/pkg/logger/context"
	"github.com/ThCompiler/go_game_constractor/scg/example/scg/internal/texts/manager"
	"github.com/ThCompiler/go_game_constractor/scg/go/types"
)

// Echo scene
type Echo struct {
	loghttp.LogObject
	TextManager manager.TextManager
	NextScene   SceneName
}

// React function of actions after scene has been played
func (sc *Echo) React(ctx *scene.Context) scene.Command {
	ctx.Set("sayed", types.MustConvert[string](ctx.Request.SearchedMessage))

	// TODO Write the actions after Echo scene has been played
	switch {

	// Matcher select
	case ctx.Request.NameMatched == base_matchers.AnyMatchedString:

	default:
		sc.NextScene = EchoScene
	}

	return scene.NoCommand
}

// Next function returning next scene
func (sc *Echo) Next() scene.Scene {
	if sc.NextScene == EchoRepeatScene {
		return &EchoRepeat{
			TextManager: sc.TextManager,
		}
	}

	return &Echo{
		TextManager: sc.TextManager,
	}
}

// GetSceneInfo function returning info about scene
func (sc *Echo) GetSceneInfo(_ *scene.Context) (scene.Info, bool) {
	// TODO Write some actions for get data for texts

	text, _ := sc.TextManager.GetEchoText()
	return scene.Info{
		Text: text,
		ExpectedMessages: []scene.MessageMatcher{
			base_matchers.AnyMatcher,
		},
		Buttons: []director.Button{},
		Err:     base_matchers.NumberError,
	}, true
}
