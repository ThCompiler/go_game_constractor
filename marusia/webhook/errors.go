package webhook

import "errors"

var (
	ErrorTooLongRunning    = errors.New("too long run scene")
	ErrorBadDirectorAnswer = errors.New("bad answer from script director")
)
