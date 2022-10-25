package marusia

import (
	drt "gameconstractor/director"
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
	RunScene(req Request) chan PlayedSceneResult
}
