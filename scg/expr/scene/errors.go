package scene

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/ThCompiler/go_game_constractor/scg/go/types"
	errors2 "github.com/ThCompiler/go_game_constractor/scg/script/errors"
	"github.com/ThCompiler/go_game_constractor/scg/script/matchers"
)

var (
	ErrorIsNotRegexMatcher              = errors.New("this matcher is not regex, but you try get regex matcher")
	ErrorIsNotSelectsMatcher            = errors.New("this matcher is not selects, but you try get selects matcher")
	ErrorEmptyNextSceneWithInfoScene    = errors.New("if scene is info scene use only flag NextScene")
	ErrorEmptyNextScenesWithNoInfoScene = errors.New("if scene is not info scene use only flag NextScenes")
	ErrorTooManyFields                  = errors.New("too many fields were passed to the matcher")
	ErrorUnknownTypeOfValue             = errors.New("the type of values is not supported. Supported type is:" +
		types.GetSupportTypes())
	ErrorNotFoundValueInText    = errors.New("settled value is not found in the text")
	ErrorNotSupportedMatherType = errors.New("the matcher name is not supported. Supported matchers is: " +
		strings.Join(matchers.GetSupportedNames(), ", "))
	ErrorNotFoundToSceneInMather = errors.New("not found scene that was bee settled in the matcher. " +
		"The name of the scene must be specified in the nextScene field of the current scene")
	ErrorNotFoundToSceneInButton = errors.New("not found scene that was bee settled in the button. " +
		"The name of the scene must be specified in the nextScene field of the current scene")
	ErrorNotSupportedErrorType = errors.New("the base error name is not supported. Supported base errors is: " +
		strings.Join(errors2.GetSupportedNames(), ", "))
)

func errorUnknownTypeOfValue(notSupportedType string) error {
	return errors.Wrap(ErrorUnknownTypeOfValue, "with user type \""+notSupportedType+"\"")
}

func errorNotFoundValueInText(notFoundValue string) error {
	return errors.Wrap(ErrorNotFoundValueInText, "with value \""+notFoundValue+
		"\".  Correct name of value {"+notFoundValue+"}, "+
		"written without space and bounded by curly braces")
}

func errorNotSupportedMatherType(matcherName string) error {
	return errors.Wrap(ErrorNotSupportedMatherType, "with matcher name \""+matcherName+"\"")
}

func errorNotFoundToSceneInMather(sceneName, matcherName string) error {
	return errors.Wrap(ErrorNotFoundToSceneInMather, "in the matcher \""+matcherName+
		"\" with next scene name \""+sceneName+"\"")
}

func errorNotFoundToSceneInButton(sceneName, buttonName string) error {
	return errors.Wrap(ErrorNotFoundToSceneInButton, "in the button \""+buttonName+
		"\" with next scene name \""+sceneName+"\"")
}

func errorNotSupportedErrorType(errorName string) error {
	return errors.Wrap(ErrorNotSupportedErrorType, "with error name \""+errorName+"\"")
}
