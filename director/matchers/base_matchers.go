package matchers

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words2num"
	"strconv"
)

var (
	NumberMatcher         = NewRegexMather(`[\-]{0,1}[0-9]+[\.][0-9]+|[\-]{0,1}[0-9]+`)
	PositiveNumberMatcher = NewRegexMather(`^\+?(0*[1-9]\d*(?:[\., ]\d+)*) *(?:\p{Sc}|°[FC])?$`)
	AnyMatcher            = NewRegexMather(`.*`)
	FirstWord             = NewRegexMather(`[^\s]+`)
)

const (
	AgreeString = "Точно!"
)

var (
	Agree = NewSelectorMatcher(
		[]string{
			"Точно",
			"Согласен",
			"Да",
			"Ага",
		},
		AgreeString,
	)
)

type NumberMatchers struct{}

func (nm NumberMatchers) Match(message string) (bool, string) {
	res, _ := words2num.Convert(message)
	return res != 0, strconv.FormatInt(res, 10)
}

var PositiveNumberInWordsMatcher = NumberMatchers{}
