package expr

import (
	"github.com/pkg/errors"
)

func errorNameSceneExists(sceneName string) error {
	return errors.New("this name scene already exist \"" + sceneName + "\"")
}
