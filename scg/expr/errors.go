package expr

import (
	"github.com/pkg/errors"
)

var (
	StartSceneNotFoundError = errors.New("start scene with it name not found")
)

func errorNameSceneExists(sceneName string) error {
	return errors.New("this name scene already exist \"" + sceneName + "\"")
}
