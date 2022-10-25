package director

import (
	"context"
	"encoding/json"
	"github.com/thcompiler/go_game_constractor/director/scene"
)

type Result struct {
	Text          scene.Text
	Buttons       []scene.Button
	IsEndOfScript bool
}

type SceneDirectorConfig struct {
	StartScene   scene.Scene
	GoodbyeScene scene.Scene
	EndCommand   string
}

type UserInfo struct {
	UserId    string
	SessionId string
}

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
