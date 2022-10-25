package director

import (
	"context"
	"github.com/thcompiler/go_game_constractor/director/scene"
	"github.com/thcompiler/go_game_constractor/marusia"
	"github.com/thcompiler/go_game_constractor/pkg/stack"
	"strings"
)

type ScriptDirector struct {
	stashedScene  stack.Stack[scene.Scene]
	currentScene  scene.Scene
	isEndOfScript bool
	ctx           context.Context
	cf            SceneDirectorConfig
}

func NewScriptDirector(cf SceneDirectorConfig) *ScriptDirector {
	return &ScriptDirector{
		stashedScene:  stack.NewStack[scene.Scene](),
		cf:            cf,
		currentScene:  nil,
		ctx:           context.Background(),
		isEndOfScript: false,
	}
}

func (so *ScriptDirector) PlayScene(req SceneRequest) Result {
	ctx := req.toSceneContext(so.ctx)

	sceneInfo := scene.Info{}
	if so.currentScene != nil {
		sceneInfo, _ = so.currentScene.GetSceneInfo(ctx)
	}

	Err := scene.Error(nil)

	switch strings.ToLower(req.Command) {
	case marusia.OnStart, "debug":
		so.currentScene = so.cf.StartScene

	case marusia.OnInterrupt, strings.ToLower(so.cf.EndCommand):
		so.stashedScene.Push(so.currentScene)
		sceneCmd := so.cf.GoodbyeScene.React(ctx)
		so.currentScene = so.cf.GoodbyeScene.Next()
		so.reactSceneCommand(sceneCmd)

	default:
		cmd := so.matchCommands(req.Command, sceneInfo.ExpectedMessages)
		if cmd != "" {
			ctx.Request.SearchedMessage = cmd
			sceneCmd := so.currentScene.React(ctx)
			so.currentScene = so.currentScene.Next()
			so.reactSceneCommand(sceneCmd)
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
		info.Buttons = append(info.Buttons, scene.Button{Title: so.cf.EndCommand})
	}

	so.ctx = ctx.Context
	return Result{
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
			Text: scene.Text{
				BaseText:     oldInfo.Text.BaseText + "\n" + info.Text.BaseText,
				TextToSpeech: oldInfo.Text.TextToSpeech + "\n" + info.Text.TextToSpeech,
			},
			Buttons:          info.Buttons,
			ExpectedMessages: info.ExpectedMessages,
		}
	}

	return currentScene, info
}

func (so *ScriptDirector) reactSceneCommand(command scene.Command) {
	switch command {
	case scene.ApplyStashedScene:
		if !so.stashedScene.Empty() {
			so.currentScene, _ = so.stashedScene.Pop()
		}
		break
	case scene.FinishScene:
		so.isEndOfScript = true
		break
	}
}

func (so *ScriptDirector) matchCommands(command string, commands []scene.MessageMatcher) string {
	for _, cmd := range commands {
		if matched, msg := cmd.Match(command); matched {
			return msg
		}
	}
	return ""
}
