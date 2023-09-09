package scriptdirector

import (
	"context"
	"strings"

	"github.com/ThCompiler/go_game_constractor/director"
	"github.com/ThCompiler/go_game_constractor/director/scriptdirector/scene"
	"github.com/ThCompiler/go_game_constractor/marusia"
	context2 "github.com/ThCompiler/go_game_constractor/pkg/logger/context"
	"github.com/ThCompiler/go_game_constractor/pkg/stringutilits"
	"github.com/ThCompiler/go_game_constractor/pkg/structures"
)

// SceneDirectorConfig - config for director.
type SceneDirectorConfig struct {
	StartScene   scene.Scene
	GoodbyeScene scene.Scene
	EndCommand   string
	//TODO System Error scene and handle system error
}

// ScriptDirector - implementation game director for script games.
type ScriptDirector struct {
	context2.LogObject
	stashedScene  structures.Stack[scene.Scene]
	currentScene  scene.Scene
	isEndOfScript bool
	ctx           context.Context
	cf            SceneDirectorConfig
}

// NewScriptDirector - create new ScriptDirector.
func NewScriptDirector(cf SceneDirectorConfig) (*ScriptDirector, error) {
	if cf.StartScene == nil || cf.GoodbyeScene == nil {
		return nil, ErrorNotSetSettingsScene
	}

	return &ScriptDirector{
		stashedScene:  structures.NewStack[scene.Scene](),
		cf:            cf,
		currentScene:  nil,
		ctx:           context.Background(),
		isEndOfScript: false,
	}, nil
}

// PlayScene - .
func (so *ScriptDirector) PlayScene(req director.SceneRequest) director.Result {
	if so.isEndOfScript {
		panic(ErrorRunFinishedDirector)
	}

	ctx := toSceneContext(so.ctx, req)

	sceneInfo := scene.Info{}
	if so.currentScene != nil {
		sceneInfo, _ = so.currentScene.GetSceneInfo(ctx)
	}

	errScene := scene.Error(nil)

	switch {
	case strings.EqualFold(marusia.OnStart, req.Request.Command) || req.Info.IsNewSession:
		so.currentScene = so.cf.StartScene

	case strings.EqualFold(marusia.OnInterrupt, req.Request.Command) ||
		strings.EqualFold(stringutilits.ClearStringFromPunctuation(so.cf.EndCommand), req.Request.Command):
		so.stashedScene.Push(so.currentScene)
		sceneCmd := so.cf.GoodbyeScene.React(ctx)
		so.reactSceneCommand(sceneCmd, so.cf.GoodbyeScene.Next())

	default:
		errScene = so.operateBaseScenes(req, sceneInfo, ctx)
	}

	info := scene.Info{}

	if errScene != nil {
		so.errorAction(&info, errScene, ctx, sceneInfo)
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

func (so *ScriptDirector) operateBaseScenes(req director.SceneRequest, sceneInfo scene.Info,
	ctx *scene.Context,
) scene.Error {
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
		return sceneInfo.Err
	}

	return nil
}

func (so *ScriptDirector) errorAction(info *scene.Info, errScene scene.Error,
	ctx *scene.Context, sceneInfo scene.Info,
) {
	if errScene.IsErrorScene() {
		errCmd := errScene.GetErrorScene().React(ctx)
		tmpScene, ErrInfo := so.baseSceneInfo(errScene.GetErrorScene().Next(), ctx)

		if errCmd != scene.ApplyStashedScene {
			so.stashedScene.Push(so.currentScene)
			so.currentScene = tmpScene
		}

		*info = ErrInfo
	} else {
		*info = sceneInfo
		info.Text.BaseText = errScene.GetErrorText()
		info.Text.TextToSpeech = errScene.GetErrorText()
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
			so.currentScene = so.stashedScene.MustPop()
		}

	case scene.FinishScene:
		so.isEndOfScript = true
		so.currentScene = nextScene

	case scene.NoCommand:
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
