package matchers

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words2num"
	"strconv"
)

const (
	NumberMatchedString                = "number"
	PositiveNumberMatchedString        = "positive_number"
	AnyMatchedString                   = "any"
	FirstWordMatchedString             = "first_word"
	PositiveNumberInWordsMatchedString = "positive_number_in_words"
)

var (
	NumberMatcher         = NewRegexMather(`[\-]{0,1}[0-9]+[\.][0-9]+|[\-]{0,1}[0-9]+`, NumberMatchedString)
	PositiveNumberMatcher = NewRegexMather(`^\+?(0*[1-9]\d*(?:[\., ]\d+)*) *(?:\p{Sc}|°[FC])?$`, PositiveNumberMatchedString)
	AnyMatcher            = NewRegexMather(`.*`, AnyMatchedString)
	FirstWordMatcher      = NewRegexMather(`[^\s]+`, FirstWordMatchedString)
)

const (
	AgreeMatchedString = "Точно!"
)

var (
	Agree = NewSelectorMatcher(
		[]string{
			"Точно",
			"Согласен",
			"Да",
			"Ага",
		},
		AgreeMatchedString,
	)
)

type NumberMatchers struct{}

func (nm NumberMatchers) Match(message string) (bool, string) {
	res, err := words2num.Convert(message)

	return res != 0 && err == nil, strconv.FormatInt(res, 10)
}

func (nm NumberMatchers) GetMatchedName() string {
	return PositiveNumberInWordsMatchedString
}

var PositiveNumberInWordsMatcher = NumberMatchers{}
