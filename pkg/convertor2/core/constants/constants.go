package constants

import "github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/words"

type NumberType string

const (
	DECIMAL_NUMBER    = NumberType(",")
	FRACTIONAL_NUMBER = NumberType("/")
)

const (
	NUMBER = words.CurrencyName("number")
)

const (
	MaxNumberPartLength = 306

	TwoSignAfterRoundForCurrency = 2

	CountDigits          = 10
	CountWordForms       = 2
	CountNumberNameForms = 3

	UnitsScale    = 1
	ThousandScale = 2

	TripletSize = 3
)
