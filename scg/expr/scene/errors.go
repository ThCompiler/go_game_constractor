package scene

import (
	"github.com/ThCompiler/go_game_constractor/scg/go/types"
	errors2 "github.com/ThCompiler/go_game_constractor/scg/script/errors"
	"github.com/ThCompiler/go_game_constractor/scg/script/matchers"
	"github.com/pkg/errors"
	"strings"
)

var (
	errorIsNotRegexMatcher              = errors.New("this matcher is not regex, but you try get regex matcher")
	errorIsNotSelectsMatcher            = errors.New("this matcher is not selects, but you try get selects matcher")
	errorIsNotStandardMatcher           = errors.New("this matcher is not standard, but you try get standard matcher")
	errorEmptyNextSceneWithInfoScene    = errors.New("if scene is info scene use only flag NextScene")
	errorEmptyNextScenesWithNoInfoScene = errors.New("if scene is not info scene use only flag NextScenes")
)

func errorUnknownTypeOfValue(notSupportedType string) error {
	return errors.New("this type \"" + notSupportedType +
		"\" not supported. Supported type is:" + types.GetSupportTypes())
}

func errorNotFoundValueInText(notFoundValue string) error {
	return errors.New("this value \"" + notFoundValue +
		"\" not found in text. Correct name of value {" + notFoundValue + "}, " +
		"written without space and bounded by curly braces")
}

func errorNotSupportedMatherType(matcherName string) error {
	return errors.New("this matcher name \"" + matcherName +
		"\" not supported. Supported mathers is: " + strings.Join(matchers.GetSupportedNames(), ", "))
}

func errorNotSupportedErrorType(errorName string) error {
	return errors.New("this base error name \"" + errorName +
		"\" not supported. Supported base errors is: " + strings.Join(errors2.GetSupportedNames(), ", "))
}
