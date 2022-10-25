package scene

import (
	"github.com/pkg/errors"
	"github.com/thcompiler/go_game_constractor/scg/go/types"
)

var (
	errorIsNotRegexMatcher    = errors.New("this matcher is not regex, but you try get regex matcher")
	errorIsNotSelectsMatcher  = errors.New("this matcher is not selects, but you try get selects matcher")
	errorIsNotStandardMatcher = errors.New("this matcher is not standard, but you try get standard matcher")
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
