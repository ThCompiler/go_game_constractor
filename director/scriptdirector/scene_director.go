package scriptdirector

import (
	"context"
	"github.com/ThCompiler/go_game_constractor/director"
	"github.com/ThCompiler/go_game_constractor/director/scriptdirector/scene"
	"github.com/ThCompiler/go_game_constractor/marusia"
	"github.com/ThCompiler/go_game_constractor/pkg/stack"
	"github.com/ThCompiler/go_game_constractor/pkg/stringutilits"
	"strings"
)

// SceneDirectorConfig - config for director
type SceneDirectorConfig struct {
	StartScene   scene.Scene
	GoodbyeScene scene.Scene
	EndCommand   string
}

// ScriptDirector - implementation game director for script games
type ScriptDirector struct {
	stashedScene  stack.Stack[scene.Scene]
	currentScene  scene.Scene
	isEndOfScript bool
	ctx           context.Context
	cf            SceneDirectorConfig
}

// NewScriptDirector - create new ScriptDirector
func NewScriptDirector(cf SceneDirectorConfig) *ScriptDirector {
	return &ScriptDirector{
		stashedScene:  stack.NewStack[scene.Scene](),
		cf:            cf,
		currentScene:  nil,
		ctx:           context.Background(),
		isEndOfScript: false,
	}
}

// PlayScene - .
func (so *ScriptDirector) PlayScene(req director.SceneRequest) director.Result {
	ctx := toSceneContext(req, so.ctx)

	sceneInfo := scene.Info{}
	if so.currentScene != nil {
		sceneInfo, _ = so.currentScene.GetSceneInfo(ctx)
	}

	Err := scene.Error(nil)

	switch {
	case strings.EqualFold(marusia.OnStart, req.Request.Command) || req.Info.IsNewSession:
		so.currentScene = so.cf.StartScene

	case strings.EqualFold(marusia.OnInterrupt, req.Request.Command),
		strings.EqualFold(stringutilits.ClearStringFromPunctuation(so.cf.EndCommand), req.Request.Command):
		so.stashedScene.Push(so.currentScene)
		sceneCmd := so.cf.GoodbyeScene.React(ctx)
		so.reactSceneCommand(sceneCmd, so.cf.GoodbyeScene.Next())

	default:
		var cmd, name string
		if req.Request.WasButton {
			cmd, name = so.matchButton(req.Request.Command, sceneInfo.Buttons)
		} else {
			cmd, name = so.matchCommands(req.Request.Command, sceneInfo.ExpectedMessages)
		}
		if cmd != "" {
			ctx.Request.SearchedMessage = cmd
			ctx.Request.NameMatched = name
			sceneCmd := so.currentScene.React(ctx)
			so.reactSceneCommand(sceneCmd, so.currentScene.Next())
		} else {
			Err = sceneInfo.Err
		}
	}

	info := scene.Info{}

	if Err != nil {
		if Err.IsErrorScene() {
			errCmd := Err.GetErrorScene().React(ctx)
			tmpScene, ErrInfo := so.baseSceneInfo(Err.GetErrorScene().Next(), ctx)
			if errCmd != scene.ApplyStashedScene {
				so.stashedScene.Push(so.currentScene)
				so.currentScene = tmpScene
			}

			info = ErrInfo
		} else {
			info = sceneInfo
			info.Text.BaseText = Err.GetErrorText()
			info.Text.TextToSpeech = Err.GetErrorText()
		}
	} else {
		so.currentScene, info = so.baseSceneInfo(so.currentScene, ctx)
	}

	if !so.isEndOfScript {
		info.Buttons = append(info.Buttons, director.Button{Title: so.cf.EndCommand})
	}

	so.ctx = ctx.Context
	return director.Result{
		Text:          info.Text,
		Buttons:       info.Buttons,
		IsEndOfScript: so.isEndOfScript,
	}
}

func (so *ScriptDirector) baseSceneInfo(currentScene scene.Scene, ctx *scene.Context) (scene.Scene, scene.Info) {
	info, withReact := currentScene.GetSceneInfo(ctx)
	for !withReact {
		currentScene = currentScene.Next()
		oldInfo := info
		info, withReact = currentScene.GetSceneInfo(ctx)
		info = scene.Info{
			Text: director.Text{
				BaseText:     oldInfo.Text.BaseText + "\n" + info.Text.BaseText,
				TextToSpeech: oldInfo.Text.TextToSpeech + "\n" + info.Text.TextToSpeech,
			},
			Buttons:          info.Buttons,
			ExpectedMessages: info.ExpectedMessages,
		}
	}

	return currentScene, info
}

func (so *ScriptDirector) reactSceneCommand(command scene.Command, nextScene scene.Scene) {
	switch command {
	case scene.StashScene:
		so.stashedScene.Push(so.currentScene)
		so.currentScene = nextScene
	case scene.ApplyStashedScene:
		if !so.stashedScene.Empty() {
			so.currentScene, _ = so.stashedScene.Pop()
		}
	case scene.FinishScene:
		so.isEndOfScript = true
		so.currentScene = nextScene
	default:
		so.currentScene = nextScene
	}
}

func (so *ScriptDirector) matchCommands(command string, commands []scene.MessageMatcher) (string, string) {
	for _, cmd := range commands {
		if matched, msg := cmd.Match(command); matched {
			return msg, cmd.GetMatchedName()
		}
	}
	return "", ""
}

func (so *ScriptDirector) matchButton(command string, buttons []director.Button) (string, string) {
	for _, button := range buttons {
		if strings.EqualFold(command, stringutilits.ClearStringFromPunctuation(button.Title)) {
			return command, button.Title
		}
	}
	return "", ""
}
