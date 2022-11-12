package marusia

import (
	game "github.com/ThCompiler/go_game_constractor/director"
	"github.com/ThCompiler/go_game_constractor/director/scene"
	"github.com/ThCompiler/go_game_constractor/marusia/hub"
	"github.com/ThCompiler/go_game_constractor/pkg/logger"
	"time"
)

const RequestTime = 60 * time.Second

func NewDefaultMarusiaWebhook(l logger.Interface, runnerHub hub.ScriptRunner, sdc game.SceneDirectorConfig) *Webhook {
	wh := &Webhook{
		l: l,
	}
	wh.OnEvent(func(r Request) (resp Response, err error) {
		err = nil

		if r.Request.Command == OnStart || r.Request.Command == "debug" {
			runnerHub.AttachDirector(r.Session.SessionID, game.NewScriptDirector(sdc))
		}

		ans := runnerHub.RunScene(r)

		ticker := time.NewTicker(RequestTime)
		select {
		case answer, ok := <-ans:
			if ok {
				resp.Text = answer.Text.BaseText
				resp.TTS = answer.Text.TextToSpeech
				resp.EndSession = answer.IsEndOfScript
				resp.Buttons = toMarusiaButtons(answer.Buttons)

				if answer.IsEndOfScript {
					answer.WorkedDirector.Close()
				}
			} else {
				err = BadDirectorAnswer
			}
			break
		case <-ticker.C:
			err = TooLongRunning
			break
		}
		ticker.Stop()

		if err != nil {
			wh.l.Error(err)
		}
		return
	})

	return wh
}

func toMarusiaButtons(buttons []scene.Button) []Button {
	res := make([]Button, 0)
	for _, button := range buttons {
		res = append(res, Button{
			Title:   button.Title,
			URL:     button.URL,
			Payload: button.Payload,
		})
	}
	return res
}
