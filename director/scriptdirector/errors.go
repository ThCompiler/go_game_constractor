package scriptdirector

import "github.com/pkg/errors"

var (
	ErrorNotSetSettingsScene = errors.New("there is not set start or goodbye scene in director config")
	ErrorRunFinishedDirector = errors.New("attempt to process a new event in the director that " +
		"is in the finish state, which is impossible")
)
