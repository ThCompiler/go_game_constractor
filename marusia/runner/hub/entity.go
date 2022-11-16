package hub

import (
    drt "github.com/ThCompiler/go_game_constractor/director"
    "github.com/ThCompiler/go_game_constractor/marusia/runner"
)

func ToDirectorRequest(req runner.Request) drt.SceneRequest {
    return drt.SceneRequest{
        Request: drt.UserInput{
            Command:      req.Request.Command,
            FullUserText: req.Request.OriginalUtterance,
            WasButton:    req.Request.IsButton,
            Payload:      req.Request.Payload,
            NLU:          req.Request.NLU,
        },
        Info: drt.SessionInfo{
            UserId:       req.Session.User.UserId,
            SessionId:    req.Session.SessionID,
            IsNewSession: req.Session.New,
            UserVKId:     req.Session.User.UserVKId,
        },
        Application: drt.Application{
            ApplicationID:   req.Session.Application.ApplicationID,
            ApplicationType: req.Session.Application.ApplicationType,
        },
        GlobalContext: req.Context,
    }
}

func ToRunnerResult(result drt.Result) runner.Result {
    return runner.Result{
        Text: runner.Text{
            BaseText:     result.Text.BaseText,
            TextToSpeech: result.Text.TextToSpeech,
        },
        Buttons:       ToRunnerButtons(result.Buttons),
        IsEndOfScript: result.IsEndOfScript,
    }
}

func ToRunnerButton(button drt.Button) runner.Button {
    return runner.Button{
        Title:   button.Title,
        URL:     button.URL,
        Payload: button.Payload,
    }
}

func ToRunnerButtons(buttons []drt.Button) []runner.Button {
    res := make([]runner.Button, 0)
    for _, button := range buttons {
        res = append(res, ToRunnerButton(button))
    }
    return res
}
