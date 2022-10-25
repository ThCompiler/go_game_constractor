package hub

import (
	drt "github.com/ThCompiler/go_game_constractor/director"
	"github.com/ThCompiler/go_game_constractor/marusia"
)

type ClosedDirector interface {
	Close()
}

type PlayedSceneResult struct {
	drt.Result
	WorkedDirector ClosedDirector
}

type ScriptRunner interface {
	AttachDirector(sessionId string, op drt.Director)
	RunScene(req marusia.Request) chan PlayedSceneResult
}
