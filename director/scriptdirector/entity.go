package scriptdirector

import (
	"context"
	"github.com/ThCompiler/go_game_constractor/director"
	"github.com/ThCompiler/go_game_constractor/director/scriptdirector/scene"
)

func toSceneContext(ctx context.Context, sr director.SceneRequest) *scene.Context {
	return scene.NewContext(
		ctx,
		sr.GlobalContext,
		scene.Request{
			SearchedMessage: sr.Request.Command,
			FullMessage:     sr.Request.FullUserText,
			WasButton:       sr.Request.WasButton,
			Payload:         sr.Request.Payload,
			ApplicationType: sr.Application.ApplicationType,
			NLU: scene.NLU{
				Tokens:   sr.Request.NLU.Tokens,
				Entities: sr.Request.NLU.Entities,
			},
		},
		scene.UserInfo{
			UserID:    sr.Info.UserID,
			SessionID: sr.Info.SessionID,
			UserVKID:  sr.Info.UserVKID,
		},
	)
}
