package director

import (
	"context"
	"encoding/json"
	"github.com/ThCompiler/go_game_constractor/director/scene"
)

// Result - of play scene
type Result struct {
	Text          scene.Text
	Buttons       []scene.Button
	IsEndOfScript bool
}

// SceneDirectorConfig - config for director
type SceneDirectorConfig struct {
	StartScene   scene.Scene
	GoodbyeScene scene.Scene
	EndCommand   string
}

// UserInfo - info about user for scene
type UserInfo struct {
	UserId    string
	SessionId string
}

// SceneRequest - request from user for scene
type SceneRequest struct {
	Command      string
	FullUserText string
	WasButton    bool
	Payload      json.RawMessage
	Info         UserInfo
}

func (sr *SceneRequest) toSceneContext(ctx context.Context) *scene.Context {
	return scene.NewContext(
		scene.Request{
			SearchedMessage: sr.Command,
			FullMessage:     sr.FullUserText,
			WasButton:       sr.WasButton,
			Payload:         sr.Payload,
		},
		scene.UserInfo{
			UserId:    sr.Info.UserId,
			SessionId: sr.Info.SessionId,
		},
		ctx,
	)
}
