package matchers

import "github.com/thcompiler/go_game_constractor/director/scene"

const (
	positiveNumberErrorString = "Я не знаю такого положительного числа"
	numberErrorString         = "Я не знаю такого положительного числа"
)

var (
	NumberMatcher         = NewRegexMather(`[\-]{0,1}[0-9]+[\.][0-9]+|[\-]{0,1}[0-9]+`)
	PositiveNumberMatcher = NewRegexMather(`^\+?(0*[1-9]\d*(?:[\., ]\d+)*) *(?:\p{Sc}|°[FC])?$`)
	AnyMatcher            = NewRegexMather(`.*`)
	FirstWord             = NewRegexMather(`[^\s]+`)
	PositiveNumberError   = scene.BaseTextError{
		Message: positiveNumberErrorString,
	}
	NumberError = scene.BaseTextError{
		Message: numberErrorString,
	}
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
