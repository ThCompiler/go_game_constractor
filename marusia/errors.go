package marusia

import "errors"

var (
	TooLongRunning    = errors.New("too long run scene")
	BadDirectorAnswer = errors.New("bad answer from script director")
)
