package webhook

import (
	game "github.com/ThCompiler/go_game_constractor/director"
	"github.com/ThCompiler/go_game_constractor/director/scene"
	"github.com/ThCompiler/go_game_constractor/marusia"
	"github.com/ThCompiler/go_game_constractor/marusia/hub"
	"github.com/ThCompiler/go_game_constractor/pkg/logger"
	"time"
)

const RequestTime = 60 * time.Second

func NewDefaultMarusiaWebhook(l logger.Interface, runnerHub hub.ScriptRunner, sdc game.SceneDirectorConfig) *marusia.Webhook {
	wh := marusia.NewWebhook(l)

	wh.OnEvent(func(r marusia.Request) (resp marusia.Response, err error) {
		err = nil

		if r.Request.Command == marusia.OnStart || r.Request.Command == "debug" {
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
				err = ErrorBadDirectorAnswer
			}
			break
		case <-ticker.C:
			err = ErrorTooLongRunning
			break
		}
		ticker.Stop()

		if err != nil {
			l.Error(err)
		}
		return
	})

	return wh
}

func toMarusiaButtons(buttons []scene.Button) []marusia.Button {
	res := make([]marusia.Button, 0)
	for _, button := range buttons {
		res = append(res, marusia.Button{
			Title:   button.Title,
			URL:     button.URL,
			Payload: button.Payload,
		})
	}
	return res
}
