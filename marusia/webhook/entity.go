package webhook

import (
	"github.com/ThCompiler/go_game_constractor/marusia"
	"github.com/ThCompiler/go_game_constractor/marusia/runner"
)

func ToHubRequest(req marusia.Request) runner.Request {
	return runner.Request{
		Meta:    ToHubMeta(req.Meta),
		Request: ToHubRequestIn(req.Request),
		Session: ToHubSession(req.Session),
	}
}

func ToHubRequestIn(req marusia.RequestIn) runner.RequestIn {
	return runner.RequestIn{
		Command:           req.Command,
		OriginalUtterance: req.OriginalUtterance,
		IsButton:          req.Type == marusia.ButtonPressed,
		Payload:           req.Payload,
		NLU:               req.NLU,
	}
}

func ToHubMeta(req marusia.Meta) runner.Meta {
	return runner.Meta{
		ClientID: req.ClientID,
		CityRu:   req.CityRu,
		Timezone: req.Timezone,
		Locale:   req.Locale,
	}
}

func ToHubSession(req marusia.Session) runner.Session {
	return runner.Session{
		SessionID:   req.SessionID,
		SkillID:     req.SkillID,
		New:         req.New,
		MessageID:   req.MessageID,
		Application: ToHubApplication(req.Application),
		User:        ToHubUser(req.User),
	}
}

func ToHubApplication(req marusia.Application) runner.Application {
	return runner.Application{
		ApplicationID:   req.ApplicationID,
		ApplicationType: req.ApplicationType,
	}
}

func ToHubUser(req marusia.User) runner.User {
	return runner.User{
		UserID:   req.UserID,
		UserVKID: req.UserVKID,
	}
}
