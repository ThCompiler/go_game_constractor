package expr

import (
	"github.com/pkg/errors"
)

var (
	StartSceneNotFoundError   = errors.New("start scene with it name not found")
	GoodbyeSceneNotFoundError = errors.New("goodbye scene with it name not found")
)

func errorNameSceneNotFound(sceneName string) error {
	return errors.New("this name scene not found: \"" + sceneName + "\"")
}
