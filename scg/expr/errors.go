package expr

import (
    "github.com/pkg/errors"
    "strings"
)

var (
    ErrorStartSceneNotFound   = errors.New("start scene with it name not found")
    ErrorGoodbyeSceneNotFound = errors.New("goodbye scene with it name not found")
    ErrorNameAlreadyOccupied  = errors.New("the name of matcher is already occupied by the standard matcher")
    ErrorNameSceneNotFound    = errors.New("the name of scene not found")
    ErrorUnknown              = errors.New("got unknown error. Please send the error information " +
        "and your configuration file to the mail: vetan22@mail.ru")
    ErrorNotFoundSLoadedContext = errors.New("not found value that you try load from context in higher-level scenes")
)

func errorNameSceneNotFound(sceneName string) error {
    return errors.Wrap(ErrorNameSceneNotFound, "with name: \""+sceneName+"\"")
}

func errorNotFoundLoadingContext(valueName string, sceneName string, visitedScenes []string) error {
    return errors.Wrap(ErrorNotFoundSLoadedContext, "with context value name: \""+valueName+"\", in the scene \""+sceneName+
        "\"not found as saved in higher-level scenes: "+strings.Join(visitedScenes, ", "))
}

func errorNotFoundLoadingContextInValues(valueName string, sceneName string, visitedScenes []string) error {
    return errors.Wrap(ErrorNotFoundSLoadedContext, "with value name: \""+valueName+"\", in text of the scene \""+sceneName+
        "\"not found as saved in higher-level scenes: "+strings.Join(visitedScenes, ", "))
}
