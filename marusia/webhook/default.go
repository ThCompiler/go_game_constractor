package webhook

import (
	"time"

	sd "github.com/ThCompiler/go_game_constractor/director/scriptdirector"
	"github.com/ThCompiler/go_game_constractor/marusia"
	"github.com/ThCompiler/go_game_constractor/marusia/runner"
	"github.com/ThCompiler/go_game_constractor/pkg/logger"
)

const RequestTime = 60 * time.Second

func NewDefaultMarusiaWebhook(l logger.Interface, runnerHub runner.ScriptRunner,
	sdc sd.SceneDirectorConfig,
) *marusia.Webhook {
	wh := marusia.NewWebhook(l)

	wh.OnEvent(func(r marusia.Request) (resp marusia.Response, err error) {
		err = nil

		if r.Request.Command == marusia.OnStart || r.Session.New {
			runnerHub.AttachDirector(r.Session.SessionID, sd.NewScriptDirector(sdc))
		}

		ans := runnerHub.RunScene(ToHubRequest(r))

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

func toMarusiaButtons(buttons []runner.Button) []marusia.Button {
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
