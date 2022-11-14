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
	"github.com/ThCompiler/go_game_constractor/scg/example/scg/echogame/manager"
)

// Hello scene
type Hello struct {
	TextManager manager.TextManager
	NextScene   SceneName
}

// React function of actions after scene has been played
func (sc *Hello) React(_ *scene.Context) scene.Command {
	return scene.NoCommand
}

// Next function returning next scene
func (sc *Hello) Next() scene.Scene {
	return &Echo{
		TextManager: sc.TextManager,
	}
}

// GetSceneInfo function returning info about scene
func (sc *Hello) GetSceneInfo(_ *scene.Context) (scene.Info, bool) {
	var (
		number int64
	)

	// TODO Write some actions for get data for texts

	text, _ := sc.TextManager.GetHelloText(
		number,
	)
	return scene.Info{
		Text:             text,
		ExpectedMessages: []scene.MessageMatcher{},
		Buttons:          []director.Button{},
		Err:              scene.BaseSceneError{Scene: &Goodbye{TextManager: sc.TextManager}},
	}, false
}
