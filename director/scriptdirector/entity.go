package scriptdirector

import (
	"context"
	"github.com/ThCompiler/go_game_constractor/director"
	"github.com/ThCompiler/go_game_constractor/director/scriptdirector/scene"
)

func toSceneContext(sr director.SceneRequest, ctx context.Context) *scene.Context {
	return scene.NewContext(
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
			UserId:    sr.Info.UserId,
			SessionId: sr.Info.SessionId,
			UserVKId:  sr.Info.UserVKId,
		},
		ctx,
	)
}
