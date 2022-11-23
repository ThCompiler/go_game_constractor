package matchers

import (
	"github.com/ThCompiler/go_game_constractor/director/scriptdirector/scene"
)

const (
	positiveNumberErrorString = "Я не знаю такого положительного числа"
	numberErrorString         = "Я не знаю такого целого числа"
)

var (
	PositiveNumberError = scene.BaseTextError{
		Message: positiveNumberErrorString,
	}
	NumberError = scene.BaseTextError{
		Message: numberErrorString,
	}
)
