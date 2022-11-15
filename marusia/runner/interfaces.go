package runner

import (
	drt "github.com/ThCompiler/go_game_constractor/director"
)

type ClosedDirector interface {
	Close()
}

type PlayedSceneResult struct {
	Result
	WorkedDirector ClosedDirector
}

type ScriptRunner interface {
	AttachDirector(sessionId string, op drt.Director)
	RunScene(req Request) chan PlayedSceneResult
}
