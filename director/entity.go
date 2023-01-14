package director

import (
	"context"
	"encoding/json"

	"github.com/ThCompiler/go_game_constractor/marusia"
	"github.com/ThCompiler/go_game_constractor/pkg/language"
)

type Button struct {
	Title   string
	URL     string
	Payload interface{}
}

type Text struct {
	BaseText     string
	TextToSpeech string
}

// Result - of play scene.
type Result struct {
	Text          Text
	Buttons       []Button
	IsEndOfScript bool
}

// SessionInfo - info about user for scene.
type SessionInfo struct {
	UserID    string
	SessionID string

	// VK ID user
	UserVKID string

	IsNewSession bool
}

// Application info about app.
type Application struct {
	// ID of app
	ApplicationID string

	// Types of app:
	//  • mobile;
	//  • speaker;
	//  • VK;
	//  • other.
	ApplicationType marusia.ApplicationType
}

type UserInput struct {
	Command      string
	FullUserText string
	WasButton    bool
	Payload      json.RawMessage
	NLU          language.NLU
}

// SceneRequest - request from user for scene.
type SceneRequest struct {
	Request       UserInput
	Info          SessionInfo
	Application   Application
	GlobalContext context.Context
}
