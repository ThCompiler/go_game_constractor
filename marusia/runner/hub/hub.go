package hub

import (
	drt "github.com/ThCompiler/go_game_constractor/director"
	"github.com/ThCompiler/go_game_constractor/marusia/runner"
)

type sceneMessage struct {
	sessionId string
	answer    chan runner.PlayedSceneResult
	req       runner.Request
}

type director struct {
	hub       *ScriptHub
	op        drt.Director
	sessionId string
}

func newDirector(hub *ScriptHub, sessionId string, op drt.Director) *director {
	return &director{
		hub:       hub,
		sessionId: sessionId,
		op:        op,
	}
}

func (c *director) Close() {
	c.hub.detachDirector(c)
}

func (c *director) PlayScene(msg sceneMessage) drt.Result {
	res := c.op.PlayScene(ToDirectorRequest(msg.req))

	return res
}

type ScriptHub struct {
	directors map[string]*director
	broadcast chan *sceneMessage
	attacher  chan *director
	dettacher chan *director
	stopHub   chan bool
}

func NewHub() *ScriptHub {
	return &ScriptHub{
		broadcast: make(chan *sceneMessage),
		attacher:  make(chan *director),
		dettacher: make(chan *director),
		directors: make(map[string]*director),
		stopHub:   make(chan bool),
	}
}

func (h *ScriptHub) AttachDirector(sessionId string, op drt.Director) {
	h.attacher <- newDirector(h, sessionId, op)
}

func (h *ScriptHub) detachDirector(drt *director) {
	h.dettacher <- drt
}

func (h *ScriptHub) RunScene(req runner.Request) chan runner.PlayedSceneResult {
	answer := make(chan runner.PlayedSceneResult)
	h.broadcast <- &(sceneMessage{
		sessionId: req.Session.SessionID,
		req:       req,
		answer:    answer,
	})
	return answer
}

func (h *ScriptHub) StopHub() {
	h.stopHub <- true
}

func (h *ScriptHub) detachAll() {
	for key := range h.directors {
		delete(h.directors, key)
	}
}

func (h *ScriptHub) runScene(msg *sceneMessage) {
	if drt, ok := h.directors[msg.sessionId]; ok {
		go func(ans chan runner.PlayedSceneResult, drt *director) {
			ans <- runner.PlayedSceneResult{
				Result:         ToRunnerResult(drt.PlayScene(*msg)),
				WorkedDirector: drt,
			}
		}(msg.answer, drt)
	}
}

func (h *ScriptHub) applyDirectorDetaching(drt *director) {
	delete(h.directors, drt.sessionId)
}

func (h *ScriptHub) Run() {
	for {
		select {
		case drt, ok := <-h.attacher:
			if ok {
				h.directors[drt.sessionId] = drt
			}
		case drt, ok := <-h.dettacher:
			if ok {
				h.applyDirectorDetaching(drt)
			}
		case msg, ok := <-h.broadcast:
			if ok {
				h.runScene(msg)
			}
		case <-h.stopHub:
			h.detachAll()
			return
		}
	}
}
