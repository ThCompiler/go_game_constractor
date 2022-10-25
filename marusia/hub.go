package marusia

import (
	"encoding/json"
	drt "gameconstractor/director"
)

type request struct {
	command      string
	fullUserText string
	payload      json.RawMessage
	wasButton    bool
}

func fromMarusiaRequest(rqm RequestIn) *request {
	return &request{
		command:      rqm.Command,
		fullUserText: rqm.OriginalUtterance,
		payload:      rqm.Payload,
		wasButton:    rqm.Type == ButtonPressed,
	}
}

type sceneMessage struct {
	userId    string
	sessionId string
	answer    chan PlayedSceneResult
	rq        request
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
	res := c.op.PlayScene(drt.SceneRequest{
		Command:      msg.rq.command,
		FullUserText: msg.rq.fullUserText,
		WasButton:    msg.rq.wasButton,
		Payload:      msg.rq.payload,
		Info: drt.UserInfo{
			UserId:    msg.userId,
			SessionId: msg.sessionId,
		},
	})

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

func (h *ScriptHub) RunScene(rq Request) chan PlayedSceneResult {
	answer := make(chan PlayedSceneResult)
	h.broadcast <- &(sceneMessage{
		userId:    rq.Session.UserID,
		sessionId: rq.Session.SessionID,
		rq:        *fromMarusiaRequest(rq.Request),
		answer:    answer,
	})
	return answer
}

func (h *ScriptHub) StopHub() {
	h.stopHub <- true
}

func (h *ScriptHub) detachAll() {
	for key, _ := range h.directors {
		delete(h.directors, key)
	}
}

func (h *ScriptHub) runScene(msg *sceneMessage) {
	if drt, ok := h.directors[msg.sessionId]; ok {
		go func(ans chan PlayedSceneResult, drt *director) {
			ans <- PlayedSceneResult{
				Result:         drt.PlayScene(*msg),
				WorkedDirector: drt,
			}
		}(msg.answer, drt)
	}
}

func (h *ScriptHub) applyDirectorDetaching(drt *director) {
	if _, ok := h.directors[drt.sessionId]; ok {
		delete(h.directors, drt.sessionId)
	}
}

func (h *ScriptHub) Run() {
	for {
		select {
		case drt, ok := <-h.attacher:
			if ok {
				h.directors[drt.sessionId] = drt
			}
			break
		case drt, ok := <-h.dettacher:
			if ok {
				h.applyDirectorDetaching(drt)
			}
			break
		case msg, ok := <-h.broadcast:
			if ok {
				h.runScene(msg)
			}
			break
		case <-h.stopHub:
			h.detachAll()
			return
		default:
			break
		}
	}
}
